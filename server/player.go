package main

type Player struct {
	name               string
	notificationStream chan Notification
	role               uint8
}

func (p *Player) TakeNotification() (*Notification, error) {
	ntf, ok := <-p.notificationStream
	if ok {
		return &ntf, nil
	}
	return &ntf, errChannelClosed
}

func (p *Player) Notify(msg *Notification) {
	p.notificationStream <- *msg
}

func CreatePlayer(name string) *Player {
	return &Player{name: name, notificationStream: make(chan Notification, 10), role: PLAYER}
}
