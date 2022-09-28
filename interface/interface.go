package interfaces

import (
	"fmt"
)

func main() {
	var getuser = []map[string]interface{}{
		// {"<==get user==>": "<==get user==>"},
		{"name": "Nugraha Muthus", "alamat": "bekasi", "pekerjaan": "IT", "alasan": "hobi"},
		{"name": "Muhammad Hilmi", "alamat": "bekasi", "pekerjaan": "IT", "alasan": "hobi"},
		{"name": "Bayu Fajar Sidik", "alamat": "bekasi", "pekerjaan": "IT", "alasan": "hobi"},
		{"name": "I Putu Eka Widiantara", "alamat": "bekasi", "pekerjaan": "IT", "alasan": "hobi"},
		{"name": "Satrio Utomo", "alamat": "bekasi", "pekerjaan": "IT", "alasan": "hobi"},
		{"name": "Agus Supriyatna", "alamat": "bekasi", "pekerjaan": "IT", "alasan": "hobi"},
		{"name": "Barru Kurniawan", "alamat": "bekasi", "pekerjaan": "IT", "alasan": "hobi"},
		{"name": "HASANUDIN", "alamat": "bekasi", "pekerjaan": "IT", "alasan": "hobi"},
		{"name": "Yudha Nugraha", "alamat": "bekasi", "pekerjaan": "IT", "alasan": "hobi"},
		{"name": "Eka Muhendra", "alamat": "bekasi", "pekerjaan": "IT", "alasan": "hobi"},
	}
	for _, each := range getuser {
		fmt.Println("get", each["name"], each["alamat"], each["pekerjaan"], each["alasan"])
	}

	var register = []interface{}{
		"Satrio Utomo", "Eka Hendra", "Agus Supriatna",
	}
	var newreg = register[0:1]
	for _, each := range register {
		fmt.Println("berhasil di daftarkan", newreg, each)
	}
}
