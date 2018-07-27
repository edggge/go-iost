package main

import (
	"flag"

	"fmt"

	. "github.com/iost-official/Go-IOS-Protocol/network"
	"github.com/iost-official/Go-IOS-Protocol/network/discover"
)

func main() {
	mode := flag.String("mode", "public", "operation mode: private | public | committee")
	flag.Parse()
	fmt.Println("[WARNING] Running in " + *mode + " mode. ")
	NetMode = *mode

	bootnodeStart()
}

func initNetConf() *NetConifg {
	conf := &NetConifg{}
	conf.SetLogPath("/tmp")
	conf.SetNodeTablePath("/tmp")
	conf.SetListenAddr("0.0.0.0")
	return conf
}

func bootnodeStart() {
	node, err := discover.ParseNode("84a8ecbeeb6d3f676da1b261c35c7cd15ae17f32b659a6f5ce7be2d60f6c16f9@0.0.0.0:30304")
	if err != nil {
		fmt.Errorf("parse boot node got err:%v", err)
	}
	conf := initNetConf()
	conf.SetNodeID(string(node.ID))
	baseNet, err := NewBaseNetwork(conf)
	if err != nil {
		fmt.Println("NewBaseNetwork ", err)
		return
	}
	ch, err := baseNet.Listen(node.TCP)
	if err != nil {
		fmt.Println("Init ", err)
		return
	}
	fmt.Println("server starting", node.Addr())
	for {
		select {
		case <-ch:
		}
	}

}
