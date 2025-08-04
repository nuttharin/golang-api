package service_user

import (
	"context"
	"golang-api/repositories/models"
	"golang-api/services/entities/request"

	"gorm.io/gorm"
)

type UserService interface {
	WithTx(txHandle *gorm.DB) UserService
	GetById(
		ctx context.Context,
		id uint,
	) (*models.User, error)
	List(
		ctx context.Context,
	) (*[]models.User, error)
	Create(ctx context.Context, data request.UserReq) (*models.User, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, id uint, data request.UserUpdateReq) error
}
