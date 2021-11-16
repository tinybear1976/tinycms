package defines

import (
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	SIGN_LOGIN  = "login"
	SIGN_LOGOUT = "logout"

	DB_MAIN          = "tcms"
	DB_LOG_CLIENTIP  = "logclientip"
	REDIS_AUTH_TOKEN = "audit_auth_token"

	VERSION = "1.1.0"
	LOGO    = "\n" +
		" _____  __  _   __  ___" + "        " + "TinyCMS Api Service " + VERSION + "\n" +
		"/_  _/,'_/ / \\,' /,' _/" + "        " + "IP:    {{{{ip}}}}" + "\n" +
		" / / / /_ / \\,' /_\\ `. " + "        " + "Port:  {{{{port}}}}" + "\n" +
		"/_/  |__//_/ /_//___,'" + "         " + "{{{{pid}}}}" + "\n" +
		"                      " + "         " + "Start: {{{{datetime}}}}" + "\n\n" +
		"Running...\n"
)

func GetLogo(ip string) string {
	pid := ""
	switch runtime.GOOS {
	case "linux":
		fallthrough
	case "darwin":
		pid = "PID:   " + strconv.Itoa(os.Getpid())
	case "windows":
	}
	spl := strings.Split(ip, ":")
	tmp := strings.ReplaceAll(LOGO, "{{{{ip}}}}", spl[0])
	tmp = strings.ReplaceAll(tmp, "{{{{port}}}}", spl[1])
	tmp = strings.ReplaceAll(tmp, "{{{{datetime}}}}", time.Now().Format(FORMATDATETIME))
	tmp = strings.ReplaceAll(tmp, "{{{{pid}}}}", pid)
	return tmp
}
