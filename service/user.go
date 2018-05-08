package service

import (
	"github.com/gin-gonic/gin"
	"web/db"
	"web/model"
	"gopkg.in/mgo.v2/bson"
)

// Query user info by id.
func GetUserInfoById(id bson.ObjectId) (gin.H, error) {
	db, close, err := db.CloneDB()

	if err != nil {
		return nil, err
	} else {
		defer close()
	}

	data := model.User{}

	err = db.C(model.UserCollection).FindId(id).One(&data)

	if err != nil {
		return nil, err
	}

	return gin.H{
		"data": data,
	}, nil
}

// Query users list with skip and limit.
func GetUsersList(skip int, limit int) (gin.H, error) {
	db, close, err := db.CloneDB()

	if err != nil {
		return nil, err
	} else {
		defer close()
	}

	data := make([]model.User, limit)

	err = db.C(model.UserCollection).Find(bson.M{}).Skip(skip).Limit(limit).All(&data)

	if err != nil {
		return nil, err
	}

	return gin.H{
		"list": data,
	}, nil
}