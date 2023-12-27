package player

// HandlerRegister 注册handler服务
func (p *Player) HandlerRegister() {
	p.handlers["add_friend"] = p.AddFriend
	p.handlers["del_friend"] = p.DelFriend
	p.handlers["resolve_chatmsg"] = p.ResolveChatMsg
}