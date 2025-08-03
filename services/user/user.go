package service_user

import (
	"context"
	"golang-api/repositories/models"
	repository_user "golang-api/repositories/user"
	"golang-api/services/entities/request"
	"log"

	"gorm.io/gorm"
)

type userSvc struct {
	userRepo repository_user.UserRepository
}

func (svc userSvc) WithTx(txHandle *gorm.DB) UserService {
	// svc.odmCheckoutSvc = svc.odmCheckoutSvc.WithTx(txHandle)
	// svc.nextdayCheckoutSvc = svc.nextdayCheckoutSvc.WithTx(txHandle)
	svc.userRepo = svc.userRepo.WithTx(txHandle)
	return svc
}

func (svc userSvc) GetById(
	ctx context.Context,
	id uint,
) (*models.User, error) {
	user, err := svc.userRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (svc userSvc) List(
	ctx context.Context,
) (*[]models.User, error) {
	users, err := svc.userRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (svc userSvc) Create(ctx context.Context, data request.UserReq) (*models.User, error) {
	log.Println("userData svc")

	// Build model
	userData := models.User{
		Name:  data.Name,
		Email: data.Email,
	}

	log.Println(userData)
	user, err := svc.userRepo.Create(ctx, &userData)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewUserService(userRepo repository_user.UserRepository) UserService {
	return &userSvc{userRepo: userRepo}
}
