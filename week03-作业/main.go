package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func HttpServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "基于errgroup的httpserver.")
}

func StartHttpServer(s *http.Server) error {
	http.HandleFunc("/week03", HttpServer)
	err := s.ListenAndServe()
	return err
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	egroup, errCtx := errgroup.WithContext(ctx)
	srv := &http.Server{Addr: "0.0.0.0:8090"}

	//结束http服务器协程
	egroup.Go(func() error {
		<-errCtx.Done()
		fmt.Println("http server stop.")
		return srv.Shutdown(errCtx)
	})

	//启动http服务器协程
	egroup.Go(func() error {
		fmt.Println("http server start...")
		return StartHttpServer(srv)
	})

	//监听linux signal信号
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Listening signals...")

	egroup.Go(func() error {
		for {
			select {
			case <-errCtx.Done(): //
				return errCtx.Err()
			case s := <-channel: //接收到linux信号
				switch s {
				case syscall.SIGINT, syscall.SIGTERM:
					fmt.Printf("Recived signal: %v\n", s)
					cancel()
				default:
					fmt.Println("Unknow signal: ", s)
				}
			}
		}
	})

	if err := egroup.Wait(); err != nil {
		fmt.Printf("errgroup error: %+v\n", err.Error())
		return
	}
	fmt.Println("All group done")
}
