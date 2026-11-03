package best_practise

import (
	"encoding/json"
	"errors"
)

type Record interface {
	json.Marshaler
	json.Unmarshaler
}

// для типа type RawMessage []byte в пакете encoding/json
// декларированы оба метода выше,
// то есть *json.RawMessage реализует интерфейс Record

/* плохо */
//type Storage interface {
//	Insert(*json.RawMessage, ID)     // метод принимает конкретный тип
//	Get(ID) (Record, error)          // метод возвращает интерфейс
//}

type ID string

/* хорошо */
type Storage interface {
	Insert(Record, ID)                // метод принимает интерфейс
	Get(ID) (*json.RawMessage, error) // метод возвращает конкретный тип
}

type MapStore map[ID]*json.RawMessage

func (m MapStore) Insert(r Record, id ID) {
	// делаем приведение типов,
	// освобождая от этого вызывающего
	m[id] = r.(*json.RawMessage)
}
func (m MapStore) Get(id ID) (*json.RawMessage, error) {
	r, ok := m[id]
	// проверяем, есть ли запись в хранилище
	if !ok {
		return r, errors.New("not found")
	}
	return r, nil
}

// конструктор
func NewMapStore() MapStore {
	s := make(map[ID]*json.RawMessage)
	return s
}
