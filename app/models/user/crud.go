package user

import (
	"jasy/goblog/pkg/logger"
	"jasy/goblog/pkg/model"
)

func (user *User) Create() (err error) {
	if err = model.DB.Create(user).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}