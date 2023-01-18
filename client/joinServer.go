package client

import (
	"github.com/Vertixx01/wannabeflipper/utils"
)

func Join() {
	client := Client()
	err := client.JoinServer("localhost")
	if err != nil {
		utils.NewLogger().Error.Println("Failed to join server: ", err)
	}
	utils.NewLogger().Info.Println("Login successful!")
	for {
		client.HandleGame()
	}
}
