package models

import (
	"app/internal/dao"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/qiniu/qmgo/field"

	"go.mongodb.org/mongo-driver/bson"
)

type Update struct {
	Updateid string `bson:"updateid" json:"updateid"`
	Url      string `bson:"url" json:"url"`
}
type News struct {
	field.DefaultField `bson:",inline"`
	Newid              string    `bson:"newid" json:"newid" binding:"required"`
	Title              string    `bson:"title" json:"title" binding:"required"`
	Body               string    `bson:"body" json:"body" binding:"required"`
	Writer             string    `bson:"writer" json:"writer"`
	Pagedate           time.Time `bson:"pagedate" json:"pagedate" binding:"required"`
	Downdate           time.Time `bson:"downdate" json:"downdate" binding:"required"`
	Url                string    `bson:"url" json:"url" binding:"required"`
	Mediaid            int       `bson:"mediaid" json:"mediaid"`
	Medianame          string    `bson:"medianame" json:"medianame"`
	Partid             int       `bson:"partid" json:"partid" binding:"required"`
	Partname           string    `bson:"partname" json:"partname" binding:"required"`
	Siteid             int       `bson:"siteid" json:"siteid" binding:"required"`
	Sitename           string    `bson:"sitename" json:"sitename" binding:"required"`
	Sectionid          int       `bson:"sectionid" json:"sectionid" binding:"required"`
	Sectionname        string    `bson:"sectionname" json:"sectionname" binding:"required"`
	Lang               string    `bson:"lang" json:"lang" binding:"required"`
	Update             []Update  `bson:"update" json:"update"`
	Offset             int       `bson:"offset"`
}

func (news *News) Save() error {
	news.CreateAt = time.Now().Local()

	lastnews := &News{}
	//查询最新createtime的news
	err := dao.FindSort("news", bson.M{}, "-createAt", lastnews)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			news.Offset = 0
			return dao.Insert("news", news)
		}
		return err
	}
	//如果存在该createtime，则offset+1
	subtime := news.CreateAt.UnixNano() - lastnews.CreateAt.UnixNano()
	if subtime < 1 {
		news.Offset = lastnews.Offset + 1
	} else {
		news.Offset = 0
	}

	return dao.Insert("news", news)
}
