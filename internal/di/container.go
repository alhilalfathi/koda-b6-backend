package di

import (
	"koda-b6-backend/internal/handlers"
	"koda-b6-backend/internal/repository"
	"koda-b6-backend/internal/service"

	"github.com/jackc/pgx/v5"
)

type Container struct {
	user        *pgx.Conn
	userRepo    *repository.UserRepository
	userService *service.UserService
	userHandler *handlers.UserHandler
}

func NewContainer(db *pgx.Conn) *Container {

	container := Container{
		user: db,
	}

	container.initDependencies()

	return &container
}

func (c *Container) initDependencies() {
	c.userRepo = repository.NewUserRepository(c.user)
	c.userService = service.NewUserService(c.userRepo)
	c.userHandler = handlers.NewUserHandler(c.userService)

}

func (c *Container) UserHandler() *handlers.UserHandler {
	return c.userHandler
}
