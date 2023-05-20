package server

import "github.com/kalogs-c/web-socrates/entities"

type GameServer struct {
	players map[*entities.Player]bool
	join    chan *entities.Player
	leave   chan *entities.Player
}

func NewGameServer() *GameServer {
	return &GameServer{
		players: make(map[*entities.Player]bool),
		join:    make(chan *entities.Player),
		leave:   make(chan *entities.Player),
	}
}

func (gs *GameServer) Handle() {
	for {
		select {
		case player := <-gs.join:
			gs.joinPlayer(player)
		case player := <-gs.leave:
			gs.removePlayer(player)
		}
	}
}

func (gs *GameServer) joinPlayer(player *entities.Player) {
	gs.players[player] = true
}

func (gs *GameServer) removePlayer(player *entities.Player) {
	if _, ok := gs.players[player]; ok {
		delete(gs.players, player)
	}
}
