package admin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"minpro2/dto"
	"minpro2/repositories"
	"net/http"
	"strconv"
)

type RequestHandlerAdmin struct {
	ctr ControllerAdmin
}

func NewAdminRequestHandler(
	dbCrud *gorm.DB,
) RequestHandlerAdmin {
	return RequestHandlerAdmin{
		ctr: controllerAdmin{
			adminUseCase: useCaseAdmin{
				adminRepo: repositories.NewAdmin(dbCrud),
			},
		}}
}

func (h RequestHandlerAdmin) CreateAdmin(c *gin.Context) {
	request := AdminParam{}
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	res, err := h.ctr.CreateAdmin(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerAdmin) GetAdminById(c *gin.Context) {
	request := AdminParam{}
	err := c.BindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	adminid, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorResponse())
		return
	}

	//batas tambahan fitur yang bisa diblock
	res, err := h.ctr.GetAdminById(uint(adminid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerAdmin) UpdateAdmin(c *gin.Context) {
	request := AdminParam{}
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	adminId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	res, err := h.ctr.UpdateAdmin(request, uint(adminId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerAdmin) DeleteAdmin(c *gin.Context) {
	email := c.Param("email")
	res, err := h.ctr.DeleteAdmin(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerAdmin) LoginAdmin(c *gin.Context) {

	request := LoginAdmin{}
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	result, err := h.ctr.LoginAdmin(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	c.JSON(http.StatusOK, result)
}
