package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// collection name
const UserCollection = "users"

// collection model
type User struct {
	Id			bson.ObjectId	`json:"id" bson:"_id"`
	NickName	string			`json:"nickname"`
	Email		string			`json:"email"`
	UserName	string			`json:"username"`
	Gid			string			`json:"gid"`
	Mobile		string			`json:"mobile"`
	Password	string			`json:"password"`
	Role		int				`json:"role"`
	CreateTime	time.Time		`json:"createTime"`
	UpdateTime	time.Time		`json:"updateTime"`
}