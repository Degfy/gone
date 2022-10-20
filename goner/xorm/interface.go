package xorm

import "xorm.io/xorm"

type Engine interface {
	xorm.Interface
	Transaction(fn func(session *xorm.Session) error) error
	Sqlx(sql string, args ...interface{}) *xorm.Session
	Start()
	GetOriginEngine() *xorm.Engine
}
