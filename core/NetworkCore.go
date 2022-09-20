package core

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/lkjfrf/content"
)

type NetworkCore struct {
}

var instance *NetworkCore
var once sync.Once

func GetNetworkCore() *NetworkCore {
	once.Do(func() {
		instance = &NetworkCore{}
	})
	return instance
}

func (nc *NetworkCore) Init(port string) {
	log.Println("INIT_NetworkCore")
	nc.Connect(port)
}

func (nc *NetworkCore) ParseHeader(header []byte) (int, int) {
	pktsize := binary.LittleEndian.Uint16(header[:2])
	pktid := binary.LittleEndian.Uint16(header[2:4])

	return int(pktsize), int(pktid)
}

func (nc *NetworkCore) Connect(port string) {

	ServerAddr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("listening on ", port)

	conn, err := net.ListenUDP("udp", ServerAddr)
	if err != nil {
		log.Println("Connect Fail : ", err)
	} else {
		log.Println("Connect Success : ", conn)
	}

	go nc.Recv(conn)

}

func (nc *NetworkCore) Recv(conn *net.UDPConn) {

	for {
		data := make([]byte, 64*1024)

		n, addr, err := conn.ReadFromUDP(data)

		if n > 0 && err == nil {
			pktsize, pktid := nc.ParseHeader(data)
			//log.Println("RecvPacket : ", pktid, "/", pktsize)

			if pktsize > 4 {
				recv := data[4:pktsize]

				if content.GetContentManager().HandlerFunc[(int)(pktid)] != nil {
					content.GetContentManager().HandlerFunc[(int)(pktid)](conn, addr, string(recv))
				}
			}

			// if pktid == 62 { // 화면공유

			// 	datasize := pktsize - 6
			// 	if datasize > 0 {
			// 		recv := make([]byte, datasize)

			// 		//time.Sleep(time.Millisecond * 10)

			// 		_n, _, _err := conn.ReadFromUDP(recv)
			// 		if _err != nil {
			// 			log.Println(_err)
			// 		}

			// 		log.Println("RecvPacket : ", conn, " - ", "pktid : ", pktid)
			// 		if _n > 0 {
			// 			pktsize, pktid := nc.ParseHeader(recv)

			// 			if pktsize > 4 {
			// 				newHeader := make([]byte, 6)
			// 				pktsize := datasize + 6

			// 				binary.LittleEndian.PutUint32(newHeader, uint32(pktsize))
			// 				binary.LittleEndian.PutUint16(newHeader[4:], uint16(pktid))

			// 				result := append(newHeader, recv...)
			// 				//time.Sleep(time.Millisecond * 10)

			// 				content.GetSession().BroadCastToSameChannelExpetMeByte(conn, result, content.EScreenShare)
			// 			}
			// 		}
			// 	}
			//} else {
			// datasize := pktsize
			// if datasize > 0 {
			// 	recv := make([]byte, datasize)
			// 	_n, addr, _ := conn.ReadFromUDP(recv)
			// 	recv = recv[4:datasize]

			// 	if _n == datasize {
			// 		if content.GetContentManager().HandlerFunc[(int)(pktid)] != nil {
			// 			content.GetContentManager().HandlerFunc[(int)(pktid)](conn, addr, string(recv))
			// 		}
			// 		if pktid != 5 && pktid != 6 && pktid != 36 {
			// 			log.Println("RecvPacket : ", pktid)
			// 		}
			// 	} else {
			// 		log.Println("Packet Size Wrong!! : ", pktid, "-Read:", _n, "-Size:", datasize)
			// 	}
			// }
			//}
		} else {
			log.Println(err)
		}
	}
}
