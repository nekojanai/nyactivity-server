package serverstate

import (
	"github.com/google/uuid"
	"neko.bar/nyactivity-server/models"
)

type ServerState struct {
	Users map[uuid.UUID]models.User
}

func New() *ServerState {
	return &ServerState{
		make(map[uuid.UUID]models.User),
	}
}

func (state *ServerState) AddUser(user *models.User) {
	state.Users[user.PrivateId] = *user
}

func (state *ServerState) RemoveUser(privateIdString uuid.UUID) {
	delete(state.Users, privateIdString)
}
