package main

import "gamebackend/world"

func main() {
	world.MM = world.NewMgrMgr()
	world.MM.Pm.Run()
}