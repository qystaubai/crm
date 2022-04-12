package gin

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (ur *UsersRoute) listAll() func(*gin.Context) {
	return func(ctx *gin.Context) {

		usersList, err := ur.core.ListUsers()

		if err != nil {
			log.Printf("%+v", errors.WithStack(err))
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, usersList)
	}
}
