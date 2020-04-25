package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/opencars/edrmvs/pkg/model"
	"github.com/opencars/edrmvs/pkg/store/teststore"
)

func TestServer_FindByNumber(t *testing.T) {
	store := teststore.New()

	registration := model.TestRegistration(t)
	assert.NoError(t, store.Registration().Create(registration))

	req := httptest.NewRequest(http.MethodGet, "http://127.0.0.1/api/v1/registrations?number=AA9359PC", nil)
	rr := httptest.NewRecorder()

	srv := newServer(store)
	srv.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, 200)
	assert.Equal(t, rr.Header().Get("Content-Type"), "application/json")
}

func TestServer_FindByVIN(t *testing.T) {
	store := teststore.New()

	registration := model.TestRegistration(t)
	assert.NoError(t, store.Registration().Create(registration))

	req := httptest.NewRequest(http.MethodGet, "http://127.0.0.1/api/v1/registrations?vin=5YJXCCE40GF010543", nil)
	rr := httptest.NewRecorder()

	srv := newServer(store)
	srv.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, 200)
	assert.Equal(t, rr.Header().Get("Content-Type"), "application/json")
}

func TestServer_FindByCode(t *testing.T) {
	store := teststore.New()

	registration := model.TestRegistration(t)
	assert.NoError(t, store.Registration().Create(registration))

	req := httptest.NewRequest(http.MethodGet, "http://127.0.0.1/api/v1/registrations/CXH484154", nil)
	rr := httptest.NewRecorder()

	srv := newServer(store)
	srv.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, 200)
	assert.Equal(t, rr.Header().Get("Content-Type"), "application/json")
}

func TestServer_Compress(t *testing.T) {
	store := teststore.New()

	registration := model.TestRegistration(t)
	assert.NoError(t, store.Registration().Create(registration))

	req := httptest.NewRequest(http.MethodGet, "http://127.0.0.1/api/v1/registrations/CXH484154", nil)
	req.Header.Set("Accept-Encoding", "gzip")
	rr := httptest.NewRecorder()

	srv := newServer(store)
	srv.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, 200)
	assert.Equal(t, rr.Header().Get("Content-Type"), "application/json")
	assert.Equal(t, rr.Header().Get("Content-Encoding"), "gzip")
}
