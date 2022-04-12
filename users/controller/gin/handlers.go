package gin

import (
	users "github.com/qystaubai/crm/users"

	"github.com/gin-gonic/gin"
)

// UsersRoute ...
type UsersRoute struct {
	core users.UsersCore
}

// New constructor that returns an instance
// of UsersRoute
func New(core users.UsersCore) *UsersRoute {
	return &UsersRoute{
		core,
	}
}

func (ur *UsersRoute) AddRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/", ur.listAll())
}
