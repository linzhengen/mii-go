package logger

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/zero-contrib/logx/zapx"
)

type Logger = logx.Logger

func init() {
	writer, err := zapx.NewZapWriter()
	logx.Must(err)
	logx.SetWriter(writer)
}

var (
	Debug       = logx.Debug
	Debugf      = logx.Debugf
	Error       = logx.Error
	Errorf      = logx.Error
	Info        = logx.Info
	Infof       = logx.Info
	Severe      = logx.Severe
	Severef     = logx.Severef
	WithContext = logx.WithContext
	WithFields  = logx.WithFields
)
