package handlers

import (
	"github.com/Valkyrie00/curriculum-telegram/internal/helpers"
	"github.com/Valkyrie00/curriculum-telegram/internal/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

var (
	homeReplyKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Raccontami chi sei"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Percorso lavorativo"),
			tgbotapi.NewKeyboardButton("Tecnologie e Progetti"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Contatti"),
		),
	)
)


func commandsHandler(update tgbotapi.Update) {
	command, _, ok := breakCommand(update.Message.Text)
	if ok {
		switch command {
		case "Torna":
			TornaCommand(update)
		case "Start":
			StartCommand(update)
		case "/Start":
			StartCommand(update)
		case "/start":
			StartCommand(update)
		case "Raccontami":
			StoryCommand(update)
		case "Percorso":
			JobsCommand(update)
		case "Tecnologie":
			TechCommand(update)
		case "Contatti":
			ContactsCommand(update)
		}
	}
}

func breakCommand(message string) (string, []string, bool) {
	var command []string
	var arguments []string
	if message == "" {
		return "", arguments, false
	}

	command = strings.Split(message, " ")
	if len(command) >= 2 {
		arguments = strings.Split(command[1], ",")
	}

	return command[0], arguments, true
}

// StartCommand - Command
func StartCommand(update tgbotapi.Update) {
	stories := []types.Message{
		{ChatID: update.Message.Chat.ID, Duration: 1, Content: "Ciao 🙂!", MsgType: "Message"},
		{ChatID: update.Message.Chat.ID, Duration: 2, Content: "Io sono Vito", MsgType: "Message"},
		{ChatID: update.Message.Chat.ID, Duration: 2, Content: "... o meglio la sua piccola copia digitale!", MsgType: "Message"},
		{ChatID: update.Message.Chat.ID, Duration: 1, Content: "Come posso aiutarti?", MsgType: "Message", Keyboard: &homeReplyKeyboard},
	}

	for _, story := range stories {
		helpers.ConsumeChainMessage(bot, story)
	}
}

// TornaCommand - Command
func TornaCommand(update tgbotapi.Update) {
	stories := []types.Message{
		{ChatID: update.Message.Chat.ID, Duration: 2, Content: "Cosa ti interessa sapere?", MsgType: "Message", Keyboard: &homeReplyKeyboard},
	}

	for _, story := range stories {
		helpers.ConsumeChainMessage(bot, story)
	}
}

// StoryCommand - Command
func StoryCommand(update tgbotapi.Update) {
	stories := []types.Message{
		{ChatID: update.Message.Chat.ID, Duration: 1, Content: "Ok!", MsgType: "Message"},
		{ChatID: update.Message.Chat.ID, Duration: 1, Content: "Questo sono io:", MsgType: "Message"},
		{ChatID: update.Message.Chat.ID, Duration: 2, Content: "assets/vito.jpg", MsgType: "NewPhotoUpload"},
		{ChatID: update.Message.Chat.ID, Duration: 1, Content: "Come dicevo mi chiamo *Vito Castellano*, ho *27 anni* e vivo a Legnano in provincia di Milano.", MsgType: "Message"},
		{ChatID: update.Message.Chat.ID, Duration: 2, Content: "Amo la palestra, il motociclismo ma soprattutto il *mio lavoro*!", MsgType: "Message"},
		{ChatID: update.Message.Chat.ID, Duration: 2, Content: "Attualmente lavoro presso Facile.it S.p.A dove ricopro il ruolo *Senior Backend Developer* e lavoro giornalmente con queste tecnologie:", MsgType: "Message"},
		// Message{ChatID: update.Message.Chat.ID, Duration: 2, Content: "Nel luogo di lavoro attuale mi ritrovo a lavorare quotidianamente con le seguenti tecnologie:", MsgType: "Message"},
		{ChatID: update.Message.Chat.ID, Duration: 2, Content: "- PHP\n- MySQL\n- Symfony (REST API)\n- RabbitMQ\n- Docker\n- Gitlab\n", MsgType: "Message"},
		{ChatID: update.Message.Chat.ID, Duration: 2, Content: "Prima di Symfony però ho conosciuto ed amato *Laravel* ❤️, di cui sono anche uno dei fondatori e attivisti dei *Meetup Laravel a Milano*.", MsgType: "Message"},
		{ChatID: update.Message.Chat.ID, Duration: 2, Content: "Questa è solo una parte delle mie conoscenze, per la lista e un dettaglio completo puoi usare il bottone *Tecnologie e Progetti* nel menù sotto.", MsgType: "Message"},
		{ChatID: update.Message.Chat.ID, Duration: 2, Content: "Nella sezione *Contatti* vi lascio il link al mio *GitHub* dove potete verificare la qualità del mio codice, come ad esempio questo Bot, senza perdere tempo a fare quei noiosi e inutili test. *#NoTest!*", MsgType: "Message"},
		{ChatID: update.Message.Chat.ID, Duration: 2, Content: "Se pensi che la mia figura possa essere utile al tuo progetto e se hai una proposta interessante, contattami pure!", MsgType: "Message"},
	}

	for _, story := range stories {
		helpers.ConsumeChainMessage(bot, story)
	}
}

// JobsCommand - Command
func JobsCommand(update tgbotapi.Update) {
	stories := []types.Message{
		{ChatID: update.Message.Chat.ID, Duration: 1, Content: "*Senior PHP Backend & Rest API Developer*\nGiugno 2018 - OGGI\n*Facile.it S.p.A*\n\nAgency; Sviluppo e mantenimento Web Application e servizi Rest API \n\nTecnologie usate/apprese: PHP - MySQL - Symfony - Docker - k8s - GitLab - Redis - Kibana - RabbitMQ", MsgType: "Message"},
		{ChatID: update.Message.Chat.ID, Duration: 1, Content: "*Senior PHP Backend & Rest API Developer*\nAprile 2015 - Giugno 2018\n*S2K Agency*\n\nAgency; Sviluppo e mantenimento Web Application e servizi Rest API \n\nTecnologie usate/apprese: PHP - MySQL - Laravel - Docker - Git - Redis - Deployer", MsgType: "Message"},
		{ChatID: update.Message.Chat.ID, Duration: 1, Content: "*Junior PHP Web Developer*\nMaggio 2014 - Marzo 2015\n*Pro Web Consulting*\n\nAgency; Sviluppo e mantenimento Web Application.\n\nTecnologie usate/apprese: PHP - MySQL - Laravel - Homestead - Git", MsgType: "Message"},
		{ChatID: update.Message.Chat.ID, Duration: 1, Content: "*Junior PHP Web Developer*\nFebbraio 2012 - Aprile 2014\n*Touring Club Italiano*\n\nSviluppo e mantenimento dei canali pubblici principali di Touring Club Italiano e Bandiere Arancioni.\n\nTecnologie usate/apprese: PHP - MySQL - CodeIgniter - Drupal", MsgType: "Message"},
		{ChatID: update.Message.Chat.ID, Duration: 1, Content: "*Tester Funzionale, PMO*\nOttobre 2011 - Febbraio 2012\n*NTT DATA Italia*\n\nMi occupavo principalmente di eseguire dei test funzionali su applicativi riguardanti la pubblicazione e gestione pubblicità a livello web, stampa e radio per il *GRUPPO SOLE 24 ORE*.", MsgType: "Message"},
	}

	for _, story := range stories {
		helpers.ConsumeChainMessage(bot, story)
	}
}

// TechCommand - Command
func TechCommand(update tgbotapi.Update) {
	stories := []types.Message{
		// Message{ChatID: update.Message.Chat.ID, Duration: 1, Content: "*Linguaggi*: \n\n -*PHP* ⭐️️️️⭐️⭐️⭐️⭐️  \n -*Go* ⭐️️️️⭐️⭐️⭐️ \n -*Python* ⭐️️️️⭐️⭐️\n -*C#* ⭐️️️️⭐️\n -*Rust* ⭐️️️️⭐️ ", MsgType: "Message"},
		// Message{ChatID: update.Message.Chat.ID, Duration: 1, Content: "*Database*: \n\n -*MySQL* ⭐️️️️⭐️⭐️⭐️⭐️  \n -*MongoDB* ⭐️️️️⭐️⭐️ \n", MsgType: "Message"},
		// Message{ChatID: update.Message.Chat.ID, Duration: 1, Content: "*Framework*: \n\n -*Symfony* ⭐️️️️⭐️⭐️⭐️⭐️  \n -*Laravel* ⭐️️️️⭐️⭐️⭐️⭐️  \n -*Codeigniter* ⭐️️️️⭐️⭐ \n -*Rocket* ⭐️️️️⭐️", MsgType: "Message"},
		// Message{ChatID: update.Message.Chat.ID, Duration: 1, Content: "*Cache*: \n\n -*Redis* ⭐️️️️⭐️⭐️⭐️ ", MsgType: "Message"},
		// Message{ChatID: update.Message.Chat.ID, Duration: 1, Content: "*Altro*: \n\n -*Docker* ⭐️️️️⭐️⭐️⭐️ \n -*RabbitMQ* ⭐️️️️⭐️⭐ \n -*k8s* ⭐️️️️", MsgType: "Message"},
		{ChatID: update.Message.Chat.ID, Duration: 1, Content: "*Linguaggi*: \n -*PHP* ⭐️️️️⭐️⭐️⭐️⭐️  \n -*Go* ⭐️️️️⭐️⭐️⭐️ \n -*Python* ⭐️️️️⭐️⭐️\n -*C#* ⭐️️️️⭐️\n -*Rust* ⭐️️️️⭐️ \n\n*Database*: \n -*MySQL* ⭐️️️️⭐️⭐️⭐️⭐️  \n -*MongoDB* ⭐️️️️⭐️⭐️ \n\n*Framework*: \n -*Symfony* ⭐️️️️⭐️⭐️⭐️⭐️  \n -*Laravel* ⭐️️️️⭐️⭐️⭐️⭐️  \n -*Codeigniter* ⭐️️️️⭐️⭐ \n -*Rocket* ⭐️️️️⭐️\n\n*Cache*: \n -*Redis* ⭐️️️️⭐️⭐️⭐️ \n\n*Altro*: \n -*Docker* ⭐️️️️⭐️⭐️⭐️ \n -*RabbitMQ* ⭐️️️️⭐️⭐ \n -*k8s* ⭐️️️️", MsgType: "Message"},
	}

	for _, story := range stories {
		helpers.ConsumeChainMessage(bot, story)
	}
}

// ContactsCommand - Command
func ContactsCommand(update tgbotapi.Update) {
	stories := []types.Message{
		{ChatID: update.Message.Chat.ID, Duration: 1, Content: "*Email*: castellano.vito@gmail.com \n*Linkedin*: https://www.linkedin.com/in/vito-castellano-27788953 \n*Github*: https://github.com/Valkyrie00", MsgType: "Message"},
	}

	for _, story := range stories {
		helpers.ConsumeChainMessage(bot, story)
	}
}