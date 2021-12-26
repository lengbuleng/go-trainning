package models

import (
	"app/internal/dao"
	"app/pkg/e"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/qiniu/qmgo/field"

	"go.mongodb.org/mongo-driver/bson"
)

type Subscribe struct {
	Sectionid int       `bson:"sectionid" json:"sectionid"`
	Siteid    int       `bson:"siteid" json:"siteid"`
	Partid    int       `bson:"partid" json:"partid"`
	Status    bool      `bson:"status"`
	Subtime   time.Time `bson:"subtime"`
}

type User struct {
	field.DefaultField `bson:",inline"`
	Name               string      `bson:"username" json:"username" binding:"required"`
	Password           string      `bson:"userpwd" json:"password" binding:"required"`
	Subscribes         []Subscribe `bson:"subscribe" json:"subscribe"`
	Role               int         `bson:"role"`
}

func (user *User) Save() error {
	// user.CreateAt = time.Now().Local()

	// lastuser := &User{}
	// //查询最新createtime的news
	// err := dao.FindSort("users", bson.M{}, "-createAt", lastuser)
	// if err != nil {
	// 	if err == mongo.ErrNoDocuments {
	// 		user.Offset = 0
	// 		return dao.Insert("users", user)
	// 	}
	// 	return err
	// }
	// //如果存在该createtime，则offset+1
	// subtime := user.CreateAt.Unix() - lastuser.CreateAt.Unix()
	// if subtime < 1 {
	// 	user.Offset = lastuser.Offset + 1
	// } else {
	// 	user.Offset = 0
	// }

	return dao.Insert("users", user)
}

func (user *User) Login() error {
	err := dao.FindOne("users", bson.M{"username": user.Name, "userpwd": user.Password}, user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return e.NewError("用户不存在")
		}
		return err
	}
	if user.Id.IsZero() {
		return e.NewError("用户不存在")
	}
	return nil
}
