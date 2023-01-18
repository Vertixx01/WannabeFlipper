package client

import (
	"os"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/bot/msg"
	"github.com/Tnze/go-mc/bot/playerlist"
	"github.com/Tnze/go-mc/bot/screen"
	"github.com/Tnze/go-mc/bot/world"
	"github.com/Vertixx01/wannabeflipper/utils"
	"github.com/tidwall/gjson"
)

var (
	client        *bot.Client
	player        *basic.Player
	playerList    *playerlist.PlayerList
	chatHandler   *msg.Manager
	worldManager  *world.World
	screenManager *screen.Manager
)

func Client() *bot.Client {
	config, _ := os.Open("config.json")
	defer config.Close()
	name := gjson.Get(utils.Read(config), "name").String()
	client = bot.NewClient()
	client.Auth.Name = name
	player = basic.NewPlayer(client, basic.DefaultSettings, basic.EventsListener{
		GameStart: onGameStart,
	})
	playerList = playerlist.New(client)
	chatHandler = msg.New(client, player, playerList, msg.EventsHandler{
		SystemChat:        onSystemMsg,
		PlayerChatMessage: onPlayerMsg,
		DisguisedChat:     onDisguisedMsg,
	})
	utils.NewLogger().Info.Println("Client created!")
	return client
}
