package manager


import (
	"gamebackend/player"
)

//PlayerMgr 维护在线玩家
type PlayerMgr struct {
	players map[uint64]player.Player 
	addPCh chan player.Player
}


//Add添加在线玩家
func (pm *PlayerMgr) Add(p player.Player) {
	pm.players[p.UId] = p 
	go p.Run() 
}

//Run 启动玩家管理
func (pm *PlayerMgr) Run() {
	for {
		select {
		case p := <- pm.addPCh:
			pm.Add(p)
		}
	}
}