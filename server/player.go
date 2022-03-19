package main

type Player struct {
	name               string
	notificationStream chan Notification
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
