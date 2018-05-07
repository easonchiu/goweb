package service

import (
	"github.com/gin-gonic/gin"
	"web/db"
	"web/model"
	"gopkg.in/mgo.v2/bson"
)

func GetUserInfoById(id bson.ObjectId) (gin.H, error) {
	db, close := db.MuseDB()
	defer close()

	data := model.User{}

	err := db.C(model.UserCollection).FindId(id).One(&data)

	if err != nil {
		return nil, err
	}

	return gin.H{
		"data": data,
	}, nil
}

func GetUsersList(skip int, limit int) (gin.H, error) {
	db, close := db.MuseDB()
	defer close()

	data := make([]model.User, limit)

	err := db.C(model.UserCollection).Find(bson.M{}).Skip(skip).Limit(limit).All(&data)

	if err != nil {
		return nil, err
	}

	return gin.H{
		"list": data,
	}, nil
}