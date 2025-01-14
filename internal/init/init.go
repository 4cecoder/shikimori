// Init package is responsible for initializing and making bot accessible
package inits

import (
	"embed"
	"fmt"
	"log"
	"os"
	"os/signal"

	ds "github.com/bwmarrin/discordgo"
)

//go:embed token.txt
var tokenTxt embed.FS

// Fnction that returns bot's token from token.txt file as a string
func ReadBotToken() string {
	tokenInBytes, err := tokenTxt.ReadFile("token.txt")
	if err != nil {
		log.Println(err)
	}
	token := string(tokenInBytes)
	return token
}

// Function that takes token value and creates a new discord session
func InitializeBot(token string) *ds.Session {
	discordSession, err := ds.New("Bot" + token)
	discordSession.AddHandler(func(session *ds.Session, r *ds.Ready) {
		fmt.Println("Shikimori is ready for her job.")
	})
	if err != nil {
		log.Println(err)
	}
	return discordSession
}

// Function that takes session and starts it
func StartBot(discordSession *ds.Session) {
	err := discordSession.Open()
	if err != nil {
		log.Println(err)
	}
	defer discordSession.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Graceful shutdown")
}

func Init() {
	StartBot(InitializeBot(ReadBotToken()))
}
