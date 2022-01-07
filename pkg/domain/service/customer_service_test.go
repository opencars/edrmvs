package service_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/opencars/edrmvs/pkg/domain/mocks"
	"github.com/opencars/edrmvs/pkg/domain/model"
	"github.com/opencars/edrmvs/pkg/domain/query"
	"github.com/opencars/edrmvs/pkg/domain/service"
)

func TestCustomerService_ListByNumber(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []model.Registration{
		*model.TestRegistration(t),
	}

	store := mocks.NewMockRegistrationStore(ctrl)
	store.EXPECT().FindByNumber(gomock.Any(), expected[0].Number).Return(expected, nil)

	producer := mocks.NewMockProducer(ctrl)
	producer.EXPECT().Produce(gomock.Any(), gomock.Any()).Return(nil)

	svc := service.NewCustomerService(store, nil, producer)
	actual, err := svc.ListByNumber(context.Background(), &query.ListByNumber{Number: expected[0].Number})
	require.NoError(t, err)

	assert.Len(t, actual, 1)
	assert.Equal(t, expected[0].Code, actual[0].Code)
	assert.Equal(t, expected[0].Number, actual[0].Number)
	assert.Equal(t, expected[0].VIN, actual[0].VIN)
}

func TestCustomerService_ListByVIN(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []model.Registration{
		*model.TestRegistration(t),
	}

	store := mocks.NewMockRegistrationStore(ctrl)
	store.EXPECT().FindByVIN(gomock.Any(), *expected[0].VIN).Return(expected, nil)

	producer := mocks.NewMockProducer(ctrl)
	producer.EXPECT().Produce(gomock.Any(), gomock.Any()).Return(nil)

	svc := service.NewCustomerService(store, nil, producer)
	actual, err := svc.ListByVIN(context.Background(), &query.ListByVIN{VIN: *expected[0].VIN}, false)
	require.NoError(t, err)

	assert.Len(t, actual, 1)
	assert.Equal(t, expected[0].Code, actual[0].Code)
	assert.Equal(t, expected[0].Number, actual[0].Number)
	assert.Equal(t, expected[0].VIN, actual[0].VIN)
}

func TestCustomerService_DetailsByCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := model.TestRegistration(t)

	store := mocks.NewMockRegistrationStore(ctrl)
	store.EXPECT().FindByCode(gomock.Any(), expected.Code).Return(expected, nil)

	producer := mocks.NewMockProducer(ctrl)
	producer.EXPECT().Produce(gomock.Any(), gomock.Any()).Return(nil)

	svc := service.NewCustomerService(store, nil, producer)
	actual, err := svc.DetailsByCode(context.Background(), &query.DetailsByCode{Code: expected.Code})
	require.NoError(t, err)

	assert.Equal(t, expected.Code, actual.Code)
	assert.Equal(t, expected.Number, actual.Number)
	assert.Equal(t, expected.VIN, actual.VIN)
}
