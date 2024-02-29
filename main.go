package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	apiToken := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelList := []string{os.Getenv("CHANNEL_ID")}
	fileList := []string{"utility.pdf"}

	for i := 0; i < len(fileList); i++ {
		parameters := slack.FileUploadParameters{
			Channels: channelList,
			File:     fileList[i],
		}
		file, err := apiToken.UploadFile(parameters)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}
