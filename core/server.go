package core

import (
	"fmt"
	"strings"
	"time"

	json "github.com/goccy/go-json"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Server struct {
	Port       int
	Engine     *fiber.App
	MainModule *Module
}

func (s *Server) Start() error {
	return s.Engine.Listen(fmt.Sprintf(":%v", s.Port))
}

func (s *Server) RegisterMainModule() {
	s.MainModule.RegisterRoutes(s.Engine)
}

var server *Server

const (
	TimeoutInSeconds = 10
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	response := map[string]interface{}{
		"error":   true,
		"message": "Internal Server Error"}
	fmt.Printf("Error: on %v with method %v: %v\n", ctx.Route().Path, ctx.Route().Method, err)
	if ConfigInstance().Env != Production {
		response["details"] = err.Error()
	}
	ctx.Status(500).JSON(response)
	return nil
}

func NewServer() *Server {
	if server == nil {
		engine := fiber.New(fiber.Config{
			AppName:               ConfigInstance().AppName,
			ReadTimeout:           time.Duration(TimeoutInSeconds) * time.Second,
			WriteTimeout:          time.Duration(TimeoutInSeconds) * time.Second,
			ErrorHandler:          ErrorHandler,
			DisableStartupMessage: false,
			JSONEncoder:           json.Marshal,
			JSONDecoder:           json.Unmarshal,
			Prefork:               true,
		})
		engine.Use(recover.New(recover.Config{
			EnableStackTrace: true,
		}))

		engine.Use(compress.New())
		engine.Use(limiter.New(limiter.Config{
			Max:               20,
			Expiration:        30 * time.Second,
			LimiterMiddleware: limiter.SlidingWindow{},
		}))
		engine.Get("/server/metrics", monitor.New(monitor.Config{
			Title:   fmt.Sprintf("%s Metrics Page", strings.ToUpper(ConfigInstance().AppName)),
			Refresh: 1 * time.Second,
		}))

		engine.Use(requestid.New())
		engine.Use(logger.New(logger.Config{
			Format: "[${time}] x-request-id ${locals:requestid} ${method} ${path} ${status}  - ${latency} ${ip} ${ua} - ${error} \n",
		}))

		go func() {
			db := &Database{}
			db.InitializeDatabase()
		}()

		server = &Server{
			Port:   ConfigInstance().Port,
			Engine: engine,
		}
		fmt.Printf("%s Server starting on Port %d\n", starting, server.Port)
	}
	return server
}
