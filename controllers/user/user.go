package controller_user

import (
	"golang-api/httpserver"
	"golang-api/pkg/utils"
	"golang-api/services/entities/request"

	service_user "golang-api/services/user"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type UserController interface {
	GetUser(ctx httpserver.Context)
	List(ctx httpserver.Context)
	CreateUser(ctx httpserver.Context)
}

type userCtrl struct {
	userSvc service_user.UserService
}

// @Summary      get user by id
// @Description  get user by id
// @Tags         user
// @Router       /v1/user/{id} [get]
// @Accept       json
// @Produce      json
// @Param        id  path  string  true  "id"
// @Success      200	{object}  models.User
func (c *userCtrl) GetUser(ctx httpserver.Context) {
	id, err := ctx.GetParamInt("id")
	if err != nil {
		httpserver.AttachError(ctx, err)
		return
	}

	log.Println(id)

	user, err := c.userSvc.GetById(ctx.GetRequestCtx(), uint(id))
	if err != nil {
		httpserver.AttachError(ctx, err)
		return
	}

	httpserver.Data(ctx, user)
}

// @Summary      get list user
// @Description  get list user
// @Tags         user
// @Router       /v1/user [get]
// @Accept       json
// @Produce      json
// @Success      200	{object}  []models.User
func (c *userCtrl) List(ctx httpserver.Context) {

	user, err := c.userSvc.List(ctx.GetRequestCtx())
	if err != nil {
		httpserver.AttachError(ctx, err)
		return
	}

	httpserver.Data(ctx, user)
}

// @Summary      create user
// @Description  create user
// @Tags         user
// @Router       /v1/user [post]
// @Accept       json
// @Produce      json
// @Param        request  body  request.UserReq  true "create user body"
// @Success      200 {object}  response.UserRes
func (c *userCtrl) CreateUser(ctx httpserver.Context) {
	var r request.UserReq

	if err := ctx.Bind(&r); err != nil {
		httpserver.AttachError(ctx, err)
		return
	}

	if err := utils.ValidateStruct(r); err != nil {
		httpserver.AttachError(ctx, err)
		return
	}

	txHandle, _ := ctx.Get("db_tx")
	data, err := c.userSvc.WithTx(txHandle.(*gorm.DB)).Create(ctx.GetRequestCtx(), r)
	if err != nil {
		httpserver.AttachError(ctx, err)
		return
	}

	httpserver.Success(ctx, &httpserver.SuccessResponse{
		Code: http.StatusOK,
		Data: data,
	})
}

func NewCancelController(userSvc service_user.UserService) UserController {
	return &userCtrl{userSvc: userSvc}
}
