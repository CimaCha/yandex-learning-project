// файл persistent/persistent.go

package persistent

import (
	"yandex-learning-project/sprint3/gomock/project/store"
)

func Lookup(s store.Store, key string) ([]byte, error) {
	// ...
	return s.Get(key)
}
