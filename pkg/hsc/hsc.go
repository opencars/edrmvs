package hsc

import (
	"fmt"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigFastest

// Registration contains details of vehicle registration.
type Registration struct {
	Brand        *string `json:"brand"`
	Capacity     *string `json:"capacity"`
	Color        string  `json:"color"`
	DFirstReg    *string `json:"dFirstReg"`
	DReg         *string `json:"dReg"`
	Fuel         *string `json:"fuel"`
	Kind         *string `json:"kind"`
	MakeYear     string  `json:"makeYear"`
	Model        *string `json:"model"`
	NDoc         string  `json:"nDoc"`
	SDoc         string  `json:"sDoc"`
	NRegNew      string  `json:"nRegNew"`
	NSeating     *string `json:"nSeating"`
	NStanding    *string `json:"nStanding"`
	OwnWeight    *string `json:"ownWeight"`
	RankCategory *string `json:"rankCategory"`
	TotalWeight *string `json:"totalWeight"`
	VIN         *string `json:"vin"`
}

// API is wrapper to Head Service Center website.
type API struct {
	baseURL string
}

// New creates an instance of API wrapper.
func New(uri string) *API {
	api := new(API)

	api.baseURL = uri

	return api
}

// VehiclePassport sends GET request to Head Service Center.
// Code is identifier of vehicle registration certificate.
// Returns array of vehicles registration details.
func (api *API) VehiclePassport(code string) ([]Registration, error) {
	uri := fmt.Sprintf(
		"%s/gateway-edrmvs/api/verification/spr/%s/%s",
		api.baseURL, code[:3], code[3:],
	)

	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	info := make([]Registration, 0)
	if resp.StatusCode != http.StatusOK {
		return info, nil
	}

	err = json.NewDecoder(resp.Body).Decode(&info)
	resp.Body.Close()

	if err != nil {
		return nil, err
	}

	return info, nil
}
