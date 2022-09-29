package content

import (
	"encoding/json"
	"log"

	"net"
	"sync"
)

const (
	Error = iota
	DBSignin
	PlayerLogout
	ChannelEnter
	NearPlayerUpdate

	PlayerMove //5
	PlayerActionEvent
	OtherPlayerMove
	PlayerLogin
	OtherPlayerSpawnInfo

	OtherPlayerDestroyInfo //10
	OtherInfo
	Voice
	RoomUserList
	RoomListUpdate

	Permission //15
	KickFromRoom
	MicToggle
	NoticeWrite
	NoticeContent

	NoticeList //20
	NoticeDelete
	NoticeModify
	ChannelCreate
	ChannelDelete

	CalenderRequest //25
	ChannelWidgetInfo
	NormalChat
	PrivateChat
	NoticeChat

	Questions // 30
	Invite
	InviteUserList
	CostumeSet
	UpdateCostume

	OtherUpdateCostume // 35
	HeartBeat
	AllFriendList
	SearchAddFriendList
	SearchDeleteFriendList

	AddFriend // 40
	DeleteFriend
	RequestAddFriend
	RequestDeleteFriend
	SpawnAvatar

	SaveFile //45
	CancelQuestion
	ModifyIntroduce
	FileList
	AccpetQuestion

	QuestionList //50
	ESaveShareData
	EPlaySaveShareData
	EEnterFileComplete
	EUploadComplete

	EScreenDataControlling //55
	RecvFileStatus
	EChannelCreateAfterEnter
	ERemoveFile
	EPlaceFBX

	ETestPlayerLogin //60
	EPlaceLoopingMP4
	EScreenShare
	EDownloadPPTtoPDF
	EGroupChat

	EGroupCreate //65
	EGroupActive
	EGroupUserListUpdate
	ERequestGroupList
	ERequestGroupUserList

	ESaveGroupAlarm

	Max
)

type ContentManager struct {
	Channel     sync.Map
	HandlerFunc map[int]func(*net.UDPConn, *net.UDPAddr, string)
}

var CM_Ins *ContentManager
var CM_once sync.Once

func GetContentManager() *ContentManager {
	CM_once.Do(func() {
		CM_Ins = &ContentManager{}
	})
	return CM_Ins
}

func JsonStrToStruct[T any](jsonstr string) T {
	var data T
	json.Unmarshal([]byte(jsonstr), &data)
	return data
}

func (cm *ContentManager) Init() {
	cm.Channel = sync.Map{}
	cm.HandlerFunc = make(map[int]func(*net.UDPConn, *net.UDPAddr, string), 0)

	cm.HandlerFunc[ChannelEnter] = cm.ChannelEnter
	cm.HandlerFunc[EScreenShare] = cm.ScreenShare
}

func (cm *ContentManager) ChannelEnter(conn *net.UDPConn, addr *net.UDPAddr, jsonstr string) {
	type S_ChannelEnter struct {
		Id          string
		ChannelNum  int32
		ChannelType int32 // 0: Auditorium, 1: Convention, 2: VirtualOffice, 3: VirtualGallery, 4: Plaza
	}
	data := S_ChannelEnter{}
	json.Unmarshal([]byte(jsonstr), &data)
	log.Println(data.Id)
	//cm.Channel.Store(data.Id, &Player{conn: conn, Channel: data.ChannelNum})
	GetSession().NewPlayer(data.Id, conn, addr, data.ChannelNum)

	// type R_NormalChat struct {
	// 	Name    string
	// 	Team    string
	// 	Grade   string
	// 	Message string
	// }

	// packet := R_NormalChat{}
	// packet.Name = "test"
	// GetSession().SendPacketByConn(conn, addr, packet, NormalChat)
}

func (cm *ContentManager) ScreenShare(conn *net.UDPConn, addr *net.UDPAddr, jsonstr string) {
	type SR_ScreenShare struct {
		Id     string
		Data   []uint16
		Status int32
		Size   int32
		Width  int32
		Height int32
	}
	data := SR_ScreenShare{}
	json.Unmarshal([]byte(jsonstr), &data)

	//GetSession().BroadCastToSameChannelExpetMeByte(conn, []byte(jsonstr), EScreenShare)
	GetSession().BroadCastToSameChannelExpetMe(data.Id, data, EScreenShare)
}
