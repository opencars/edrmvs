package mockstore

//go:generate mockgen -destination=./store.go -package=mockstore github.com/opencars/edrmvs/pkg/domain FullRegistrationStore
