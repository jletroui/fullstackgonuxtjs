package config

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func CreateWebServer(cfg *Config, installApiRoutes func(*gin.RouterGroup)) (*gin.Engine, error) {
	gin.DefaultWriter = new(slogWriter)
	err := ConfigureSuperTokens(cfg)
	if err != nil {
		return nil, err
	}

	router := gin.New()
	router.Use(
		createStructuredAccessLogHandler(),
		createStaticHandler(),
		gin.Recovery(),
		createCorsHandler(cfg),
		createSuperTokensHandler(),
	)

	apiRoutes := router.Group("/api")
	installApiRoutes(apiRoutes)

	return router, nil
}

func createCorsHandler(cfg *Config) gin.HandlerFunc {
	// https://supertokens.com/docs/emailpassword/pre-built-ui/setup/backend#3-add-the-supertokens-apis--cors-setup
	origins := append(cfg.AllowOrigins, cfg.SuperTokensUrl)
	slog.Info(fmt.Sprintf("Origins: %s", origins))
	return cors.New(cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowHeaders:     append([]string{"content-type"}, supertokens.GetAllCORSHeaders()...),
		AllowCredentials: true,
	})
}

func createStaticHandler() gin.HandlerFunc {
	innerHandler := static.Serve("/", static.LocalFile("www/", true))

	return func(ctx *gin.Context) {
		// We only want to pay the price of looking on the file system for non-API requests
		if !strings.HasPrefix(ctx.Request.URL.Path, "/api") {
			innerHandler(ctx)
		}
	}
}

func createSuperTokensHandler() gin.HandlerFunc {
	// https://supertokens.com/docs/emailpassword/pre-built-ui/setup/backend#3-add-the-supertokens-apis--cors-setup
	return func(c *gin.Context) {
		supertokens.Middleware(http.HandlerFunc(
			func(rw http.ResponseWriter, r *http.Request) {
				c.Next()
			})).ServeHTTP(c.Writer, c.Request)
		// we call Abort so that the next handler in the chain is not called, unless we call Next explicitly
		c.Abort()
	}
}

func createStructuredAccessLogHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Start timer
		start := time.Now()
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		// Process request
		ctx.Next()

		slog.Info(
			ctx.Errors.ByType(gin.ErrorTypePrivate).String(),
			slog.String("timestamp", start.Format(time.RFC1123)),
			slog.Int64("latency_ms", time.Since(start).Milliseconds()),
			slog.String("method", ctx.Request.Method),
			slog.String("remote_host", ctx.ClientIP()),
			slog.String("requested_uri", path),
			slog.String("protocol", ctx.Request.Proto),
			slog.Int("status_code", ctx.Writer.Status()),
			slog.Int("content_length", ctx.Writer.Size()),
			slog.String("referer", ctx.Request.Referer()),
			slog.String("user_agent", ctx.Request.UserAgent()),
		)
	}
}

type slogWriter struct{}

func (*slogWriter) Write(p []byte) (n int, err error) {
	// Note: terrible performance, but logging by GIN happens very rarely after boot, so ok.
	logFunc := slog.Debug
	s := string(p)
	if strings.Contains(s, "[WARNING]") {
		logFunc = slog.Warn
	}
	s = strings.Replace(s, "[GIN-debug] ", "", -1)
	s = strings.Replace(s, "[WARNING] ", "", -1)
	logFunc("[GIN] " + s)
	return len(p), nil
}
