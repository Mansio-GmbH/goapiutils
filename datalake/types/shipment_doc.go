package types

import (
	"github.com/elliotchance/pie/v2"
	"github.com/mansio-gmbh/goapiutils/chrono"
	"github.com/mansio-gmbh/goapiutils/ct"
)

type ShipmentPosition struct {
	Position          string        `json:"position" mapstructure:"position"`
	Count             string        `json:"count" mapstructure:"count"`
	PackagingKind     string        `json:"packagingKind" mapstructure:"packagingKind"`
	GoodDescription   string        `json:"goodDescription" mapstructure:"goodDescription"`
	KindAndNumber     string        `json:"kindAndNumber" mapstructure:"kindAndNumber"`
	FreightableWeight *ct.UnitValue `json:"freightableWeight" mapstructure:"freightableWeight"`
	Weight            *ct.UnitValue `json:"weight" mapstructure:"weight"`
	Length            *ct.UnitValue `json:"length" mapstructure:"length"`
	Width             *ct.UnitValue `json:"width" mapstructure:"width"`
	Height            *ct.UnitValue `json:"height" mapstructure:"height"`
	Volume            *ct.UnitValue `json:"volume" mapstructure:"volume"`
	MonetaryValue     *ct.UnitValue `json:"monetaryValue" mapstructure:"monetaryValue"`
	LoadingMeter      *ct.UnitValue `json:"loadingMeter" mapstructure:"loadingMeter"`
	Note              string        `json:"note" mapstructure:"note"`
	PalletSpace       int           `json:"palletSpace" mapstructure:"palletSpace"`
}

type CityLocation struct {
	PostalCode         string          `json:"postalCode" mapstructure:"postalCode"`
	City               string          `json:"city" mapstructure:"city"`
	CountryCode        string          `json:"countryCode" mapstructure:"countryCode"`
	CenterCoordinates  *ct.Coordinates `json:"centerCoordinates,omitempty" mapstructure:"centerCoordinates,omitempty"`
	SnappedCoordinates *ct.Coordinates `json:"snappedCoordinates,omitempty" mapstructure:"snappedCoordinates,omitempty"`
	SnappedDistance    *float64        `json:"snappedDistance,omitempty" mapstructure:"snappedDistance,omitempty"`
}

type ShipmentDoc struct {
	BaseDoc
	Number                string             `json:"number" mapstructure:"id"`
	TenantID              string             `json:"tenantID" mapstructure:"tenantID"`
	TourID                string             `json:"tourID" mapstructure:"tourID"`
	OrderNumber           *string            `json:"orderNumber,omitempty" mapstructure:"orderNumber,omitempty"`
	ReferenceNumber       *string            `json:"referenceNumber,omitempty" mapstructure:"referenceNumber,omitempty"`
	SenderLocation        *ct.Location       `json:"senderLocation,omitempty" mapstructure:"senderLocation,omitempty"`
	SenderCityLocation    *CityLocation      `json:"senderCityLocation,omitempty" mapstructure:"senderCityLocation,omitempty"`
	PickupLocation        *ct.Location       `json:"pickupLocation,omitempty" mapstructure:"pickupLocation,omitempty"`
	PickupCityLocation    *CityLocation      `json:"pickupCityLocation,omitempty" mapstructure:"pickupCityLocation,omitempty"`
	ConsigneeLocation     *ct.Location       `json:"consigneeLocation,omitempty" mapstructure:"consigneeLocation,omitempty"`
	ConsigneeCityLocation *CityLocation      `json:"consigneeCityLocation,omitempty" mapstructure:"consigneeCityLocation,omitempty"`
	DeliveryLocation      *ct.Location       `json:"deliveryLocation,omitempty" mapstructure:"deliveryLocation,omitempty"`
	DeliveryCityLocation  *CityLocation      `json:"deliveryCityLocation,omitempty" mapstructure:"deliveryCityLocation,omitempty"`
	PickupDate            *chrono.Date       `json:"pickupDate,omitempty" mapstructure:"pickupDate,omitempty"`
	DeliveryDate          *chrono.Date       `json:"deliveryDate,omitempty" mapstructure:"deliveryDate,omitempty"`
	Positions             []ShipmentPosition `json:"positions" mapstructure:"positions"`
	TransshipmentBan      *bool              `json:"transshipmentBan,omitempty" mapstructure:"transshipmentBan,omitempty"`
	NoPaletteSwap         *bool              `json:"noPaletteSwap,omitempty" mapstructure:"noPaletteSwap,omitempty"`
	StackableLoad         *bool              `json:"stackableLoad,omitempty" mapstructure:"stackableLoad,omitempty"`
	Errors                []string           `json:"errors,omitempty" mapstructure:"errors,omitempty"`
	ImportedAt            chrono.Time        `json:"importedAt,omitempty" mapstructure:"importedAt,omitempty"`
	ImportReference       *string            `json:"importReference,omitempty" mapstructure:"importReference,omitempty"`
	CaseID                *string            `json:"caseID,omitempty" mapstructure:"caseID,omitempty"`
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

func (s ShipmentDoc) AggregatedWeight() (*ct.UnitValue, error) {
	return ct.AddUnitValues(nil, pie.Map(s.Positions, func(position ShipmentPosition) *ct.UnitValue { return position.Weight })...)
}

func (s ShipmentDoc) AggregatedVolume() (*ct.UnitValue, error) {
	return ct.AddUnitValues(nil, pie.Map(s.Positions, func(position ShipmentPosition) *ct.UnitValue { return position.Volume })...)
}

func (s ShipmentDoc) AggregatedLoadingMeter() (*ct.UnitValue, error) {
	return ct.AddUnitValues(nil, pie.Map(s.Positions, func(position ShipmentPosition) *ct.UnitValue { return position.LoadingMeter })...)
}

func (s ShipmentDoc) AggregatedMonetaryValue() (*ct.UnitValue, error) {
	return ct.AddUnitValues(nil, pie.Map(s.Positions, func(position ShipmentPosition) *ct.UnitValue { return position.MonetaryValue })...)

}
