package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

// User используется для тестирования.
type User struct {
	Nick string
	Age  int `limit:"18"`
	Rate int `limit:"0,100"`
}

// Str2Int конвертирует строку в int.
func Str2Int(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}

// Validate проверяет min и max для int c тегом limit.
func Validate(obj interface{}) bool {
	vobj := reflect.ValueOf(obj)
	objType := vobj.Type() // получаем описание типа

	// перебираем все поля структуры
	for i := range objType.NumField() {
		field := objType.Field(i)

		// field.Tag имеет тип StructTag и содержит все теги i-го поля
		// выводим имя поля и значение тега `json`
		fmt.Println(field.Name, field.Tag.Get("limit"))
		// берём значение текущего поля и проверяем, что это int
		if val, ok := vobj.Field(i).Interface().(int); ok {

			// получаем тег limit
			limit, ok := objType.Field(i).Tag.Lookup("limit")
			if !ok {
				panic(fmt.Errorf("нет тега limit у %s", objType.Field(i).Name))
			}
			if strings.Contains(limit, ",") {
				minInt := Str2Int(strings.Split(limit, ",")[0])
				maxInt := Str2Int(strings.Split(limit, ",")[1])
				if val < minInt || val > maxInt {
					return false
				}
			} else {
				if val < Str2Int(limit) {
					return false
				}
			}
		}
	}
	return true
}

func TestValidate(t *testing.T) {
	var table = []struct {
		name string
		age  int
		rate int
		want bool
	}{
		{"admin", 20, 88, true},
		{"su", 45, 10, true},
		{"", 16, 5, false},
		{"usr", 24, -2, false},
		{"john", 18, 0, true},
		{"usr2", 30, 200, false},
	}
	for _, v := range table {
		if Validate(User{v.name, v.age, v.rate}) != v.want {
			t.Errorf("Не прошла проверку запись %s", v.name)
		}
	}
}
