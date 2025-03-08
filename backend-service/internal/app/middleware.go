package app

import (
	"bytes"
	"io"
	"sync"
	"time"

	"backend-service/config"
	"backend-service/internal/pkg/errors"
	"backend-service/internal/pkg/helper"

	"github.com/gin-gonic/gin"
	"github.com/madevara24/go-common/constant"
	"github.com/madevara24/go-common/logger"
	"github.com/madevara24/go-common/response"
	"github.com/spf13/cast"

	"golang.org/x/time/rate"
)

type bodyWriter struct {
	gin.ResponseWriter
	body *[]byte
}

func (w *bodyWriter) Write(b []byte) (int, error) {
	*w.body = append(*w.body, b...)
	return w.ResponseWriter.Write(b)
}

func SetTDRMiddleware() gin.HandlerFunc {
	beginTime := time.Now()

	return func(c *gin.Context) {
		loggerCtx := logger.Context{
			ServiceName:    config.Get().AppName,
			ServiceVersion: config.Get().AppVersion,
			ServicePort:    cast.ToInt(config.Get().Port),
			ReqMethod:      c.Request.Method,
			ReqURI:         c.Request.URL.String(),
		}

		var (
			reqBody []byte
			resBody []byte
		)

		if c.Request.Body != nil {
			reqBody, _ = io.ReadAll(c.Request.Body)
			loggerCtx.ReqBody = helper.CleanJSON(string(reqBody))
			c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
		}

		writer := &bodyWriter{ResponseWriter: c.Writer, body: &resBody}
		c.Writer = writer

		ctx := logger.InjectCtx(c.Request.Context(), loggerCtx)
		c.Request = c.Request.WithContext(ctx)

		if logger.IsSkipLog(c.Request.Header.Get("Content-Type")) {
			logger.Log.Info(ctx, "Request Not Log Because Unsupported Content-Type")
		}

		// Continue Request Until Executed
		c.Next()

		if !logger.IsSkipLog(c.Writer.Header().Get("Content-Type")) {
			loggerCtx.RespBody = helper.CleanJSON(string(resBody))
		}

		if errMessage, exists := c.Get(constant.ErrorMessageKey); exists {
			loggerCtx.Error = cast.ToString(errMessage)
		}

		loggerCtx.ReqHeader = helper.ConvertMapStringToString(c.Request.Header)
		loggerCtx.RespTime = cast.ToString(time.Since(beginTime).Milliseconds())
		loggerCtx.RespCode = c.Writer.Status()

		ctx = logger.InjectCtx(c.Request.Context(), loggerCtx)

		// Logger Trace Distributed Request (TDR)
		// List all information in one request log
		logger.Log.TDR(ctx, "Request Information")
	}
}

func StaticKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := config.Get()

		// Get the API key from the request header
		apiKey := c.GetHeader("X-API-Key")

		// Compare with the configured key
		if apiKey != cfg.PrivateApiKey {
			response.WriteError(c, errors.ERR_UNAUTHORIZED)
			return
		}

		// Continue if the key is valid
		c.Next()
	}
}

type RateLimiter struct {
	visitors sync.Map
}

func (r *RateLimiter) GetLimiter(ip string) *rate.Limiter {
	cfg := config.Get()

	limiter, exists := r.visitors.Load(ip)
	if exists {
		return limiter.(*rate.Limiter)
	}

	newLimiter := rate.NewLimiter(rate.Every(time.Minute/time.Duration(cfg.RateLimitPerMinute)), cfg.RateLimitBurst)
	r.visitors.Store(ip, newLimiter)

	// Remove the limiter after inactivity
	go func() {
		time.Sleep(1 * time.Minute)
		r.visitors.Delete(ip)
	}()

	return newLimiter
}

func (r *RateLimiter) RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := r.GetLimiter(ip)

		if !limiter.Allow() {
			response.WriteError(c, errors.ERR_RATE_LIMIT_EXCEED)
			return
		}
		c.Next()
	}
}
