package hsc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/opencars/edrmvs/pkg/config"
)

const (
	username = "username"
	password = "password"
)

type RegistrationStatus struct {
	StatusID int    `json:"statusId"`
	Status   string `json:"status"`
}

type Registration struct {
	PaperID        json.Number        `json:"paperId"`
	CarID          json.Number        `json:"carId"`
	LicencePlate   string             `json:"licencePlate"`
	MakeYear       int                `json:"makeYear"`
	Brand          *string            `json:"brand"`
	Model          *string            `json:"model"`
	CommercialDesc string             `json:"commercialDesc"`
	Vin            string             `json:"vin"`
	TotalWeight    *string            `json:"totalWeight"`
	OwnWeight      *string            `json:"ownWeight"`
	Category       *string            `json:"category"`
	Capacity       *string            `json:"capacity"`
	Fuel           *string            `json:"fuel"`
	Color          string             `json:"color"`
	Note           string             `json:"note"`
	Series         string             `json:"seria"`
	Number         string             `json:"number"`
	Status         RegistrationStatus `json:"status"`
	Date           *string            `json:"dreg"`
	FirstDate      *string            `json:"dfirstReg"`
	EndDate        interface{}        `json:"dend"`
	NChassis       *string            `json:"nchassis"`
	NSeating       *string            `json:"nseating"`
	NStanding      *string            `json:"nstandup"`
}

type Session struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	JTI          string `json:"jti"`
}

// API is wrapper to Head Service Center website.
type API struct {
	client  *http.Client
	baseURL string

	username string
	password string
}

// New creates an instance of API wrapper.
func New(conf *config.HSC) *API {
	return &API{
		client:   &http.Client{},
		baseURL:  conf.BaseURL,
		username: conf.Username,
		password: conf.Password,
	}
}

// VehiclePassport sends GET request to Head Service Center.
// Code is identifier of vehicle registration certificate.
// Returns array of vehicles registration details.
func (api *API) VehiclePassport(ctx context.Context, token string, code string) ([]Registration, error) {
	uri := fmt.Sprintf(
		"%s/sprlics-service/sprlics?seria=%s&number=%s",
		api.baseURL, code[:3], code[3:],
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := api.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	info := make([]Registration, 0)
	if resp.StatusCode != http.StatusOK {
		return info, nil
	}

	if err = json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, err
	}

	return info, nil
}

func (api *API) Authorize(ctx context.Context) (*Session, error) {
	uri := fmt.Sprintf(
		"%s/auth-server/oauth/token",
		api.baseURL,
	)

	data := make(url.Values)
	data.Set(username, api.username)
	data.Set(password, api.password)
	data.Set("grant_type", "password")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth("opendata_cl_id_nopd", "open_secret_encrypt_nopd")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := api.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("session: failed to authorize")
	}

	var session Session
	if err = json.NewDecoder(resp.Body).Decode(&session); err != nil {
		return nil, err
	}

	return &session, nil
}
