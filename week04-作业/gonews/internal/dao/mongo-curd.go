package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func Insert(table string, data interface{}) error {
	cli := getColl(table)
	_, err := cli.InsertOne(context.Background(), data)
	return err
}

func InsertMany(table string, data interface{}) error {
	cli := getColl(table)
	_, err := cli.InsertMany(context.Background(), data)
	return err
}

func Find(table string, where, data interface{}) error {
	cli := getColl(table)
	err := cli.Find(context.Background(), where).All(data)
	return err
}

func FindOne(table string, where, data interface{}) error {
	cli := getColl(table)
	err := cli.Find(context.Background(), where).One(data)
	return err
}

func FindSort(table string, where interface{}, sortfield string, data interface{}) error {
	cli := getColl(table)
	err := cli.Find(context.Background(), where).Sort(sortfield).Limit(1).One(data)
	return err
}

func FindByPage(table string, where, data interface{}, page, limit int) error {
	cli := getColl(table)
	err := cli.Find(context.Background(), where).Skip(int64((page - 1) * limit)).Limit(int64(limit)).All(data)
	return err
}

func Count(table string, where interface{}, count int) error {
	cli := getColl(table)
	ct, err := cli.Find(context.Background(), where).Count()
	count = int(ct)
	return err
}

func UpdateMany(table string, where, update interface{}) error {
	cli := getColl(table)
	_, err := cli.UpdateAll(context.Background(), where, update)
	return err
}

func UpdateOne(table string, where, update interface{}) error {
	cli := getColl(table)
	err := cli.UpdateOne(context.Background(), where, update)
	return err
}

func DelOne(table string, where interface{}) error {
	cli := getColl(table)
	err := cli.Remove(context.Background(), where)
	return err
}

func DelMany(table string, where interface{}) error {
	cli := getColl(table)
	_, err := cli.RemoveAll(context.Background(), where)
	return err
}

func DelById(table string, id interface{}) error {
	cli := getColl(table)
	err := cli.RemoveId(context.Background(), id)
	return err
}

func Aggregate(table string, args []bson.D, data interface{}) error {
	cli := getColl(table)

	err := cli.Aggregate(context.Background(), args).All(data)
	return err
}
