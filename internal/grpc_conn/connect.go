package grpc_conn

// import (
// 	"fmt"
// 	"strconv"

// 	"wuliangxiaoseng/openapi/internal/config"

// 	"google.golang.org/grpc"
// )

// var conn *grpc.ClientConn

// func setConn() {
// 	var err error
// 	port := strconv.Itoa(config.GrpcConfig.ServerPort)
// 	addr := ":" + port
// 	conn, err = grpc.Dial(addr, grpc.WithInsecure())
// 	if err != nil {
// 		fmt.Printf("连接服务端失败: %s", err)
// 		return
// 	}
// }

// func Getconn() *grpc.ClientConn {
// 	return conn
// }

// func Init() {
// 	setConn()
// }
