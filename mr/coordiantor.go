package mr

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"sync"
)

type TaskStatus uint

// 定义任务状态
const (
	Idle TaskStatus = iota + 1
	InProgress
	Completed
)

type Task struct {
	State    TaskStatus // 任务状态
	WorkerId string     // 任务所属worker
	In       any        // 任务输入
}

// Coordinator 负责调度任务
type Coordinator struct {
	In []any // 输入数据

	mu sync.RWMutex // 读写锁

	// 任务状态
	mapDone    bool
	reduceDone bool
}

// MakeCoordinator 创建一个Coordinator实例
// in 输入数据
// nReduce reduce任务数量
func MakeCoordinator(in []any, nReduce int) *Coordinator {
	c := Coordinator{
		In: in,
	}

	c.server() // 启动rpc服务
	return &c
}

// 判断任务是否完成
func (c *Coordinator) Done() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.mapDone && c.reduceDone
}

func (c *Coordinator) server() {
	rpc.Register(c) // 注册rpc服务
	rpc.HandleHTTP()

	// 生成一个UNIX 域套接字，用于worker和coordinator之间的通信
	sockname := coordinatorSock()
	os.Remove(sockname)
	l, err := net.Listen("unix", sockname)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	go http.Serve(l, nil)
	log.Println("[Coordinator] Start...")
}

// 定义rpc服务
// Ping
func (c *Coordinator) Ping(request *PingRequest, reply *PingReply) error {
	id := request.Id
	log.Printf("[Coordinator] Ping from worker %s\n", id)

	reply.Ok = true
	return nil
}
