package logging

import (
	"fhir-sirs/app/config/constant"
	log "fhir-sirs/app/logger/singleton"
	srvReq "fhir-sirs/app/server/request"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"time"
)

func Logging() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqID := srvReq.ID()
			c.Set("requestID", reqID)
			defer func(now time.Time) {

				// flush cache
				// cache.Delete(reqID)

				message := constant.LLvlAccess
				fields := []zap.Field{
					zap.String("at", now.Format(constant.DateFormatWithTime)),
					zap.String("method", c.Request().Method),
					zap.String("uri", c.Request().URL.String()),
					zap.String("ip", c.RealIP()),
					zap.String("host", c.Request().Host),
					zap.String("user_agent", c.Request().UserAgent()),
					zap.Int("code", c.Response().Status),
				}
				log.WithRequestID(reqID).Info(message, fields...)
			}(time.Now())
			return next(c)
		}
	}
}
