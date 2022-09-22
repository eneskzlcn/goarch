package main

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
)

const (
	Fiber    = "f"
	Gin      = "g"
	Kafka    = "k"
	Rabbitmq = "r"
	Postgres = "p"
	Mongo    = "m"
	Logger   = "l"
)

const (
	NLayeredWeb         = "nl-web"
	Microservice        = "ms-rkpfml"
	DefaultArchitecture = Microservice
	DefaultDirectory    = "."
)

func CreateArchitecture(arch, targetPath string) error {
	switch arch {
	case NLayeredWeb:
		return CreateNLayeredWebArchitecture()
	case Microservice:
		return CreateMicroserviceArchitecture()
	}
	return errors.New("not valid operation")
}

const (
	ConfigurationYamlContent = `---
db:
  host: "hostname"
  port: "port"
  dbName: "dbname"`
)
const (
	FiberServerContent = `
package server

import (
	"fmt"
	config "your_module_name/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
	"os/signal"
	"syscall"
)

type Handler interface {
	RegisterRoutes(app *fiber.App)
}

type Logger interface {
	Error(args ...interface{})
}
type Server struct {
	app    *fiber.App
	config config.Server
	logger Logger
}

func New(handlers []Handler, config config.Server, logger Logger) *Server {
	app := fiber.New()
	app.Use(cors.New())
	for _, handler := range handlers {
		handler.RegisterRoutes(app)
	}
	server := &Server{
		app:    app,
		config: config,
		logger: logger,
	}
	server.AddRoutes()
	return server
}
func (s *Server) AddRoutes() {
	s.app.Get("/health", s.healthCheck)
}
func (s *Server) healthCheck(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK)
}

func (s *Server) Start() error {
	address := fmt.Sprintf(":%s", s.config.Port)
	shutDownChan := make(chan os.Signal, 1)
	signal.Notify(shutDownChan, os.Interrupt, syscall.SIGTERM, os.Kill)
	go func() {
		<-shutDownChan
		err := s.app.Shutdown()
		if err != nil {
			s.logger.Error(err)
	}()
	return s.app.Listen(address)
}`
)
const (
	GinServerContent = `
package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler interface {
	RegisterRoutes(engine *gin.Engine)
}
type Logger interface {
	Debugf(template string, args ...interface{})
}
type Server struct {
	engine *gin.Engine
	port   string
	logger Logger
}

func NewServer(handlers []Handler, trustedProxies []string, port string, logger Logger) *Server {
	engine := gin.Default()
	engine.SetTrustedProxies(trustedProxies)
	server := Server{port: port, engine: engine, logger: logger}
	for _, handler := range handlers {
		handler.RegisterRoutes(server.engine)
	}
	server.AddRoutes()
	return &server
}

func (s *Server) AddRoutes() {
	s.engine.GET("/health", s.healthCheck)
}

func (s *Server) healthCheck(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

func (s *Server) Start() error {
	s.logger.Debugf("server started to listen at: %s", s.port)
	return s.engine.Run(fmt.Sprintf("localhost:%s", s.port))
}`
)

func CreateDirectory(name string) error {
	return exec.Command("mkdir", name).Run()
}

func CreateFile(absolutePathWithName string) error {
	return exec.Command("touch", absolutePathWithName).Run()
}

func CreateFileWithContent(absolutePathWithName, content string) error {
	workingDir, _ := os.Getwd()
	file, err := os.OpenFile(workingDir+"/"+absolutePathWithName, os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(file, content)
	if err != nil {
		return err
	}
	return nil
}

func PrepareDevDirectory() error {
	if err := CreateDirectory(".dev"); err != nil {
		return err
	}

	if err := CreateFileWithContent(".dev/dev.yaml", ConfigurationYamlContent); err != nil {
		return err
	}

	if err := CreateFileWithContent(".dev/qa.yaml", ConfigurationYamlContent); err != nil {
		return err
	}

	if err := CreateFileWithContent(".dev/prod.yaml", ConfigurationYamlContent); err != nil {
		return err
	}

	if err := CreateFileWithContent(".dev/local.yaml", ConfigurationYamlContent); err != nil {
		return err
	}

	return nil
}

func CreateCmdDirectory() error {
	if err := CreateDirectory("cmd"); err != nil {
		return err
	}

	currentDirectory, _ := os.Getwd()
	projectDirName := path.Base(currentDirectory)
	if err := CreateDirectory(fmt.Sprintf("cmd/%s", projectDirName)); err != nil {
		return err
	}

	if err := CreateFile(fmt.Sprintf("cmd/%s/main.go", projectDirName)); err != nil {

	}
	return nil
}

func PrepareServerDirectory(serverTechnology string) error {
	if err := CreateDirectory("server"); err != nil {
		return err
	}
	switch serverTechnology {
	case Fiber:
		if err := CreateFileWithContent("server/server.go", FiberServerContent); err != nil {
			return err
		}
		break
	case Gin:
		if err := CreateFileWithContent("server/server.go", GinServerContent); err != nil {
			return err
		}
		break
	default:
		if err := CreateFileWithContent("server/server.go", `package server`); err != nil {
			return err
		}
		break
	}
	if err := CreateFileWithContent("server/server_test.go", `package server_test`); err != nil {
		return err
	}

	return nil
}

func CreateMicroserviceArchitecture() error {

	//if err := PrepareDevDirectory(); err != nil {
	//	return err
	//}
	//if err := CreateCmdDirectory(); err != nil {
	//	return err
	//}
	if err := PrepareServerDirectory(Gin); err != nil {
		return err
	}

	//if err := PrepareInternalDirectory(); err != nil {
	//	return err
	//}
	//
	//if err := PreparePostgresDirectory(); err != nil {
	//	return err
	//}
	//if err := PrepareLoggerDirectory(); err != nil {
	//	return err
	//}
	//if err := PrepareRabbitMQDirectory(); err != nil {
	//	return err
	//}
	//if err := PrepareSeedDirectory(); err != nil {
	//	return err
	//}
	return nil
}

func CreateNLayeredWebArchitecture() error {
	return nil
}

func main() {
	var err error
	if len(os.Args) <= 1 {
		err = CreateArchitecture(DefaultArchitecture, DefaultDirectory)
	} else {
		err = CreateArchitecture(os.Args[1], DefaultDirectory)
	}
	if err != nil {
		fmt.Println("error occurred when creating architecture", err.Error())
	}
}
