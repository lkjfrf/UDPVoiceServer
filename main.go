package main

import (
	"log"
	"sync"

	"github.com/lkjfrf/content"
	"github.com/lkjfrf/core"
	"github.com/lkjfrf/setting"
)

func main() {
	setting.GetStManager().Init()
	setting.GetLogManager().SetLogFile()
	core.GetNetworkCore().Init(setting.St_Ins.Port)
	content.GetContentManager().Init()
	log.Println("VoiceServer Start")

	mu := sync.Mutex{}
	mu.Lock()
	mu.Lock()
}
