package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/opencars/edrmvs/pkg/domain/mocks"
	"github.com/opencars/edrmvs/pkg/domain/model"
	"github.com/opencars/edrmvs/pkg/domain/query"
)

func TestServer_FindByNumber(t *testing.T) {
	type args struct {
		name          string
		number        string
		registrations []model.Registration
		wantErr       error
		httpStatus    int
	}

	record := model.TestRegistration(t)

	tests := []args{
		{
			name:   "ok",
			number: "AA9359PC",
			registrations: []model.Registration{
				*record,
			},
			wantErr:    nil,
			httpStatus: http.StatusOK,
		},
		{
			name:          "bad_request",
			number:        "BLAH-BLAH",
			registrations: []model.Registration{},
			wantErr:       model.ErrBadNumber,
			httpStatus:    http.StatusBadRequest,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			query := query.ListByNumber{
				UserID:  "aeda406c-1f27-47e0-a1de-c0ec2a339206",
				TokenID: "befb8aeb-922e-45e0-b89a-c936200a455c",
				Number:  tests[i].number,
			}

			svc := mocks.NewMockCustomerService(ctrl)
			svc.EXPECT().ListByNumber(gomock.Any(), &query).Return(tests[i].registrations, tests[i].wantErr)

			url := fmt.Sprintf("/api/v1/registrations?number=%s", tests[i].number)

			req := httptest.NewRequest(http.MethodGet, url, nil)
			req.Header.Set(HeaderUserID, "aeda406c-1f27-47e0-a1de-c0ec2a339206")
			req.Header.Set(HeaderTokenID, "befb8aeb-922e-45e0-b89a-c936200a455c")
			req.Header.Set(HeaderTokenName, "test-token")

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
		registrations []model.Registration
		wantErr       error
		httpStatus    int
	}

	record := model.TestRegistration(t)

	tests := []args{
		{
			name: "ok",
			vin:  "5YJXCCE40GF010543",
			registrations: []model.Registration{
				*record,
			},
			wantErr:    nil,
			httpStatus: http.StatusOK,
		},
		{
			name:          "bad_request",
			vin:           "BLAH-BLAH",
			registrations: make([]model.Registration, 0),
			wantErr:       model.ErrBadVIN,
			httpStatus:    http.StatusBadRequest,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			query := query.ListByVIN{
				UserID:  "aeda406c-1f27-47e0-a1de-c0ec2a339206",
				TokenID: "befb8aeb-922e-45e0-b89a-c936200a455c",
				VIN:     tests[i].vin,
			}

			svc := mocks.NewMockCustomerService(ctrl)
			svc.EXPECT().ListByVIN(gomock.Any(), &query, false).Return(tests[i].registrations, tests[i].wantErr)

			url := fmt.Sprintf("/api/v1/registrations?vin=%s", tests[i].vin)

			req := httptest.NewRequest(http.MethodGet, url, nil)
			req.Header.Set(HeaderUserID, "aeda406c-1f27-47e0-a1de-c0ec2a339206")
			req.Header.Set(HeaderTokenID, "befb8aeb-922e-45e0-b89a-c936200a455c")
			req.Header.Set(HeaderTokenName, "test-token")

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
		registration *model.Registration
		wantErr      error
		httpStatus   int
	}

	record := model.TestRegistration(t)

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
			wantErr:      model.ErrBadCode,
			httpStatus:   http.StatusBadRequest,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			query := query.DetailsByCode{
				UserID:  "aeda406c-1f27-47e0-a1de-c0ec2a339206",
				TokenID: "befb8aeb-922e-45e0-b89a-c936200a455c",
				Code:    tests[i].code,
			}

			svc := mocks.NewMockCustomerService(ctrl)
			svc.EXPECT().DetailsByCode(gomock.Any(), &query).Return(tests[i].registration, tests[i].wantErr)

			url := fmt.Sprintf("/api/v1/registrations/%s", tests[i].code)

			req := httptest.NewRequest(http.MethodGet, url, nil)
			req.Header.Set(HeaderUserID, "aeda406c-1f27-47e0-a1de-c0ec2a339206")
			req.Header.Set(HeaderTokenID, "befb8aeb-922e-45e0-b89a-c936200a455c")
			req.Header.Set(HeaderTokenName, "test-token")

			rr := httptest.NewRecorder()

			srv := newServer(svc)
			srv.router.ServeHTTP(rr, req)

			assert.Equal(t, tests[i].httpStatus, rr.Code)
			assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
		})
	}
}
