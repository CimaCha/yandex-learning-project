package store

//go:generate mockgen -source=store.go -destination=mocks/mock_store.go -package=mocks project/store Store
type Store interface {
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
	Delete(key string) error
}
