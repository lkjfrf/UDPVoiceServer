package main

import (
	"log"
	"sync"

	"github.com/lkjfrf/content"
	"github.com/lkjfrf/core"
)

func main() {
	core.GetLogManager().SetLogFile()
	core.GetNetworkCore().Init(":8005")
	content.GetContentManager().Init()
	log.Println("VoiceServer Start")

	mu := sync.Mutex{}
	mu.Lock()
	mu.Lock()
}
