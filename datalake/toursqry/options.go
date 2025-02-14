package toursqry

const (
	DepotLoadingModeNone           = "none"
	DepotLoadingModeOwner          = "owner"
	DepotLoadingModeAll            = "all"
	DistanceModeNone               = "none"
	DistanceModeHaversine          = "haversine"
	DistanceModeReal               = "real"
	HandoverStationLoadingModeNone = "none"
	HandoverStationLoadingModeAll  = "all"
)

type Options struct {
	DepotLoadingMode           string `json:"depotLoadingMode"`           // "none", "owner", "all"
	DistanceMode               string `json:"distanceMode"`               // "none", "haversine", "real"
	HandoverStationLoadingMode string `json:"handoverStationLoadingMode"` // "none", "all"
	Shuffle                    bool   `json:"shuffle"`
	ShuffleSeed                int64  `json:"shuffleSeed"`
}
