package main

import (
	"context"
	"demo/model"
	"demo/pb"
	"errors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/get/userInfo", func(c *gin.Context) {
		userName := c.Query("name")
		if len(userName) == 0 {
			c.Error(errors.New("用户名为空"))
			return
		}
		conn, err := grpc.Dial("127.0.0.1:8972", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatal("GRPC CONNECT FAIL")
		}
		defer conn.Close()
		client := pb.NewGreeterClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		res, err := client.GetUserInfo(ctx, &pb.GetUserRequest{Username: userName})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		c.JSON(200, model.UserInfo{
			Id:       res.Id,
			UserName: res.Username,
			NickName: res.Nickname,
		})
	})
	r.Run()
}
