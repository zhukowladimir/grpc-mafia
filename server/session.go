package main

type Session struct {
	name    string
	id      uint64
	players map[string]Player
	ready   int
	game    *Game
}

func createSession(name string, id uint64) *Session {
	return &Session{name: name, id: id, players: make(map[string]Player), ready: 0, game: nil}
}

func (s *Session) IsValidUsername(username string) (bool, error) {
	if username == "SYSTEM" || username == "" {
		return false, errOccupiedUsername
	}
	for name := range s.players {
		if name == username {
			return false, errOccupiedUsername
		}
	}
	return true, nil
}

func (s *Session) NotifyPlayers(msg *Notification) {
	for _, player := range s.players {
		player.Notify(msg)
	}
}

func (s *Session) TakePlayerNotification(username string) (*Notification, error) {
	player, ok := s.players[username]
	if ok {
		return player.TakeNotification()
	}
	return nil, errDisconnectedPlayer
}

func (s *Session) GetPlayersList() []string {
	list := make([]string, len(s.players))
	i := 0
	for k := range s.players {
		list[i] = k
		i++
	}
	return list
}

func (s *Session) CreateGame() {
	s.game = &Game{players: &s.players, isNight: false}
}

func (s *Session) StartGame() {
	s.CreateGame()
	s.game.StartPhase()
}
