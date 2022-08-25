package main

import "fmt"

type person struct{ name string }

func clousur() {
	var orang = []*person{
		{name: "Nugraha Muthus"},
		{name: "Muhammad Hilmi"},
		{name: "Bayu Fajar Sidik"},
		{name: "I Putu Eka Widiantara"},
		{name: "Satrio Utomo"},
		{name: "Agus Supriyatna"},
		{name: "Barru Kurniawan"},
		{name: "HASANUDIN"},
		{name: "Yudha Nugraha"},
		{name: "Eka Muhendra"},
	}

	clousure := func(persons []*person) {

		for _, v := range persons {
			fmt.Println(v.name)
		}

	}
	clousure(orang)
}

func main() {

	clousur()

}
