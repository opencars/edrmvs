package sqlstore_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/opencars/edrmvs/pkg/domain"
	"github.com/opencars/edrmvs/pkg/store/sqlstore"
)

func TestRegistrationRepository_Create(t *testing.T) {
	s, teardown := sqlstore.TestDB(t, conf)
	defer teardown("registrations")

	registration := domain.TestRegistration(t)
	require.NoError(t, s.Create(context.Background(), registration))
}

func TestRegistrationRepository_FindByNumber(t *testing.T) {
	s, teardown := sqlstore.TestDB(t, conf)
	defer teardown("registrations")

	registration := domain.TestRegistration(t)
	require.NoError(t, s.Create(context.Background(), registration))

	actual, err := s.FindByNumber(context.Background(), registration.Number)
	require.NoError(t, err)
	assert.Len(t, actual, 1)
	assert.Equal(t, registration.Code, actual[0].Code)
	assert.Equal(t, registration.Number, actual[0].Number)
	assert.Equal(t, registration.VIN, actual[0].VIN)
}

func TestRegistrationRepository_FindByCode(t *testing.T) {
	s, teardown := sqlstore.TestDB(t, conf)
	defer teardown("registrations")

	registration := domain.TestRegistration(t)
	require.NoError(t, s.Create(context.Background(), registration))

	actual, err := s.FindByCode(context.Background(), registration.Code)
	require.NoError(t, err)
	assert.Equal(t, registration.Code, actual.Code)
	assert.Equal(t, registration.Number, actual.Number)
	assert.Equal(t, registration.VIN, actual.VIN)
}

func TestRegistrationRepository_FindByVIN(t *testing.T) {
	s, teardown := sqlstore.TestDB(t, conf)
	defer teardown("registrations")

	registration := domain.TestRegistration(t)
	require.NoError(t, s.Create(context.Background(), registration))

	actual, err := s.FindByVIN(context.Background(), *registration.VIN)
	require.NoError(t, err)
	assert.Len(t, actual, 1)
	assert.Equal(t, registration.Code, actual[0].Code)
	assert.Equal(t, registration.Number, actual[0].Number)
	assert.Equal(t, registration.VIN, actual[0].VIN)
}
