package store

import (
	"github.com/ismdeep/digest"

	"github.com/ismdeep/go-rest-api-demo/internal/model"
)

type userStore struct {
}

var User *userStore

func (receiver *userStore) Count() (int64, error) {
	var cnt int64
	if err := db.Model(&model.User{}).Count(&cnt).Error; err != nil {
		return 0, err
	}

	return cnt, nil
}

func (receiver *userStore) GetAll() ([]model.User, error) {
	var users []model.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (receiver *userStore) Create(username string, password string) error {
	user := model.User{
		ID:       username,
		Digest:   digest.Generate(password),
		Nickname: username,
	}

	return db.Create(&user).Error
}

func (receiver *userStore) Delete(username string) error {
	return db.Where("id = ?", username).Delete(&model.User{}).Error
}
