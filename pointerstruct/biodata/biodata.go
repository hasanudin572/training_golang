package main

import (
	"fmt"
	"os"
	"strconv"
)

type person struct {
	name      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {

	var orang = []*person{
		{name: "Nugraha Muthus", alamat: "bekasi", pekerjaan: "IT", alasan: "hobi"},
		{name: "Muhammad Hilmi", alamat: "bekasi", pekerjaan: "IT", alasan: "hobi"},
		{name: "Bayu Fajar Sidik", alamat: "bekasi", pekerjaan: "IT", alasan: "hobi"},
		{name: "I Putu Eka Widiantara", alamat: "bekasi", pekerjaan: "IT", alasan: "hobi"},
		{name: "Satrio Utomo", alamat: "bekasi", pekerjaan: "IT", alasan: "hobi"},
		{name: "Agus Supriyatna", alamat: "bekasi", pekerjaan: "IT", alasan: "hobi"},
		{name: "Barru Kurniawan", alamat: "bekasi", pekerjaan: "IT", alasan: "hobi"},
		{name: "HASANUDIN", alamat: "bekasi", pekerjaan: "IT", alasan: "hobi"},
		{name: "Yudha Nugraha", alamat: "bekasi", pekerjaan: "IT", alasan: "hobi"},
		{name: "Eka Muhendra", alamat: "bekasi", pekerjaan: "IT", alasan: "hobi"},
	}

	var args = os.Args
	var params = args[1]
	intparams, _ := strconv.Atoi(params)

	fmt.Println(orang[intparams])

}
