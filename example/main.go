package main

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/skyscrapr/openai-sdk-go/openai"
// )

func main() {
	// 	ftreq := openai.CreateFineTunesRequest{
	// 		TrainingFile: "file-QotXmqJbY8PPgDmjOcRMe6E4",
	// 		ValidationFile: "file-0J5DkJ7KWC1QOyBJtb5r66dj",
	// 	}
	// 	fineTune, err := client.FineTunes().CreateFineTune(&ftreq)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("finetune-id: %s - state: %s \n", fineTune.Id, fineTune.Status)
	// 	err = client.FineTunes().SubscribeFineTuneEvents(fineTune.Id, fineTuneEventHandler, eventErrorHandler)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	//     fmt.Println("Done!")
	// }

	// func eventErrorHandler(err error) error {
	// 	log.Printf("error : %s", err)
	// 	return err
	// }

	//	func fineTuneEventHandler(event *openai.SSEEvent) error {
	//		var fineTuneEvent openai.FineTuneEvent
	//		err := json.Unmarshal([]byte(event.Data), &fineTuneEvent)
	//		log.Printf("Event : %s", fineTuneEvent.Message)
	//		return err
}
