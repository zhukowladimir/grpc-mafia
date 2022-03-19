package main

import (
	"errors"
)

const (
	NOTIFICATION_USER_JOIN uint8 = iota
	NOTIFICATION_USER_LEAVE
	NOTIFICATION_USER_ROLE
	NOTIFICATION_USER_GHOST
	NOTIFICATION_SESSION_START
	NOTIFICATION_SESSION_END
	NOTIFICATION_GAME_START
	NOTIFICATION_GAME_END
)

var (
	errOccupiedUsername   error = errors.New("this username is already occupied!\nEnter another one: ")
	errChannelClosed      error = errors.New("this notification channel was closed")
	errDisconnectedPlayer error = errors.New("this player was disconnected")
)

type Notification struct {
	ntype uint8
	msg   string
}
