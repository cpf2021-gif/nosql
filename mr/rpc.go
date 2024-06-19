package mr

import (
	"os"
	"strconv"
)

// worker主动向coordiantor发送任务请求
// 定义rpc服务的request和reply参数

// 生成一个UNIX 域套接字，用于worker和coordinator之间的通信
func coordinatorSock() string {
	s := "/var/tmp/nosql-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}

// Ping 测试worker和coordinator之间的连接是否正常
type PingRequest struct {
	Id string
}

type PingReply struct {
	Ok bool
}
