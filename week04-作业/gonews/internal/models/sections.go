package models

import (
	"app/internal/dao"
	"app/pkg/e"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/field"

	"go.mongodb.org/mongo-driver/bson"
)

type Part struct {
	Partid   string `bson:"partid" json:"partid"`
	Partname string `bson:"partname" json:"partname"`
}
type Site struct {
	Siteid   string `bson:"siteid" json:"siteid"`
	Sitename string `bson:"sitename" json:"sitename"`
	Part     []Part `bson:"parts" json:"parts"`
}
type Section struct {
	field.DefaultField `bson:",inline"`
	Sectionid          string `bson:"sectionid" json:"sectionid" binding:"required"`
	Sectionname        string `bson:"sectionname" json:"sectionname" binding:"required"`
	Site               []Site `bson:"sites" json:"sites"`
	Offset             int    `bson:"offset"`
}

type RetSection struct {
	Sectionid   string `bson:"sectionid" json:"id"`
	Sectionname string `bson:"sectionname" json:"name"`
}

type RetSite struct {
	Siteid   string `bson:"siteid" json:"id"`
	Sitename string `bson:"sitename" json:"name"`
}

func GetSectionList() (ret []RetSection, err error) {

	//var retM []bson.M
	var retSec []RetSection
	projectStage := bson.D{{"$project", []bson.E{{"sectionid", 1}, {"sectionname", 1}, {"_id", 0}}}}
	//matchStage := bson.D{{"$match", []bson.E{{"sectionid", "99001"}}}}
	stage := qmgo.Pipeline{projectStage}
	err = dao.Aggregate("sections", stage, &retSec)
	fmt.Printf("1111111\n")
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, e.NewError("地区不存在")
		}
		return nil, err
	}
	fmt.Printf("retSection:%v\n", retSec, retSec)

	return retSec, nil
}

func GetSiteList() (ret []RetSection, err error) {

	//var retM []bson.M
	var retSec []RetSection
	projectStage := bson.D{{"$project", []bson.E{{"sites.siteid", 1}, {"sites.sitename", 1}, {"_id", 0}}}}
	//matchStage := bson.D{{"$match", []bson.E{{"sectionid", "99001"}}}}
	stage := qmgo.Pipeline{projectStage}
	err = dao.Aggregate("sections", stage, &retSec)
	fmt.Printf("1111111\n")
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, e.NewError("地区不存在")
		}
		return nil, err
	}
	fmt.Printf("retSection:%v\n", retSec, retSec)

	return retSec, nil
}
