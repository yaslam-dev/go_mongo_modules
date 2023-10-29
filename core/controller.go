package core

import (
	"fmt"
	"path"
	"strings"

	"github.com/gofiber/fiber/v2"
)

/**
- Base Handler
- with Builder pattern
*/

type Handler struct {
	Method      HTTPMethods
	Path        string
	Middleware  []FiberHandler
	HandlerFunc FiberHandler
	RequestDto  interface{}
	ResponseDto interface{}
	Description string
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) SetMethod(method HTTPMethods) *Handler {
	h.Method = method
	return h
}

func (h *Handler) SetPath(path string) *Handler {
	h.Path = path
	return h
}

func (h *Handler) AddMiddleware(m ...FiberHandler) *Handler {
	h.Middleware = append(h.Middleware, m...)
	return h
}

func (h *Handler) SetHandlerFunc(handlerFunc FiberHandler) *Handler {
	h.HandlerFunc = handlerFunc
	return h
}

func (h *Handler) SetDescription(description string) *Handler {
	h.Description = description
	return h
}

func (h *Handler) SetRequestDto(requestDto interface{}) *Handler {
	h.RequestDto = requestDto
	return h
}

func (h *Handler) SetResponseDto(responseDto interface{}) *Handler {
	h.ResponseDto = responseDto
	return h
}

/**
- Base controller
- with Builder pattern
*/

type Controller struct {
	Name     string
	Version  string
	Path     string
	Handlers []*Handler
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) SetName(name string) *Controller {
	c.Name = name
	return c
}

func (c *Controller) SetVersion(version string) *Controller {
	c.Version = version
	return c
}

func (c *Controller) SetPath(path string) *Controller {
	c.Path = path
	return c
}

func (c *Controller) AddHandler(h *Handler) *Controller {
	c.Handlers = append(c.Handlers, h)
	return c
}

func (c *Controller) RegisterRoutes(e *fiber.App) {
	for _, h := range c.Handlers {
		fullPath := path.Join(c.Version, c.Path, h.Path)
		fmt.Println("    H", register, " NewHandler: ", GetStatusString(h.Method), fullPath)
		h.Middleware = append(h.Middleware, h.HandlerFunc)
		e.Add(strings.ToUpper(string(h.Method)), fullPath, h.Middleware...)
	}
}
