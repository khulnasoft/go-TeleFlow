# Teleflow's API v1 Go Library

Teleflow's API exposes the entire Teleflow features via a standardized programmatic interface. Please refer to the full [documentation](https://docs.teleflow.khulnasoft.com/docs/overview/introduction) to learn more.

## Installation & Usage

Install the package to your GoLang project.

```golang
go get github.com/khulnasoft/go-teleflow
```

## Getting Started

Please follow the [installation procedure](#installation--usage) and then run the following:

```golang
package main

import (
	"context"
	"fmt"
	teleflow "github.com/khulnasoft/go-teleflow/lib"
	"log"
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

	data := teleflow.ITriggerPayloadOptions{To: to, Payload: payload}
	teleflowClient := teleflow.NewAPIClient(apiKey, &teleflow.Config{})

	resp, err := teleflowClient.EventApi.Trigger(ctx, eventId, data)
	if err != nil {
		log.Fatal("teleflow error", err.Error())
		return
	}

	fmt.Println(resp)

	// get integrations
	integrations, err := teleflowClient.IntegrationsApi.GetAll(ctx)
	if err != nil {
		log.Fatal("Get all integrations error: ", err.Error())
	}
	fmt.Println(integrations)
}
```

**NOTE**
Check the `cmd` directory to see a sample implementation and test files to see sample tests

## Documentation for API Endpoints

Class | Method                                                                           | HTTP request                            | Description
------------ |----------------------------------------------------------------------------------|-----------------------------------------| -------------
*EventApi* | [**Trigger**](https://docs.teleflow.khulnasoft.com/platform/subscribers#removing-a-subscriber)   | **Post** /events/trigger                | Trigger
*EventApi* | [**TriggerBulk**](https://docs.teleflow.khulnasoft.com/api/trigger-event/)   | **Post** /v1/events/trigger/bulk               | Bulk trigger event
*EventApi* | [**BroadcastToAll**](https://docs.teleflow.khulnasoft.com/api/broadcast-event-to-all/)   | **Post** /v1/events/trigger/broadcast               | Broadcast event to all
*EventApi* | [**CancelTrigger**](https://docs.teleflow.khulnasoft.com/api/cancel-triggered-event/)   | **Delete** /v1/events/trigger/:transactionId                | Cancel triggered event
*SubscriberApi* | [**Get**](https://docs.teleflow.khulnasoft.com/api/get-subscriber/) | **Get** /subscribers/:subscriberId                 | Get a subscriber
*SubscriberApi* | [**Identify**](https://docs.teleflow.khulnasoft.com/platform/subscribers#creating-a-subscriber) | **Post** /subscribers                 | Create a subscriber
*SubscriberApi* | [**Update**](https://docs.teleflow.khulnasoft.com/platform/subscribers#updating-subscriber-data)     | **Put** /subscribers/:subscriberID    | Update subscriber data
*SubscriberApi* | [**Delete**](https://docs.teleflow.khulnasoft.com/platform/subscribers#removing-a-subscriber)     | **Delete** /subscribers/:subscriberID | Removing a subscriber
*SubscriberApi* | [**Get**](https://docs.teleflow.khulnasoft.com/api/get-a-notification-feed-for-a-particular-subscriber)     | **Get** /subscribers/:subscriberId/notifications/feed | Get a notification feed for a particular subscriber
*SubscriberApi* | [**Get**](https://docs.teleflow.khulnasoft.com/api/get-the-unseen-notification-count-for-subscribers-feed)     | **Get** /subscribers/:subscriberId/notifications/feed | Get the unseen notification count for subscribers feed
*SubscriberApi* | [**Post**](https://docs.teleflow.khulnasoft.com/api/mark-a-subscriber-feed-message-as-seen)     | **Post** /v1/subscribers/:subscriberId/messages/markAs | Mark a subscriber feed message as seen
*SubscriberApi* | [**Get**](https://docs.teleflow.khulnasoft.com/api/get-subscriber-preferences/)     | **Get** /subscribers/:subscriberId/preferences | Get subscriber preferences
*SubscriberApi* | [**Patch**](https://docs.teleflow.khulnasoft.com/api/update-subscriber-preference/)     | **Patch** /subscribers/:subscriberId/preferences/:templateId | Update subscriber preference
*TopicsApi* | [**Get**](https://docs.teleflow.khulnasoft.com/api/filter-topics/) | **Get** /topics | Get a list of topics
*TopicsApi* | [**Get**](https://docs.teleflow.khulnasoft.com/api/get-topic/) | **Get** /topics/:topicKey | Get a topic by its topic key
*TopicsApi* | [**Post**](https://docs.teleflow.khulnasoft.com/api/topic-creation/) | **Post** /topics | Create a topic
*TopicsApi* | [**Patch**](https://docs.teleflow.khulnasoft.com/api/rename-a-topic/) | **Patch** /topics/:topicKey | Rename a topic
*TopicsApi* | [**Delete**](https://docs.teleflow.khulnasoft.com/api/delete-topic/) | **Delete** /topics/:topicKey | Delete a topic
*TopicsApi* | [**Post**](https://docs.teleflow.khulnasoft.com/api/subscribers-addition/) | **Post** /topics/:topicKey/subscribers | Add subscribers to a topic by key
*TopicsApi* | [**Post**](https://docs.teleflow.khulnasoft.com/api/subscribers-removal/) | **Post** /topics/:topicKey/subscribers/removal |Remove subscribers from a topic
*IntegrationsApi* | [**Create**](https://docs.teleflow.khulnasoft.com/platform/integrations)                         | **Post** /integrations                  | Create an integration
*IntegrationsApi* | [**Update**](https://docs.teleflow.khulnasoft.com/platform/integrations)                         | **Put** /integrations/:integrationId    | Update an integration
*IntegrationsApi* | [**Delete**](https://docs.teleflow.khulnasoft.com/platform/integrations)                         | **Delete** /integrations/:integrationId | Delete an integration
*IntegrationsApi* | [**Get**](https://docs.teleflow.khulnasoft.com/platform/integrations)                            | **Get** /integrations                   | Get all integrations
*IntegrationsApi* | [**GetActive**](https://docs.teleflow.khulnasoft.com/platform/integrations)                      | **Get** /integrations/active            | Get all active integrations
_InboundParserApi_ | [**Get**](https://docs.teleflow.khulnasoft.com/platform/inbound-parse-webhook/) | **Get** /inbound-parse/mx/status | Validate the mx record setup for the inbound parse functionality

## Authorization (api-key)

- **Type**: API key
- **API key parameter name**: ApiKey
- **Location**: HTTP header

### For more information about these methods and their parameters, see the [API documentation](https://docs.teleflow.khulnasoft.com/api-reference/overview).  

## Support and Feedback

Be sure to visit the Teleflow official [documentation website](https://docs.teleflow.khulnasoft.com/docs) for additional information about our API.

If you find a bug, please post the issue on [Github](https://github.com/khulnasoft/go-teleflow/issues).

As always, if you need additional assistance, join our Discord us a note [here](https://discord.gg/teleflow).

## Contributors

Name |
------------ |
[Oyewole Samuel](https://github.com/samsoft00) |
[Dima Grossman](https://github.com/scopsy) |
