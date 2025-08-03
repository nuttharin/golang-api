package repository_user

import (
	"context"
	"golang-api/repositories/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	WithTx(txHandle *gorm.DB) UserRepository
	List(ctx context.Context) (m *[]models.User, err error)
	GetById(ctx context.Context, id uint) (m *models.User, err error)
	Create(ctx context.Context, data *models.User) (m *models.User, err error)
	DeleteById(ctx context.Context, id uint) error
	UpdateById(ctx context.Context, id uint, m *models.User) error
}

type userRepo struct {
	conn *gorm.DB
}

func (repo userRepo) WithTx(txHandle *gorm.DB) UserRepository {
	if txHandle == nil {
		return repo
	}
	repo.conn = txHandle
	return repo
}

func (repo userRepo) List(ctx context.Context) (m *[]models.User, err error) {

	if err := repo.conn.WithContext(ctx).Find(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (repo userRepo) GetById(ctx context.Context, id uint) (m *models.User, err error) {
	if err := repo.conn.WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (repo userRepo) Create(ctx context.Context, data *models.User) (m *models.User, err error) {
	if err := repo.conn.WithContext(ctx).Create(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (repo userRepo) DeleteById(ctx context.Context, id uint) error {

	// Check have data
	if _, err := repo.GetById(ctx, id); err != nil {
		return err
	}

	return repo.conn.WithContext(ctx).Where("id = ?", id).Delete(&models.User{}).Error
}

func (repo userRepo) UpdateById(ctx context.Context, id uint, m *models.User) error {
	if err := repo.conn.WithContext(ctx).Model(models.User{}).Where("id = ? ", id).Updates(&m).Error; err != nil {
		return err
	}

	return nil
}

func (repo userRepo) preload() *gorm.DB {
	// Set preload of data relation
	return repo.conn
}

func New(conn *gorm.DB) UserRepository {
	return &userRepo{conn}
}
