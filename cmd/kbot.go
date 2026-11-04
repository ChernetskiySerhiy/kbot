package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v4"
)

// TeleToken is the Telegram bot token from the environment
var TeleToken = os.Getenv("TELE_TOKEN")

// kbotCmd represents the kbot command
var kbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"start"},
	Short:   "Start the Telegram bot",
	Long: `kbot is a Telegram bot written in Go.
It replies to text messages:
  /start - greeting
  hello  - bot answers with its version
  any other text - bot repeats it back`,
	Run: func(cmd *cobra.Command, args []string) {
		if TeleToken == "" {
			log.Fatal("TELE_TOKEN environment variable is not set")
		}

		fmt.Printf("kbot %s started\n", appVersion)

		kbot, err := telebot.NewBot(telebot.Settings{
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})
		if err != nil {
			log.Fatalf("Please check TELE_TOKEN env variable. %s", err)
		}

		kbot.Handle("/start", func(m telebot.Context) error {
			return m.Send(fmt.Sprintf("Hi! I'm Kbot %s. Send me hello or any text.", appVersion))
		})

		kbot.Handle(telebot.OnText, func(m telebot.Context) error {
			log.Printf("Received message: %s", m.Text())

			switch m.Text() {
			case "hello":
				return m.Send(fmt.Sprintf("Hello I'm Kbot %s!", appVersion))
			default:
				return m.Send("You wrote: " + m.Text())
			}
		})

		kbot.Start()
	},
}

func init() {
	rootCmd.AddCommand(kbotCmd)
}