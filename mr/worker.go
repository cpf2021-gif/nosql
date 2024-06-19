package mr

import (
	"log"
	"net/rpc"
)

// 启动worker
func Worker() {
	id := NewId()
	Ping(id)
}

func Ping(id string) {
	args := PingRequest{id}
	reply := PingReply{}

	if ok := call("Coordinator.Ping", &args, &reply); ok {
		log.Printf("[Worker-%s] Ping success\n", id)
	} else {
		log.Printf("[Worker-%s] Ping failed\n", id)
	}
}

// call 调用rpc服务
func call(rpcname string, request any, reply any) bool {
	sockname := coordinatorSock()
	c, err := rpc.DialHTTP("unix", sockname)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer c.Close()

	err = c.Call(rpcname, request, reply)
	if err == nil {
		return true
	}

	// 当发生错误时, 返回false
	log.Println(err)
	return false
}
