package client

import (
	"encoding/json"
)

type EventQueue struct {
	Label   string `json:"label"`
	Type    string `json:"type"`
	PushUrl string `json:"push-url,omitempty"`
}

func (a *Api) EventQueuesList(factory string) ([]EventQueue, error) {
	url := a.serverUrl + "/ota/factories/" + factory + "/event-queues/"
	body, err := a.Get(url)
	if err != nil {
		return nil, err
	}
	var queues []EventQueue
	err = json.Unmarshal(*body, &queues)
	return queues, err
}

func (a *Api) EventQueuesDelete(factory, label string) error {
	url := a.serverUrl + "/ota/factories/" + factory + "/event-queues/" + label + "/"
	_, err := a.Delete(url, []byte{})
	return err
}
