package client

import (
	"github.com/joekendal/sendcloud-sdk-go/integration"
	"github.com/joekendal/sendcloud-sdk-go/method"
	"github.com/joekendal/sendcloud-sdk-go/parcel"
	"github.com/joekendal/sendcloud-sdk-go/sender"
	"github.com/joekendal/sendcloud-sdk-go/servicepoint"
)

type API struct {
	Parcel       *parcel.Client
	Method       *method.Client
	Sender       *sender.Client
	ServicePoint *servicepoint.Client
	Integration  *integration.Client
}

//Initialize the client
func (a *API) Init(apiKey string, apiSecret string) {
	a.Parcel = parcel.New(apiKey, apiSecret)
	a.Method = method.New(apiKey, apiSecret)
	a.Sender = sender.New(apiKey, apiSecret)
	a.ServicePoint = servicepoint.New(apiKey, apiSecret)
	a.Integration = integration.New(apiKey, apiSecret)
}
