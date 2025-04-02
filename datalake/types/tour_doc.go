package types

import (
	"github.com/mansio-gmbh/goapiutils/chrono"
	"github.com/mansio-gmbh/goapiutils/ct"
)

type TourDoc struct {
	BaseDoc
	Number                                string        `json:"number"`
	TenantID                              string        `json:"tenantID"`
	PickupCountryCodes                    []string      `json:"pickupCountryCodes"`
	DeliveryCountryCodes                  []string      `json:"deliveryCountryCodes"`
	AllCountryCodes                       []string      `json:"allCountryCodes"`
	PickupPostalCodes                     []string      `json:"pickupPostalCodes"`
	DeliveryPostalCodes                   []string      `json:"deliveryPostalCodes"`
	AllPostalCodes                        []string      `json:"allPostalCodes"`
	TourDateFrom                          *chrono.Date  `json:"tourDateFrom"`
	TourDateUntil                         *chrono.Date  `json:"tourDateUntil"`
	PickupTimeWindows                     []string      `json:"pickupTimeWindows"`
	DeliveryTimeWindows                   []string      `json:"deliveryTimeWindows"`
	ShipmentIDs                           []string      `json:"shipmentIDs"`
	ShipmentsPickupLocation               []ct.Location `json:"shipmentsPickupLocation"`
	ShipmentsDeliveryLocation             []ct.Location `json:"shipmentsDeliveryLocation"`
	ShipmentsDistanceMatrix               [][]float64   `json:"shipmentsDistanceMatrix"`
	ShipmentsDurationSecsMatrix           [][]int       `json:"shipmentsDurationSecsMatrix"`
	TotalShipmentPickupDistance           float64       `json:"totalShipmentPickupDistance"`
	TotalShipmentPickupDurationSecs       int           `json:"totalShipmentPickupDurationSecs"`
	TotalShipmentDeliveryDistance         float64       `json:"totalShipmentDeliveryDistance"`
	TotalShipmentDeliveryDurationSecs     int           `json:"totalShipmentDeliveryDurationSecs"`
	LastPickupToFirstDeliveryDistance     float64       `json:"lastPickupToFirstDeliveryDistance"`
	LastPickupToFirstDeliveryDurationSecs int           `json:"lastPickupToFirstDeliveryDurationSecs"`
	TotalShipmentsDistance                float64       `json:"totalShipmentsDistance"`
	TotalShipmentsDurationSecs            int           `json:"totalShipmentsDurationSecs"`
	TotalWeight                           *Unit         `json:"totalWeight"`
	TotalLoadingMeter                     *Unit         `json:"totalLoadingMeter"`
	TotalVolume                           *Unit         `json:"totalVolume"`
	TotalMonetaryValue                    *Unit         `json:"totalMonetaryValue"`
	ImportedAt                            chrono.Date   `json:"importedAt,omitempty" mapstructure:"importedAt,omitempty"`
	ImportReference                       *string       `json:"importReference,omitempty" mapstructure:"importReference,omitempty"`
	CaseID                                *string       `json:"caseID,omitempty" mapstructure:"caseID,omitempty"`
}
