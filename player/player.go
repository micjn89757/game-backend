package player

import (
	"gamebackend/define"
)

type Player struct {
	UId        uint64
	FriendList []uint64 // 好友列表
	handlers map[string]Handler
	HandlerParamCh chan define.HandlerParam
}


//NewPlayer 创建新Player
func NewPlayer() *Player {
	p := &Player{
		UId: 0, // ?需要生成
		FriendList: make([]uint64, 0),
		handlers: make(map[string]Handler, 0),
	}

	p.HandlerRegister()
	return p
}



//Run 
func (p *Player) Run() {
	for {
		select{
		case handlerParam := <- p.HandlerParamCh:
			if fn, ok := p.handlers[handlerParam.HandlerKey]; ok {
				fn(handlerParam.Data)
			}
		}
	}
}


