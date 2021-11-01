# sendcloud-sdk-go

Fork of afosto/sendcloud-go

## Overview [![GoDoc](https://godoc.org/github.com/joekendal/sendcloud-sdk-go?status.svg)](https://godoc.org/github.com/joekendal/sendcloud-sdk-go)

This package currently supports:
- parcels
- labels
- methods
- addresses
- integrations

## Install

```
go get github.com/joekendal/sendcloud-sdk-go
```

## Examples

Some examples on how to use the client are found below.

To initialize the client:
```go
api := client.API{}
api.Init("api_key", "api_secret")
```

To list shipping methods:
```go
methods, err := api.Method.GetMethods()
if err != nil {
    log.Fatal(err)
}

for _, m := range methods {
    log.Print(*m)
}
```
Create a parcel:
```go
params := &sendcloud.ParcelParams{
	Name:             "Sendcloud-GO",
	CompanyName:      "Afosto SaaS BV",
	Street:           "Grondzijl",
	HouseNumber:      "16",
	City:             "Groningen",
	PostalCode:       "9731DG",
	PhoneNumber:      "0507119519",
	EmailAddress:     "peter@afosto.io",
	CountryCode:      "NL",
	IsLabelRequested: true,
	Method:           8,
	ExternalID:       uuid.New().String(),
}
parcel, err := api.Parcel.New(params)
```

resolve a service point id from postnl,dhl,dpd to a sendcloud service point id 
```go
spid, err := api.ServicePoint.GetServicePoint(servicepoint.Matcher{
    SPID:        "NL-972301",
    Carrier:     "dhl",
    Country:     "NL",
    PostalCode:  "9731AR",
    HouseNumber: "41",
    Latitude:    53.226994,
    Longitude:   6.608120,
})

fmt.PrintLn(spid) 
```
this service point id can then be placed on an parcel on creation
