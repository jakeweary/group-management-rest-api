package api

import (
	db "api/internal/database"

	"github.com/gofiber/fiber/v2"
)

func (api *API) SetRoutes(mux *fiber.App) {
	mux.Get("/groups", api.GetAllGroups)

	mux.Post("/group", api.CreateGroup)
	mux.Put("/group/:id", api.UpdateGroup)
	mux.Delete("/group/:id", api.DeleteGroup)
	mux.Get("/group/:id", api.GetGroup)

	mux.Post("/user", api.CreateUser)
	mux.Put("/user/:id", api.UpdateUser)
	mux.Delete("/user/:id", api.DeleteUser)
	mux.Get("/user/:id", api.GetUser)
}

// ---

func (api *API) GetAllGroups(c *fiber.Ctx) error {
	type query struct {
		Subgroups bool `query:"subgroups"`
	}

	q := query{}
	if err := c.QueryParser(&q); err != nil {
		return err
	}

	groups, err := api.db.GetAllGroups(q.Subgroups)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"groups": groups})
}

// ---

func (api *API) CreateGroup(c *fiber.Ctx) error {
	group := db.NewGroup{}
	if err := c.BodyParser(&group); err != nil {
		return ErrInvalidRequestBody
	}

	id, err := api.db.CreateGroup(group)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"group_id": id})
}

func (api *API) UpdateGroup(c *fiber.Ctx) error {
	type params struct {
		Id db.Id `param:"id"`
	}

	p := params{}
	if err := c.ParamsParser(&p); err != nil {
		return ErrInvalidUrlParams
	}

	group := db.NewGroup{}
	if err := c.BodyParser(&group); err != nil {
		return ErrInvalidRequestBody
	}

	err := api.db.UpdateGroup(p.Id, group)
	if err != nil {
		return err
	}

	return nil
}

func (api *API) DeleteGroup(c *fiber.Ctx) error {
	type params struct {
		Id db.Id `param:"id"`
	}

	p := params{}
	if err := c.ParamsParser(&p); err != nil {
		return ErrInvalidUrlParams
	}

	err := api.db.DeleteGroup(p.Id)
	if err != nil {
		return err
	}

	return nil
}

func (api *API) GetGroup(c *fiber.Ctx) error {
	type params struct {
		Id db.Id `param:"id"`
	}

	p := params{}
	if err := c.ParamsParser(&p); err != nil {
		return ErrInvalidUrlParams
	}

	type query struct {
		Subgroups bool `query:"subgroups"`
	}

	q := query{}
	if err := c.QueryParser(&q); err != nil {
		return ErrInvalidUrlQueryParams
	}

	group, users, err := api.db.GetGroup(p.Id, q.Subgroups)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"group": group, "users": users})
}

// ---

func (api *API) CreateUser(c *fiber.Ctx) error {
	group := db.NewUser{}
	if err := c.BodyParser(&group); err != nil {
		return ErrInvalidRequestBody
	}

	id, err := api.db.CreateUser(group)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"user_id": id})
}

func (api *API) UpdateUser(c *fiber.Ctx) error {
	type params struct {
		Id db.Id `param:"id"`
	}

	p := params{}
	if err := c.ParamsParser(&p); err != nil {
		return ErrInvalidUrlParams
	}

	group := db.NewUser{}
	if err := c.BodyParser(&group); err != nil {
		return ErrInvalidRequestBody
	}

	err := api.db.UpdateUser(p.Id, group)
	if err != nil {
		return err
	}

	return nil
}

func (api *API) DeleteUser(c *fiber.Ctx) error {
	type params struct {
		Id db.Id `param:"id"`
	}

	p := params{}
	if err := c.ParamsParser(&p); err != nil {
		return ErrInvalidUrlParams
	}

	err := api.db.DeleteUser(p.Id)
	if err != nil {
		return err
	}

	return nil
}

func (api *API) GetUser(c *fiber.Ctx) error {
	type params struct {
		Id db.Id `param:"id"`
	}

	p := params{}
	if err := c.ParamsParser(&p); err != nil {
		return ErrInvalidUrlParams
	}

	user, err := api.db.GetUser(p.Id)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"user": user})
}
