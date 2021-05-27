package controller

import (
	"github.com/gin-gonic/gin"
	"go-swagger-sample/models"
	"net/http"
)

// User godoc
// @Summary 查询用户列表
// @Description 查询用户列表描述
// @Tags 用户接口
// @Id /users
// @Accept json
// @Produce json
// @Param username query string false "登录名称"
// @Success 200 {object} models.Users "success"
// @Router /users [get]
func ListUsers(ctx *gin.Context) {
	user := &models.User{}
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	// TODO: list all users

	users := []models.User{*user}
	res := &models.Users{
		TotalCount: 1,
		Items:      users,
	}
	ctx.JSON(http.StatusOK, res)
}

// User godoc
// @Summary 创建用户
// @Description 创建用户描述
// @Tags 用户接口
// @Id /users
// @Accept json
// @Produce json
// @Param createBody body models.User true "创建用户请求主体"
// @Success 200 {object} models.User "success"
// @Router /users [post]
func CreateUser(ctx *gin.Context) {
	user := &models.User{}
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	// TODO: create a user

	ctx.JSON(http.StatusOK, user)
}

// User godoc
// @Summary 查询用户详情
// @Description 查询用户详情描述
// @Tags 用户接口
// @Id /users/:id
// @Accept json
// @Produce json
// @Param id path string true "用户标识"
// @Success 200 {object} models.User "success"
// @Router /users/{id} [get]
func GetUser(ctx *gin.Context) {
	user := &models.User{}
	if err := ctx.ShouldBindUri(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	// TODO: get a user

	ctx.JSON(http.StatusOK, user)
}
