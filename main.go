package main

import (
	"crypto/tls"
	"github.com/thoj/go-ircevent"
	"log"
	"math/rand"
	"strings"
	"time"
)

func main() {
	SERVER := "irc.swehack.org:6697"
	NICK := "faggotbot"
	bannedwords := []string{"facebook", "tumblr", "haha aa", "bajsar", "skiter", "haha mm", "kek", "pepsi max", "SSD är värdelöst", "dkn ingen fv", "9gag", "flash xss", "jag bor i stockholm", "steam", "löfven", "åkesson", "deepweb", "norton", "antivirus", "ebin", "yolo", "jålå", "yålå", "swag", "sväg", "swäg"}
	responses := []string{"faggot detected", "nice meme", "episkt, helt enkelt episkt", "( ͡° ͜ʖ ͡°)", "(ノಠ益ಠ)ノ彡┻━┻ ", "varför kan jag inte suga alla dessa kukar samtidigt", "bajsar ner mig", "kräks på golvet", "skiter sönder", "helvete vad smart ich bin", "https://www.youtube.com/watch?v=8iO5-ic0Ug4", "https://i.imgur.com/O4gRAHW.gif", "https://i.imgur.com/Fo4eamA.jpg"}

	i := irc.IRC(NICK, NICK)
	i.Version = "666"
	i.UseTLS = true
	i.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := i.Connect(SERVER)
	if err != nil {
		log.Fatal(err)
	}
	rand.Seed(time.Now().UnixNano())
	i.AddCallback("PRIVMSG", func(e *irc.Event) {
		if e.Nick == i.GetNick() {
			return
		}
		s := false
		for _, val := range bannedwords {
			if strings.Contains(e.Message(), val) {
				s = true
			}
		}
		if s {
			i.Privmsg(e.Arguments[0], responses[rand.Intn(len(responses))])

		}
	})

	i.Loop()
}
