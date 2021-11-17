package debugging

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/tinybear1976/localsystem/logger"
	"github.com/tinybear1976/tinycms/defines"
)

func Debug_ShowSql(tag, sql string) {
	if viper.GetBool(defines.YML_DEBUG_SQL) {
		o := "[" + tag + "] " + sql
		logger.LogContainer[defines.LOG_SQL].Debug(o)
		fmt.Println("Debug SQL: ", o)
	}
}
