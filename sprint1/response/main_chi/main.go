package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
	"strings"
)

var cars = map[string]string{
	"id1": "Renault Logan",
	"id2": "Renault Duster",
	"id3": "BMW X6",
	"id4": "BMW M5",
	"id5": "VW Passat",
	"id6": "VW Jetta",
	"id7": "Audi A4",
	"id8": "Audi Q7",
}

// carsListFunc — вспомогательная функция для вывода всех машин.
func carsListFunc() []string {
	var list []string
	for _, c := range cars {
		list = append(list, c)
	}
	return list
}

// carFunc — вспомогательная функция для вывода определённой машины.
func carFunc(id string) string {
	if c, ok := cars[id]; ok {
		return c
	}
	return "unknown identifier " + id
}

// brandFunc — вспомогательная функция для вывода машин определенного бренда.
func brandFunc(brand string) []string {
	brandCars := make([]string, 0)
	for _, car := range cars {
		if strings.Contains(strings.ToLower(car), strings.ToLower(brand)) {
			fmt.Println("Found a match!")
			brandCars = append(brandCars, car)
		} else {
			fmt.Println("No match found.")
		}
	}
	return brandCars
}

// modelFunc — вспомогательная функция для вывода машин определенного бренда и модели.
func modelFunc(brand string, model string) []string {
	modelCars := make([]string, 0)
	for _, car := range cars {
		if strings.Contains(strings.ToLower(car), strings.ToLower(brand)) {
			if strings.Contains(strings.ToLower(car), strings.ToLower(model)) {
				fmt.Println("Found a match!")
				modelCars = append(modelCars, car)
			}
		} else {
			fmt.Println("No match found.")
		}
	}
	return modelCars
}

func carsHandle(rw http.ResponseWriter, r *http.Request) {
	carsList := carsListFunc()
	io.WriteString(rw, strings.Join(carsList, ", "))
}

func carHandle(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte(carFunc(chi.URLParam(r, "id"))))
}

func brandHandle(rw http.ResponseWriter, r *http.Request) {
	marshal, err := json.Marshal(brandFunc(chi.URLParam(r, "brand")))
	if err != nil {
		return
	}
	rw.Write(marshal)
}

func modelHandle(rw http.ResponseWriter, r *http.Request) {
	brand := chi.URLParam(r, "brand")
	model := chi.URLParam(r, "model")
	marshal, err := json.Marshal(modelFunc(brand, model))
	if err != nil {
		return
	}
	rw.Write(marshal)
}

func main() {
	r := chi.NewRouter()

	r.Route("/cars", func(r chi.Router) {
		r.Get("/", carsHandle) // GET /cars
		// Route можно вкладывать один в другой
		r.Route("/{brand}", func(r chi.Router) {
			r.Get("/", brandHandle)        // GET /cars/renault
			r.Get("/{model}", modelHandle) // GET /cars/renault/duster
		})
	})
	r.Get("/car/{id}", carHandle)
	// r передаётся как http.Handler
	log.Fatal(http.ListenAndServe(":8080", r))
}
