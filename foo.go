package main

import (
	"github.com/Xjs/protoTest/a"
	"google.golang.org/grpc"
)

func main() {
	var ao a.AObject
	_ = ao
	_ = grpc.Dial
}
