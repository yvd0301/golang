package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Calc int // RPC 서버에 등록하기 위해 임의의 타입으로 정의

type Args struct { // 매개변수
	A,B int
}

type Reply struct { // 리턴값
	C int
}

func (c *Calc) Sum(args Args, reply *Reply) error {
	reply.C = args.A + args.B // 값을 더하여 리턴값 구조체에 넣어줌
	return nil
}


func main(){
	rpc.Register(new(Calc))
	ln, err := net.Listen("tcp", ":6000") // TCP 프로토콜에 6000번 포트로 연결을 받음
	if err != nil {
		fmt.Println(err)
		return
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept() // 클라이언트가 연결되면 TCP 연결을 리턴
		fmt.Println("connected")
		if err != nil {
			continue
		}
		defer conn.Close() // main 함수가 끝나기 직전에 TCP 연결을 닫음

		go rpc.ServeConn(conn) // RPC를 처리하는 함수를 고루틴으로 실행
	}
}
