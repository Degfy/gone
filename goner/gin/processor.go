package gin

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gone-io/gone"
	"github.com/gone-io/gone/goner/logrus"
	"github.com/gone-io/gone/goner/tracer"
	"io/ioutil"
	"strings"
	"time"
)

// NewGinProcessor 创建系统级的`Processor`，该处理器将在路由上挂载多个中间件
// `trace`，负责提供健康检查响应(需要在配置文件中设置健康检查的路径`server.health-check`) 和 为请求绑定唯一的`traceID`
// `recovery`，负责恢复请求内发生panic
// `statRequestTime`，用于在日志中打印统计的请求耗时，可以通过设置配置项(`server.log.show-request-time=false`)来关闭
// `accessLog`，用于在日志中打印请求、响应信息
func NewGinProcessor() (gone.Goner, gone.GonerId) {
	return &sysProcessor{}, gone.IdGoneGinProcessor
}

type sysProcessor struct {
	gone.Flag
	logrus.Logger `gone:"gone-logger"`

	tracer     tracer.Tracer `gone:"gone-tracer"`
	router     IRouter       `gone:"gone-gin-router"`
	resHandler Responser     `gone:"gone-gin-responser"`

	// HealthCheckUrl 健康检查路劲
	// 对应配置项为: `server.health-check`
	// 默认为空，不开启；
	// 配置后，能够在该路劲提供一个http-status等于200的空响应
	HealthCheckUrl string `gone:"config,server.health-check"`

	// ShowAccessLog 展示access日志
	// 对应配置项为：`server.log.show-access-log`
	// 默认为`true`;
	// 开启后，日志中将使用`Info`级别打印请求的request和response信息
	ShowAccessLog bool `gone:"config,server.log.show-access-log,default=true"`

	// ShowRequestTime 展示请求时间
	// 对应配置项为：`server.log.show-request-time`
	// 默认为`true`;
	// 开启后，日志中将使用`Info`级别打印请求的 耗时
	ShowRequestTime bool `gone:"config,server.log.show-request-time,default=true"`
}

func (p *sysProcessor) AfterRevive(gone.Cemetery, gone.Tomb) gone.ReviveAfterError {
	m := []gin.HandlerFunc{p.trace, p.recovery}
	if p.ShowRequestTime {
		m = append(m, p.statRequestTime)
	}
	if p.ShowAccessLog {
		m = append(m, p.accessLog)
	}
	p.router.Use(m...)
	return nil
}

var RequestIdHeaderKey = "X-Request-ID"
var TraceIdHeaderKey = "X-Trace-ID"

func (p *sysProcessor) trace(context *gin.Context) {
	if p.HealthCheckUrl != "" && context.Request.URL.Path == "/api/health-check" {
		context.AbortWithStatus(200)
		return
	}

	traceId := context.GetHeader(TraceIdHeaderKey)
	p.tracer.SetTraceId(traceId, func() {
		requestID := context.GetHeader(RequestIdHeaderKey)
		p.Infof("bind requestId:%s", requestID)
		context.Next()
	})
}

func (p *sysProcessor) recovery(context *gin.Context) {
	traceID := p.tracer.GetTraceId()
	defer func() {
		if r := recover(); r != nil {
			p.Errorf("[%s] handle panic: %v, %s", traceID, r, gone.PanicTrace(2))
			err, ok := r.(error)
			if !ok {
				err = errors.New(fmt.Sprintf("%v", r))
			}
			p.resHandler.Failed(context, err)
			context.Abort()
		}
	}()
	context.Next()
}

func (p *sysProcessor) statRequestTime(c *gin.Context) {
	beginTime := time.Now()
	defer func() {
		p.Infof("request(%s %s) use time: %v", c.Request.Method, c.Request.URL.Path, time.Now().Sub(beginTime))
	}()
	c.Next()
}

func (p *sysProcessor) accessLog(c *gin.Context) {
	remoteIP := c.GetHeader("X-Forwarded-For")
	if remoteIP == "" {
		remoteIP = c.RemoteIP()
	}

	data, err := cloneRequestBody(c)
	if err != nil {
		p.Error("accessLog - cloneRequestBody error:", err)
	}

	p.Infof("api-request|%s %s %s %s %s %s\n",
		remoteIP,
		c.Request.Method,
		c.Request.RequestURI,
		c.Request.UserAgent(),
		c.GetHeader("Referer"),
		data,
	)

	blw := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw
	c.Next()

	contentType := c.Writer.Header().Get("Content-Type")
	if strings.Contains(contentType, "json") {
		p.Infof("api-response|%v %s\n", c.Writer.Status(), blw.body.String())
	} else {
		p.Infof("api-response|%v %s\n", c.Writer.Status(), contentType)
	}
}

func cloneRequestBody(c *gin.Context) ([]byte, error) {
	data, err := c.GetRawData()
	if err != nil {
		return nil, err
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	return data, nil
}

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w CustomResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}