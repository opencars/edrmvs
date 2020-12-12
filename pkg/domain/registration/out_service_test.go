package registration_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/opencars/edrmvs/pkg/domain"
	"github.com/opencars/edrmvs/pkg/domain/mocks"
	"github.com/opencars/edrmvs/pkg/domain/registration"
)

func TestOutService_FindByNumber(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []domain.Registration{
		*domain.TestRegistration(t),
	}

	store := mocks.NewMockRegistrationStore(ctrl)
	store.EXPECT().FindByNumber(gomock.Any(), expected[0].Number).Return(expected, nil)

	svc := registration.NewRegistrationService(store)
	actual, err := svc.FindByNumber(context.Background(), expected[0].Number)
	require.NoError(t, err)

	assert.Len(t, actual, 1)
	assert.Equal(t, expected[0].Code, actual[0].Code)
	assert.Equal(t, expected[0].Number, actual[0].Number)
	assert.Equal(t, expected[0].VIN, actual[0].VIN)
}

func TestOutService_FindByVIN(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := []domain.Registration{
		*domain.TestRegistration(t),
	}

	store := mocks.NewMockRegistrationStore(ctrl)
	store.EXPECT().FindByVIN(gomock.Any(), expected[0].Number).Return(expected, nil)

	svc := registration.NewRegistrationService(store)
	actual, err := svc.FindByVIN(context.Background(), expected[0].Number)
	require.NoError(t, err)

	assert.Len(t, actual, 1)
	assert.Equal(t, expected[0].Code, actual[0].Code)
	assert.Equal(t, expected[0].Number, actual[0].Number)
	assert.Equal(t, expected[0].VIN, actual[0].VIN)
}

func TestOutService_FindByCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := domain.TestRegistration(t)

	store := mocks.NewMockRegistrationStore(ctrl)
	store.EXPECT().FindByCode(gomock.Any(), expected.Code).Return(expected, nil)

	svc := registration.NewRegistrationService(store)
	actual, err := svc.FindByCode(context.Background(), expected.Code)
	require.NoError(t, err)

	assert.Equal(t, expected.Code, actual.Code)
	assert.Equal(t, expected.Number, actual.Number)
	assert.Equal(t, expected.VIN, actual.VIN)
}
