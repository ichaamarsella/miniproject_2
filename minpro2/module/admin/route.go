package admin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"minpro2/middleware"
)

type RouterAdmin struct {
	AdminRequestHandler RequestHandlerAdmin
}

func NewRouter(
	dbCrud *gorm.DB,
) RouterAdmin {
	return RouterAdmin{AdminRequestHandler: NewAdminRequestHandler(
		dbCrud,
	)}
}

func (r RouterAdmin) Handle(routerVersion *gin.RouterGroup) {

	basepath := "/admin"
	admin := routerVersion.Group(basepath)

	// Apply the JWT middleware to protected routes
	// Apply authentication middleware to all admin routes

	admin.POST("/",
		r.AdminRequestHandler.CreateAdmin,
	)

	admin.POST("/login",
		r.AdminRequestHandler.LoginAdmin,
	)

	admin.Use(middleware.AuthMiddleware())

	admin.GET("/:id",
		r.AdminRequestHandler.GetAdminById,
	)

	admin.PUT("/:id",
		r.AdminRequestHandler.UpdateAdmin,
	)

	admin.DELETE("/:email",
		r.AdminRequestHandler.DeleteAdmin,
	)

}
