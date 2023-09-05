package menu_handler

import (
	"menu-service/dtos"
	"menu-service/ent"
	service "menu-service/service/menu"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

type MenuHandler struct {
	menuService service.IMenuService
}

type IMenuHandler interface {
	GetMenuById(c *fiber.Ctx) error
	GetAllMenus(c *fiber.Ctx) error
	CreateMenu(c *fiber.Ctx) error
	UpdateMenu(c *fiber.Ctx) error
	DeleteMenu(c *fiber.Ctx) error
}

func NewHandler(menuService service.IMenuService) IMenuHandler {
	return &MenuHandler{
		menuService: menuService,
	}
}

func (h *MenuHandler) GetAllMenus(c *fiber.Ctx) error {
	menus, err := h.menuService.GetAllMenus()

	if err != nil {
		c.Status(err.Status)
		c.JSON(dtos.ResponseError{
			Message: err.Message,
		})
		return nil
	}

	c.Status(fiber.StatusOK)
	c.JSON(menus)

	return nil
}

func (h *MenuHandler) GetMenuById(c *fiber.Ctx) error {
	id := utils.CopyString(c.Params("id"))

	num_id, err := strconv.Atoi(id)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		c.JSON(dtos.ResponseError{
			Message: "id must be a number",
		})
		return nil
	}

	menu, another_err := h.menuService.GetMenuById(num_id)

	if another_err != nil {
		c.Status(another_err.Status)
		c.JSON(dtos.ResponseError{
			Message: another_err.Message,
		})
		return nil
	}

	c.Status(fiber.StatusOK)
	c.JSON(menu)

	return nil
}

func (h *MenuHandler) CreateMenu(c *fiber.Ctx) error {
	var createBody *ent.Menu

	err := c.BodyParser(&createBody)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		c.JSON(dtos.ResponseError{
			Message: "invalid body",
		})
		return err
	}

	menu, another_err := h.menuService.CreateMenu(createBody)

	if another_err != nil {
		c.Status(another_err.Status)
		c.JSON(dtos.ResponseError{
			Message: another_err.Message,
		})
		return err
	}

	c.Status(fiber.StatusCreated)
	c.JSON(dtos.CreateMenuResponseDtos{
		ID: menu.ID,
	})

	return nil
}

func (h *MenuHandler) UpdateMenu(c *fiber.Ctx) error {
	id := utils.CopyString(c.Params("id"))

	num_id, err := strconv.Atoi(id)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		c.JSON(dtos.ResponseError{
			Message: "id must be a number",
		})
		return nil
	}

	var updateBody *ent.Menu

	err = c.BodyParser(&updateBody)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		c.JSON(dtos.ResponseError{
			Message: "invalid body",
		})
		return nil
	}

	another_err := h.menuService.UpdateMenu(num_id, updateBody)

	if another_err != nil {
		c.Status(another_err.Status)
		c.JSON(dtos.ResponseError{
			Message: another_err.Message,
		})
		return nil
	}

	c.Status(fiber.StatusOK)

	return nil
}

func (h *MenuHandler) DeleteMenu(c *fiber.Ctx) error {
	id := utils.CopyString(c.Params("id"))

	num_id, err := strconv.Atoi(id)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		c.JSON(dtos.ResponseError{
			Message: "id must be a number",
		})
		return nil
	}

	another_err := h.menuService.DeleteMenu(num_id)

	if another_err != nil {
		c.Status(another_err.Status)
		c.JSON(dtos.ResponseError{
			Message: another_err.Message,
		})
		return nil
	}

	c.Status(fiber.StatusOK)

	return nil
}
