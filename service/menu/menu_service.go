package menu_service

import (
	"context"
	"fmt"
	"menu-service/dtos"
	"menu-service/ent"
	"menu-service/ent/menu"
)

type MenuService struct {
	db *ent.Client
}

type IMenuService interface {
	GetMenuById(id int) (*ent.Menu, *dtos.Error)
	GetAllMenus() ([]*ent.Menu, *dtos.Error)
	CreateMenu(menu *ent.Menu) (*ent.Menu, *dtos.Error)
	UpdateMenu(id int, menu *ent.Menu) *dtos.Error
	DeleteMenu(id int) *dtos.Error
}

func NewService(db *ent.Client) IMenuService {
	return &MenuService{
		db: db,
	}
}

func (s *MenuService) GetMenuById(id int) (*ent.Menu, *dtos.Error) {
	menu, err := s.db.Menu.Query().Where(menu.ID(id)).Only(context.Background())

	if ent.IsNotFound(err) {
		return nil, &dtos.Error{
			Status:  404,
			Message: fmt.Sprintf("menu with id %d not found", id),
		}
	} else if err != nil {
		return nil, &dtos.Error{
			Status:  500,
			Message: err.Error(),
		}
	}

	return menu, nil
}

func (s *MenuService) GetAllMenus() ([]*ent.Menu, *dtos.Error) {
	menus, err := s.db.Menu.Query().All(context.Background())

	if err != nil {
		return nil, &dtos.Error{
			Status:  500,
			Message: err.Error(),
		}
	}

	return menus, nil
}

func (s *MenuService) CreateMenu(menu *ent.Menu) (*ent.Menu, *dtos.Error) {
	menu, err := s.db.Menu.Create().SetName(menu.Name).SetDescription(menu.Description).SetPrice(menu.Price).Save(context.Background())

	if ent.IsConstraintError(err) {
		return nil, &dtos.Error{
			Status:  400,
			Message: err.Error(),
		}
	} else if err != nil {
		return nil, &dtos.Error{
			Status:  500,
			Message: err.Error(),
		}
	}

	return menu, nil
}

func (s *MenuService) UpdateMenu(id int, menu *ent.Menu) *dtos.Error {
	_, err := s.db.Menu.UpdateOneID(id).SetName(menu.Name).SetDescription(menu.Description).SetPrice(menu.Price).Save(context.Background())

	if ent.IsNotFound(err) {
		return &dtos.Error{
			Status:  404,
			Message: fmt.Sprintf("menu with id %d not found", id),
		}
	}

	if err != nil {
		return &dtos.Error{
			Status:  500,
			Message: err.Error(),
		}
	}

	return nil
}

func (s *MenuService) DeleteMenu(id int) *dtos.Error {
	err := s.db.Menu.DeleteOneID(id).Exec(context.Background())

	if ent.IsNotFound(err) {
		return &dtos.Error{
			Status:  404,
			Message: fmt.Sprintf("menu with id %d not found", id),
		}
	}

	if err != nil {
		return &dtos.Error{
			Status:  500,
			Message: err.Error(),
		}
	}

	return nil
}
