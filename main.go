package main

import (
	"log"

	"github.com/lkjfrf/content"
	"github.com/lkjfrf/core"
)

func main() {
	core.GetLogManager().SetLogFile()
	core.GetNetworkCore().Init(":8001")
	content.GetContentManager().Init()
	log.Println("VoiceServerStart")
	for {
	}
}
