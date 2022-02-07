package main

import (
	"fmt"

	"github.com/google/uuid"
	"neko.bar/nyactivity-server/models"
	"neko.bar/nyactivity-server/serverstate"
)

func main() {
	serverState := serverstate.New()
	user := models.NewUser(uuid.NewString())
	serverState.AddUser(user)
	fmt.Println(serverState)
}
