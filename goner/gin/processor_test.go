package gin

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/gone-io/gone"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func (p *sysProcessor) Errorf(format string, args ...any) {}
func (p *sysProcessor) Warnf(format string, args ...any)  {}
func (p *sysProcessor) Infof(format string, args ...any)  {}

func Test_sysProcessor_AfterRevive(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Run("ShowRequestTime", func(t *testing.T) {
		iRouter := NewMockIRouter(controller)
		iRouter.EXPECT().Use(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any())

		processor := sysProcessor{
			router:          iRouter,
			ShowRequestTime: true,
			ShowAccessLog:   true,
		}
		_ = processor.AfterRevive()
	})
}

type testTracer struct {
	gone.Tracer
}

func (t *testTracer) SetTraceId(traceId string, fn func()) {
	if traceId == "" {
		traceId = "1234567890"
	}
	fn()
}

func (t *testTracer) GetTraceId() string {
	return "1234567890"
}

func Test_sysProcessor_trace(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Run("HealthCheckUrl", func(t *testing.T) {
		writer := NewMockResponseWriter(controller)
		writer.EXPECT().WriteHeader(200)
		writer.EXPECT().WriteHeaderNow()

		processor := sysProcessor{
			HealthCheckUrl: "/health",
		}

		Url, _ := url.Parse("https://goner.fun/health")

		context := Context{
			Context: &gin.Context{
				Request: &http.Request{
					URL: Url,
				},
				Writer: writer,
			},
		}

		_, err := processor.trace(&context)
		assert.Nil(t, err)
	})

	t.Run("trace", func(t *testing.T) {
		processor := sysProcessor{
			tracer: &testTracer{},
		}

		Url, _ := url.Parse("https://goner.fun/health")

		context := Context{
			Context: &gin.Context{
				Request: &http.Request{
					URL: Url,
				},
			},
		}

		_, err := processor.trace(&context)
		assert.Nil(t, err)
	})
}

func Test_sysProcessor(t *testing.T) {
	processor := sysProcessor{
		tracer: &testTracer{},
	}

	Url, _ := url.Parse("https://goner.fun/health")

	context := Context{
		Context: &gin.Context{
			Request: &http.Request{
				URL: Url,
			},
		},
	}

	t.Run("recovery", func(t *testing.T) {
		_, err := processor.recovery(&context)
		assert.Nil(t, err)
	})

	t.Run("statRequestTime", func(t *testing.T) {
		_, err := processor.statRequestTime(&context)
		assert.Nil(t, err)
	})

	t.Run("accessLog", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		writer := NewMockResponseWriter(controller)
		writer.EXPECT().Header().Return(http.Header{})
		writer.EXPECT().Status().Return(200)

		context.Context.Writer = writer

		context.Context.Request.Body = io.NopCloser(strings.NewReader("test"))

		_, err := processor.accessLog(&context)
		assert.Nil(t, err)
	})

	t.Run("accessLog-logDataMaxLength", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		writer := NewMockResponseWriter(controller)
		writer.EXPECT().Header().Return(http.Header{})
		writer.EXPECT().Status().Return(200)
		processor.logDataMaxLength = 2

		context.Context.Writer = writer

		context.Context.Request.Body = io.NopCloser(strings.NewReader("test"))

		_, err := processor.accessLog(&context)
		assert.Nil(t, err)
	})

	t.Run("accessLog-json", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		writer := NewMockResponseWriter(controller)
		writer.EXPECT().Header().Return(http.Header{
			"Content-Type": {"application/json"},
		})
		writer.EXPECT().Status().Return(200)
		processor.logDataMaxLength = 2

		context.Context.Writer = writer

		context.Context.Request.Body = io.NopCloser(strings.NewReader("test"))

		_, err := processor.accessLog(&context)
		assert.Nil(t, err)
	})
}

func Test_sysProcessor_recover(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockResponser := NewMockResponser(controller)

	processor := sysProcessor{
		tracer:     &testTracer{},
		resHandler: mockResponser,
	}

	Url, _ := url.Parse("https://goner.fun/health")

	context := Context{
		Context: &gin.Context{
			Request: &http.Request{
				URL: Url,
			},
		},
	}

	mockResponser.EXPECT().Failed(gomock.Any(), gomock.Any()).Do(func(ctx any, err gone.InnerError) {
		assert.Equal(t, 500, err.Code())
	})

	func() {
		defer processor.recover(&context)
		panic("err")
	}()

	assert.Nil(t, nil)
}

func TestCustomResponseWrite(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	writer := NewMockResponseWriter(controller)
	writer.EXPECT().Write(gomock.Any()).Return(100, nil)
	writer.EXPECT().WriteString(gomock.Any()).Return(100, nil)

	blw := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: writer}

	write, err := blw.Write([]byte("test"))
	assert.Equal(t, 100, write)
	assert.Nil(t, err)

	write, err = blw.WriteString("test2")
	assert.Equal(t, 100, write)
	assert.Nil(t, err)

	assert.Equal(t, "testtest2", blw.body.String())
}
