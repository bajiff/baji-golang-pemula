// ! cabang/restaurant.go
package cabang

import "fmt"

// ? Cara membuat struct atau blueprint gampangnya
type Restaurant struct {
	Name string;
	totalChef int;
	Is24Hour bool;
}

// ? r itu adalah alias yang nantinya Restaurant itu jadi singkat yaitu r, dan ini tanpa * yang artinya
func (r Restaurant) BukaToko() {
	fmt.Println("Perhatian!, Restaurant: ", r.Name, " sudah dibuka")
}

func (r Restaurant) BuatPesanan(namaPesanan string) {
	fmt.Println("Ada pesanan di Restaurant ", r.Name, " pesanan: ", namaPesanan)
}

func (r *Restaurant) TambahKoki(tambahan int) {
	r.totalChef = r.totalChef + tambahan
}

func (r Restaurant) GetTotalChef() int {
	return r.totalChef
}