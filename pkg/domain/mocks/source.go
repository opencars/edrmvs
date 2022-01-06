package mocks

//go:generate mockgen -destination=./provider.go -package=mocks github.com/opencars/edrmvs/pkg/domain RegistrationProvider
//go:generate mockgen -destination=./service.go -package=mocks github.com/opencars/edrmvs/pkg/domain CustomerService
//go:generate mockgen -destination=./store.go -package=mocks github.com/opencars/edrmvs/pkg/domain RegistrationStore
