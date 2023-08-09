package utils

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/msalemor/contoso-crm-common/pkg/models"
)

func GetSBClient(connectionString string) *azservicebus.Client {
	client, err := azservicebus.NewClientFromConnectionString(connectionString, nil)
	if err != nil {
		panic(err)
	}
	return client
}

func SendMessage(queue string, payload []byte, client *azservicebus.Client) error {
	sender, err := client.NewSender(queue, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	defer sender.Close(context.TODO())

	sbMessage := &azservicebus.Message{
		Body: payload,
	}
	err = sender.SendMessage(context.TODO(), sbMessage, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func SendEvent(queue, action, model string, modelPayload any, client *azservicebus.Client) {
	if client == nil {
		log.Println("utils.SendEvent: Unable to send message to ServiceBus")
		return
	}
	sender, err := client.NewSender(queue, nil)
	if err != nil {
		log.Println(err)
	}
	defer sender.Close(context.TODO())

	sbMessagePayload := models.EventMessage{
		Action:  action,
		Model:   model,
		Payload: modelPayload,
	}
	payload, _ := json.Marshal(sbMessagePayload)

	sbMessage := &azservicebus.Message{
		Body: payload,
	}
	err = sender.SendMessage(context.TODO(), sbMessage, nil)
	if err != nil {
		log.Println(err)
	}
}
