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
	CHANNEL := "#laidback"
	bannedwords := []string{"java", "xD", "facebook", "tumblr", "haha aa", "bajsar", "skiter", "haha mm", "kek", "pepsi max", "SSD är värdelös", "dkn ingen fv", "9gag", "flash xss", "jag bor i stockholm", "steam", "löfven", "romson", "fridolin", "schyman", "kolbit", "hipster", "åkesson", "deepweb", "norton", "antivirus", "ebin", "yolo", "jålå", "yålå", "swag", "sväg", "swäg", "meme", "feminism", "enema", "mema", "bajsrunk"}
	responses := []string{"faggot detected", "nice meme", "episkt, helt enkelt episkt", "( ͡° ͜ʖ ͡°)", "(ノಠ益ಠ)ノ彡┻━┻ ", "varför kan jag inte suga alla dessa kukar samtidigt", "bajsar ner mig", "kräks på golvet", "skiter sönder", "helvete vad smart ich bin", "https://www.youtube.com/watch?v=8iO5-ic0Ug4", "https://i.imgur.com/O4gRAHW.gif", "https://i.imgur.com/Fo4eamA.jpg", "ok", "jag vet inte varför jag borde bry mig", "du suger", "Det du nyss sade bidrar inte till konversationen. Gör oss alla en tjänst och håll käften.", "tkr du e irriterande favä", "dkn ingen pv", "jag heter per och har problem", "ibland när mamma inte tittar tar jag tre skedar oboy", "ibland luktar mina pruttar jordnötssmör", "smr är tokiga", "😼", "rösta ja! på att döda alla ¬katter!!", "katter är bättre än alla levande raser", "hellre en back i hallen än ett hack i ballen favä", "jag tycker du ska tona ner ditt bajsbrevande", "politik är ganska dumt egentligen!! rösta på kungen", "jag önskar jag var ett moln", "jag identifierar mig ibland som en attackhelikopter och jag hade gillat om ni kallade mig Apache-kun", "thinkpads!! gentoo!! krossa systemet", "äppleprodukter är ofta inte värda vad de kostar", "tid är bara värdefull om du lämnar ditt rum och faktiskt får något gjort", "din kommentar irriterar mig och jag vet inte varför", "hjälp min dator funkar inte kna någon hjälpa", "vet någon var man köper w33d online", "ibland gråter jag i duschen och önskar att grannarna hör", "fhu~", "dkn inga berg", "dkn inte född som ett moln", "jag gillar stjärnor!! astronomi är coolt", "jag önskar jag hade ett flygplan", ">du kommer aldrig vara mannen som kör glassbilen", "sluta bajsbreva !! var snällare", "finns här några snälla barn? :)", "kom sitt i farbrors knä", "jag är egentligen en utomjording", "när internet går sönder blir jag ledsen", "*hick*", "grönt ansikte och blåa händer", "grattis på födelsedagen !!", "hets mot folkgrupp??? tror du menar häst mot folkgrupp hahahahHHAHAHA", "har du övervägt att acceptera herren jesus som din ledare här i livet", "du behöver jesus!!", "må allah vara med dig, inshallah", "Deus Vult!", "lite facism har ingen dött av", "magiska flickor som slår på varandra!! var snälla så får ni julklapp,ar", "Hitler hängde i IRC", "NSA inside", "dkn när red inte kan hålla käften", "dkn polaren har en finare 240"}

	i := irc.IRC(NICK, NICK)
	i.Version = "666"
	i.UseTLS = true
	i.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := i.Connect(SERVER)
	if err != nil {
		log.Fatal(err)
	}
	rand.Seed(time.Now().UnixNano())
	i.AddCallback("001", func(e *irc.Event) {
		i.Join(CHANNEL)
	})
	i.AddCallback("PRIVMSG", func(e *irc.Event) {
		if e.Nick == i.GetNick() {
			return
		}
		s := false
		for _, val := range bannedwords {
			if strings.Contains(strings.ToUpper(e.Message()), strings.ToUpper(val)) {
				s = true
			}
		}
		if s {
			i.Privmsg(e.Arguments[0], responses[rand.Intn(len(responses))])

		}
	})

	i.Loop()
}
