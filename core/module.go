package core

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Entity = interface{}

type Module struct {
	Name        string
	Description string
	Controllers []*Controller
	Imports     []*Module
	Entities    []Entity
	Extensions  interface{}
}

func NewModule() *Module {
	return &Module{}
}

func (m *Module) SetName(name string) *Module {
	m.Name = name
	return m
}

func (m *Module) SetDescription(description string) *Module {
	m.Description = description
	return m
}

func (m *Module) AddController(c *Controller) *Module {
	m.Controllers = append(m.Controllers, c)
	return m
}

func (m *Module) AddImport(im *Module) *Module {
	m.Imports = append(m.Imports, im)
	return m
}

func (m *Module) AddEntity(e Entity) *Module {
	m.Entities = append(m.Entities, e)
	return m
}

func (m *Module) SetExtensions(ext interface{}) *Module {
	m.Extensions = ext
	return m
}

func (m *Module) RegisterRoutes(e *fiber.App) {
	fmt.Println("M", register, " NewModule: ", m.Name)
	for _, c := range m.Controllers {
		fmt.Println("   C", register, " NewController: ", c.Path)
		c.RegisterRoutes(e)
	}
	for _, imported := range m.Imports {
		imported.RegisterRoutes(e)
	}
}
