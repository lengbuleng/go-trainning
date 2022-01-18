package main

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
	analyse "github.com/hhxsv5/go-redis-memory-analysis"
)

var client redis.UniversalClient

const (
	ip   string = "127.0.0.1"
	port uint16 = 6379
	char string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

//初始化redis连接
func initClient(ctx context.Context) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", ip, port),
		Password: "",
		DB:       0,
	})
	// var cancel context.CancelFunc
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	_, err = client.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	ctx := context.Background()
	err := initClient(ctx)
	if err != nil {
		fmt.Printf("redis connect error,err:%s", err.Error())
		panic(err)
	}

	//1w数据量下
	write(ctx, 10000, "len10_1w", RandChar(10))
	write(ctx, 10000, "len1024_1w", RandChar(1024))
	write(ctx, 10000, "len5120_1w", RandChar(5120))

	//10w数据量下
	write(ctx, 100000, "len10_10w", RandChar(10))
	write(ctx, 100000, "len1024_10w", RandChar(1024))
	write(ctx, 100000, "len5120_10w", RandChar(5120))

	analysis()
}

func write(ctx context.Context, num int, key, value string) {
	for i := 0; i < num; i++ {
		k := fmt.Sprintf("%s:%v", key, i)
		cmd := client.Set(ctx, k, value, 0)
		err := cmd.Err()
		if err != nil {
			fmt.Printf("set (%s) error. cmd status:(%s). err:(%s)\n", k, cmd.String(), err)
		}
	}
}

//随机生成长度为size的字符串
func RandChar(size int) string {
	rand.NewSource(time.Now().UnixNano()) // 产生随机种子
	var s bytes.Buffer
	for i := 0; i < size; i++ {
		s.WriteByte(char[rand.Int63()%int64(len(char))])
	}
	return s.String()
}

func analysis() {
	analysis, err := analyse.NewAnalysisConnection(ip, port, "")
	defer analysis.Close()
	if err != nil {
		fmt.Println("something wrong:", err)
		return
	}

	analysis.Start([]string{":"})

	err = analysis.SaveReports("./reports")
	if err == nil {
		fmt.Println("done")
	} else {
		fmt.Println("error:", err)
	}
}
