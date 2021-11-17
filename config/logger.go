package config

import (
	"path"

	"github.com/tinybear1976/localsystem/logger"
	"github.com/tinybear1976/tinycms/defines"
)

func InitLoggers(p string) {
	initBaseLog(p)
	initDebugSqlLog(p)
}

func initBaseLog(p string) {
	logger.NewLogger(defines.LOG_APP, path.Join(p, defines.LOG_APP+".log"), "debug") // info, error, warn
	if logger.LogContainer[defines.LOG_APP] == nil {
		panic("logger creation failed, program operation aborted")
	}
}
func initDebugSqlLog(p string) {
	logger.NewLogger(defines.LOG_SQL, path.Join(p, defines.LOG_SQL+".log"), "debug") // info, error, warn
	if logger.LogContainer[defines.LOG_SQL] == nil {
		panic("logger creation failed, program operation aborted")
	}
}
