package setting

import (
	"sync"
)

type SettingHandler struct {
	ServerType int // 0: 나, 1: 원효로1번서버, 2: 원효로2번서버
	Port       string
	NasPath    string
	CTSAddress string
	LogPath    string
}

var St_Ins *SettingHandler
var St_once sync.Once

func GetStManager() *SettingHandler {
	St_once.Do(func() {
		St_Ins = &SettingHandler{}
	})
	return St_Ins
}

func (st *SettingHandler) Init() {
	st.ServerType = 0 // 0: 나, 1: 원효로1번서버, 2: 원효로2번서버

	switch st.ServerType {
	case 0:
		st.Port = ":8009"
		st.LogPath = "E:\\DIPServerLog\\VoiceServer\\"
	case 1:
		st.Port = ":8000"
		st.LogPath = "D:\\DIPServerLog\\VoiceServer1\\"
	case 2:
		st.Port = ":8000"
		st.LogPath = "D:\\DIPServerLog\\VoiceServer2\\"
	}
}
