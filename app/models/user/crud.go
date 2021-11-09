package user

import (
	"jasy/goblog/pkg/logger"
	"jasy/goblog/pkg/model"
	"jasy/goblog/pkg/types"
)

func (user *User) Create() (err error) {
	if err = model.DB.Create(user).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

func Get(idstr string) (User, error) {
	var user User
	id := types.StringToUint64(idstr)
	if err := model.DB.First(&user, id).Error; err != nil{
		return user, err
	}
	return user, nil
}

func GetByEmail(email string) (User, error) {
	var user User
	if err := model.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}