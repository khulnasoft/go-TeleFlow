package main

import (
	"context"
	"fmt"
	"log"

	teleflow "github.com/khulnasoft/go-teleflow/lib"
)

func main() {
	subscriberID := "<<REPLACE_WITH_YOUR_SUBSCRIBER>"
	apiKey := "<REPLACE_WITH_YOUR_API_KEY>"
	eventId := "<REPLACE_WITH_YOUR_EVENT_ID>"

	ctx := context.Background()
	to := map[string]interface{}{
		"lastName":     "Doe",
		"firstName":    "John",
		"subscriberId": subscriberID,
		"email":        "john@doemail.com",
	}

	payload := map[string]interface{}{
		"name": "Hello World",
		"organization": map[string]interface{}{
			"logo": "https://happycorp.com/logo.png",
		},
	}

	teleflowClient := teleflow.NewAPIClient(apiKey, &teleflow.Config{})

	// Trigger
	triggerResp, err := teleflowClient.EventApi.Trigger(ctx, eventId, teleflow.ITriggerPayloadOptions{
		To:      to,
		Payload: payload,
	})
	if err != nil {
		log.Fatal("Teleflow error", err.Error())
		return
	}

	fmt.Println(triggerResp)

	// Subscriber
	subscriber := teleflow.SubscriberPayload{
		LastName: "Skjæveland",
		Email:    "benedicte.skjaeveland@example.com",
		Avatar:   "https://randomuser.me/api/portraits/thumb/women/79.jpg",
		Data: map[string]interface{}{
			"location": map[string]interface{}{
				"city":     "Ballangen",
				"state":    "Aust-Agder",
				"country":  "Norway",
				"postcode": "7481",
			},
		},
	}

	resp, err := teleflowClient.SubscriberApi.Identify(ctx, subscriberID, subscriber)
	if err != nil {
		log.Fatal("Subscriber error: ", err.Error())
		return
	}

	fmt.Println(resp)

	// update subscriber
	updateSubscriber := teleflow.SubscriberPayload{FirstName: "Susan"}

	updateResp, err := teleflowClient.SubscriberApi.Update(ctx, subscriberID, updateSubscriber)
	if err != nil {
		log.Fatal("Update subscriber error: ", err.Error())
		return
	}

	fmt.Println(updateResp)

	// delete subscriber
	deleteResp, err := teleflowClient.SubscriberApi.Delete(ctx, subscriberID)
	if err != nil {
		log.Fatal("Update subscriber error: ", err.Error())
		return
	}
	fmt.Println(deleteResp)

	// get integrations
	integrations, err := teleflowClient.IntegrationsApi.GetAll(ctx)
	if err != nil {
		log.Fatal("Get all integrations error: ", err.Error())
	}
	fmt.Println(integrations)
}
