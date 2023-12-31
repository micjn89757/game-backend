package player

import (
	"gamebackend/common"
	"gamebackend/chat"
	"log"
)

type Handler func(any)

//AddFriend 添加好友
func (p *Player) AddFriend(data any) {
	fId := data.(uint64)
	if !common.CheckInNumberSlice(fId, p.FriendList){
		p.FriendList = append(p.FriendList, fId)
	}
}

//DelFriend 删除好友
func (p *Player) DelFriend(data any) {
	fId := data.(uint64)
	p.FriendList = common.DelEleInSlice(fId, p.FriendList)
}

//ResolveChatMsg 处理消息
func (p *Player) ResolveChatMsg(data any) {
	chatMsg := data.(chat.Msg)

	log.Println(chatMsg)
	//todo 接收到消息然后发送给客户端
}