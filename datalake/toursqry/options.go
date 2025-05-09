package toursqry

const (
	DepotLoadingModeNone               = "none"
	DepotLoadingModeOwner              = "owner"
	DepotLoadingModeAll                = "all"
	DepotLoadingModeByTenant           = "byTenant"
	DistanceModeNone                   = "none"
	DistanceModeHaversine              = "haversine"
	DistanceModeReal                   = "real"
	HandoverStationLoadingModeNone     = "none"
	HandoverStationLoadingModeAll      = "all"
	HandoverStationLoadingModeByTenant = "byTenant"
)

type Options struct {
	DistanceMode                   string   `json:"distanceMode"`                       // "none", "haversine", "real"
	DepotLoadingMode               string   `json:"depotLoadingMode"`                   // "none", "owner", "all", "byTenant"
	HandoverStationLoadingMode     string   `json:"handoverStationLoadingMode"`         // "none", "all", "byTenant"
	DepotTenantIDs                 []string `json:"depotTenantIDs,omitempty"`           // only used if DepotLoadingMode == "byTenant"
	HandoverStationTenantIDs       []string `json:"handoverStationTenantIDs,omitempty"` // only used if HandoverStationLoadingMode == "byTenant"
	Shuffle                        bool     `json:"shuffle"`
	ShuffleSeed                    int64    `json:"shuffleSeed"`
	FilterToursWithoutInOrOutDepot bool     `json:"filterToursWithoutInOrOutDepot"` // if true, filter out tours without in or out depot
}
