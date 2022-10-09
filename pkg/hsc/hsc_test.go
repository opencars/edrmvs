package hsc

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/opencars/edrmvs/pkg/config"
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
				fmt.Fprintln(w, registrationsData)
			},
		),
	)
	defer server.Close()

	api := New(&config.HSC{
		BaseURL:  server.URL,
		Username: "username",
		Password: "password",
	})

	arr, err := api.VehiclePassport(context.Background(), "", "АА9359РС")
	assert.NoError(t, err)
	assert.Equal(t, registrations, arr)
}

func TestMain(m *testing.M) {
	f, err := os.Open(registrationsPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to open golden file")
		os.Exit(1)
	}

	data, err := io.ReadAll(f)
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
