package state

import "neko.bar/nyactivity-server/models"

type State struct {
	users map[string]models.User
}

func New() State {
	return State{
		make(map[string]models.User),
	}
}

func (state *State) AddUser(*user models.User) {
	state.users[user.PrivateId] = user
}

func (state *State) RemoveUser(privateIdString string) {
	delete(state.users, name)
}
