package hsc

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	registrationsPath = "../../testdata/registrations.json"
	registrationsData string
	registrations     []Registration
)

func TestVehiclePassport(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, string(registrationsData))
			},
		),
	)
	defer server.Close()

	api := New(server.URL)

	arr, err := api.VehiclePassport("АА9359РС")
	assert.NoError(t, err)
	assert.Equal(t, registrations, arr)
}

func TestMain(m *testing.M) {
	f, err := os.Open(registrationsPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to open golden file")
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to read golden file")
		os.Exit(1)
	}

	registrationsData = string(data)
	if err := json.Unmarshal(data, &registrations); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to parse JSON")
		os.Exit(1)
	}

	status := m.Run()

	os.Exit(status)
}
