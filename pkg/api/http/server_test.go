package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/opencars/edrmvs/pkg/domain"
	"github.com/opencars/edrmvs/pkg/domain/mocks"
)

func TestServer_FindByNumber(t *testing.T) {
	type args struct {
		name          string
		number        string
		registrations []domain.Registration
		wantErr       error
		httpStatus    int
	}

	record := domain.TestRegistration(t)

	tests := []args{
		{
			name:   "ok",
			number: "AA9359PC",
			registrations: []domain.Registration{
				*record,
			},
			wantErr:    nil,
			httpStatus: http.StatusOK,
		},
		{
			name:          "bad_request",
			number:        "BLAH-BLAH",
			registrations: []domain.Registration{},
			wantErr:       domain.ErrBadNumber,
			httpStatus:    http.StatusBadRequest,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			svc := mocks.NewMockRegistrationService(ctrl)
			svc.EXPECT().FindByNumber(gomock.Any(), tests[i].number).Return(tests[i].registrations, tests[i].wantErr)

			url := fmt.Sprintf("/api/v1/registrations?number=%s", tests[i].number)
			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			srv := newServer(svc)
			srv.router.ServeHTTP(rr, req)

			assert.Equal(t, tests[i].httpStatus, rr.Code)
			assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
		})
	}
}

func TestServer_FindByVIN(t *testing.T) {
	type args struct {
		name          string
		vin           string
		registrations []domain.Registration
		wantErr       error
		httpStatus    int
	}

	record := domain.TestRegistration(t)

	tests := []args{
		{
			name: "ok",
			vin:  "5YJXCCE40GF010543",
			registrations: []domain.Registration{
				*record,
			},
			wantErr:    nil,
			httpStatus: http.StatusOK,
		},
		{
			name:          "bad_request",
			vin:           "BLAH-BLAH",
			registrations: make([]domain.Registration, 0),
			wantErr:       domain.ErrBadVIN,
			httpStatus:    http.StatusBadRequest,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			svc := mocks.NewMockRegistrationService(ctrl)
			svc.EXPECT().FindByVIN(gomock.Any(), tests[i].vin).Return(tests[i].registrations, tests[i].wantErr)

			url := fmt.Sprintf("/api/v1/registrations?vin=%s", tests[i].vin)
			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			srv := newServer(svc)
			srv.router.ServeHTTP(rr, req)

			assert.Equal(t, tests[i].httpStatus, rr.Code)
			assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
		})
	}
}

func TestServer_FindByCode(t *testing.T) {
	type args struct {
		name         string
		code         string
		registration *domain.Registration
		wantErr      error
		httpStatus   int
	}

	record := domain.TestRegistration(t)

	tests := []args{
		{
			name:         "ok",
			code:         "CXH484154",
			registration: record,
			wantErr:      nil,
			httpStatus:   http.StatusOK,
		},
		{
			name:         "bad_request",
			code:         "BLAH-BLAH",
			registration: nil,
			wantErr:      domain.ErrBadCode,
			httpStatus:   http.StatusBadRequest,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			svc := mocks.NewMockRegistrationService(ctrl)
			svc.EXPECT().FindByCode(gomock.Any(), tests[i].code).Return(tests[i].registration, tests[i].wantErr)

			url := fmt.Sprintf("/api/v1/registrations/%s", tests[i].code)
			req := httptest.NewRequest(http.MethodGet, url, nil)
			rr := httptest.NewRecorder()

			srv := newServer(svc)
			srv.router.ServeHTTP(rr, req)

			assert.Equal(t, tests[i].httpStatus, rr.Code)
			assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
		})
	}
}
