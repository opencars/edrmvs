package sqlstore_test

import (
	"github.com/opencars/edrmvs/pkg/model"
	"github.com/opencars/edrmvs/pkg/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegistrationRepository_Create(t *testing.T) {
	s, teardown := sqlstore.TestDB(t, conf)
	defer teardown("registrations")

	registration := model.TestRegistration(t)
	assert.NoError(t, s.Registration().Create(registration))
}

func TestRegistrationRepository_FindByNumber(t *testing.T) {
	s, teardown := sqlstore.TestDB(t, conf)
	defer teardown("registrations")

	registration := model.TestRegistration(t)
	assert.NoError(t, s.Registration().Create(registration))

	actual, err := s.Registration().FindByNumber(registration.Number)
	assert.NoError(t, err)
	assert.Len(t, actual,1)
	assert.Equal(t, registration.Code, actual[0].Code)
	assert.Equal(t, registration.Number, actual[0].Number)
	assert.Equal(t, registration.VIN, actual[0].VIN)

}

func TestRegistrationRepository_FindByCode(t *testing.T) {
	s, teardown := sqlstore.TestDB(t, conf)
	defer teardown("registrations")

	registration := model.TestRegistration(t)
	assert.NoError(t, s.Registration().Create(registration))

	actual, err := s.Registration().FindByCode(registration.Code)
	assert.NoError(t, err)
	assert.Equal(t, registration.Code, actual.Code)
	assert.Equal(t, registration.Number, actual.Number)
	assert.Equal(t, registration.VIN, actual.VIN)
}

func TestRegistrationRepository_FindByVIN(t *testing.T) {
	s, teardown := sqlstore.TestDB(t, conf)
	defer teardown("registrations")

	registration := model.TestRegistration(t)
	assert.NoError(t, s.Registration().Create(registration))

	actual, err := s.Registration().FindByVIN(*registration.VIN)
	assert.NoError(t, err)
	assert.Len(t, actual,1)
	assert.Equal(t, registration.Code, actual[0].Code)
	assert.Equal(t, registration.Number, actual[0].Number)
	assert.Equal(t, registration.VIN, actual[0].VIN)
}

func TestRegistrationRepository_GetLast(t *testing.T) {
	s, teardown := sqlstore.TestDB(t, conf)
	defer teardown("registrations")

	registration := model.TestRegistration(t)
	assert.NoError(t, s.Registration().Create(registration))

	actual, err := s.Registration().GetLast(registration.SDoc)
	assert.NoError(t, err)
	assert.Equal(t, registration.Code, actual.Code)
	assert.Equal(t, registration.Number, actual.Number)
	assert.Equal(t, registration.VIN, actual.VIN)
}
