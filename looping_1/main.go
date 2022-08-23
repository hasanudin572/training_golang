package main

import (
	"fmt"
)

func main() {

	nama := [10]string{"Nugraha Muthus", "Muhammad Hilmi", "Bayu Fajar Sidik", "I Putu Eka Widiantara", "Satrio Utomo", "Agus Supriyatna", "Barru Kurniawan", "HASANUDIN", "Yudha Nugraha", "Eka Muhendra"}

	for index, element := range nama {

		fmt.Println(index)
		fmt.Println(element)

	}

}
