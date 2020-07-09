package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"os"
)

func main() {
	//create a new Node with a Node number of 1
	node, e := snowflake.NewNode(1)
	if e != nil {
		println(e)
		os.Exit(1)
	}
	for i := 0; i<3;i++  {
		id := node.Generate()
		fmt.Println("id",id,"Node:",id.Node(),"step:",id.Step(),"time:",id.Time())

	}
}
