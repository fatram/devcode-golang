package pkg

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

const DefaultLogHeader = `{"time":"${time_rfc3339_nano}","level":"${level}","prefix":"${prefix}",` +
	`"file":"${long_file}:${line}"}`

func LoadEcho(level ...string) (e *echo.Echo) {
	e = echo.New()
	e.Logger.SetHeader(DefaultLogHeader)
	e.Logger.SetLevel(log.DEBUG)
	return e
}
