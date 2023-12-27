package world

import "gamebackend/manager"

// 管理所有的manager
type MgrMgr struct {
	Pm manager.PlayerMgr
}

func NewMgrMgr() *MgrMgr {
	return &MgrMgr{
		Pm: manager.PlayerMgr{},
	}
}

var MM *MgrMgr