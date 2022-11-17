package services

import (
	"github.com/XecFardin/go_tamplate/model"
	"github.com/XecFardin/go_tamplate/utils"
)

func GetUser(userId int64) (*model.User, *utils.ApplicationError) {
	return model.GetUser(userId)
}
