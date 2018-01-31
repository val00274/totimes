package main

import (
	"bufio"
	"bytes"
	"flag"
	"os"

	"github.com/nlopes/slack"
)

func opts() string {
	var message string
	flag.StringVar(&message, "m", "", "")
	flag.StringVar(&message, "message", "", "")
	flag.Parse()

	return message
}

func send(text string) {
	client := slack.New(os.Getenv("TOTIMES_TOKEN"))
	client.PostMessage(os.Getenv("TOTIMES_CHANNEL"), text, slack.PostMessageParameters{AsUser: true})
}

func readlines() string {
	var buf bytes.Buffer
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		text := stdin.Text()
		buf.WriteString(text)
		buf.WriteString("\n")
	}
	return buf.String()
}

func main() {
	message := opts()

	if len(message) > 0 {
		send(message)
	} else {
		send(readlines())
	}
}
