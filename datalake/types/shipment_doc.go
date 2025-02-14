package types

import (
	"github.com/mansio-gmbh/goapiutils/chrono"
	"github.com/mansio-gmbh/goapiutils/ct"
)

type ShipmentPosition struct {
	Position        string `json:"position" mapstructure:"position"`
	Count           string `json:"count" mapstructure:"count"`
	PackagingKind   string `json:"packagingKind" mapstructure:"packagingKind"`
	GoodDescription string `json:"goodDescription" mapstructure:"goodDescription"`
	Weight          *Unit  `json:"weight" mapstructure:"weight"`
	Length          *Unit  `json:"length" mapstructure:"length"`
	Width           *Unit  `json:"width" mapstructure:"width"`
	Height          *Unit  `json:"height" mapstructure:"height"`
	Volume          *Unit  `json:"volume" mapstructure:"volume"`
	MonetaryValue   *Unit  `json:"monetaryValue" mapstructure:"monetaryValue"`
	LoadingMeter    *Unit  `json:"loadingMeter" mapstructure:"loadingMeter"`
	Note            string `json:"note" mapstructure:"note"`
}

type ShipmentDoc struct {
	BaseDoc
	Number            string             `json:"number" mapstructure:"id"`
	TenantID          string             `json:"tenantID" mapstructure:"tenantID"`
	TourID            string             `json:"tourID" mapstructure:"tourID"`
	OrderNumber       *string            `json:"orderNumber,omitempty" mapstructure:"orderNumber,omitempty"`
	ReferenceNumber   *string            `json:"referenceNumber,omitempty" mapstructure:"referenceNumber,omitempty"`
	SenderLocation    *ct.Location       `json:"senderLocation,omitempty" mapstructure:"senderLocation,omitempty"`
	PickupLocation    *ct.Location       `json:"pickupLocation,omitempty" mapstructure:"pickupLocation,omitempty"`
	ConsigneeLocation *ct.Location       `json:"consigneeLocation,omitempty" mapstructure:"consigneeLocation,omitempty"`
	DeliveryLocation  *ct.Location       `json:"deliveryLocation,omitempty" mapstructure:"deliveryLocation,omitempty"`
	PickupDate        *chrono.Date       `json:"pickupDate,omitempty" mapstructure:"pickupDate,omitempty"`
	DeliveryDate      *chrono.Date       `json:"deliveryDate,omitempty" mapstructure:"deliveryDate,omitempty"`
	Positions         []ShipmentPosition `json:"positions" mapstructure:"positions"`
	TransshipmentBan  *bool              `json:"transshipmentBan,omitempty" mapstructure:"transshipmentBan,omitempty"`
	NoPaletteSwap     *bool              `json:"noPaletteSwap,omitempty" mapstructure:"noPaletteSwap,omitempty"`
	StackableLoad     *bool              `json:"stackableLoad,omitempty" mapstructure:"stackableLoad,omitempty"`
	Errors            []string           `json:"errors,omitempty" mapstructure:"errors,omitempty"`
}

func (s ShipmentDoc) PickupAt() *ct.Location {
	if s.PickupLocation != nil && !s.PickupLocation.IsEmpty() {
		return s.PickupLocation
	}
	return s.SenderLocation
}

func (s ShipmentDoc) DeliveryAt() *ct.Location {
	if s.DeliveryLocation != nil && !s.DeliveryLocation.IsEmpty() {
		return s.DeliveryLocation
	}
	return s.ConsigneeLocation
}

func (s ShipmentDoc) AggregatedWeight() *Unit {
	var weight *Unit
	for _, position := range s.Positions {
		weight = AddUnit(weight, position.Weight)
	}
	return weight
}

func (s ShipmentDoc) AggregatedVolume() *Unit {
	var volume *Unit
	for _, position := range s.Positions {
		volume = AddUnit(volume, position.Volume)
	}
	return volume
}

func (s ShipmentDoc) AggregatedLoadingMeter() *Unit {
	var loadingMeter *Unit
	for _, position := range s.Positions {
		loadingMeter = AddUnit(loadingMeter, position.LoadingMeter)
	}
	return loadingMeter
}

func (s ShipmentDoc) AggregatedMonetaryValue() *Unit {
	var monetaryValue *Unit
	for _, position := range s.Positions {
		monetaryValue = AddUnit(monetaryValue, position.MonetaryValue)
	}
	return monetaryValue
}
