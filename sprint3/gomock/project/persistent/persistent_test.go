package persistent

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
	"yandex-learning-project/sprint3/gomock/project/mocks"
)

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockStore(ctrl)

	// возвращаемая ошибка
	errEmptyKey := errors.New("Указан пустой ключ")

	m.EXPECT().Get("").Return(nil, errEmptyKey)

	_, err := Lookup(m, "")
	require.ErrorIs(t, err, errEmptyKey)
}
