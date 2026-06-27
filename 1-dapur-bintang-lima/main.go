package main

import "fmt"

// ? Cara membuat struct atau blueprint gampangnya
type Restaurant struct {
	Name string;
	TotalChef int;
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
	r.TotalChef = r.TotalChef + tambahan
}

func (r Restaurant) TambahKokii(tambahan int) {
	r.TotalChef = r.TotalChef + tambahan
}

func main () {
	center := Restaurant{
		Name: "Baji Restaurant Pusat",
		TotalChef: 100,
		Is24Hour: true,
	}

	cirebonCity := Restaurant{
		Name: "Baji Restaurant Cirebon City",
		TotalChef: 40,
		Is24Hour: true,
	}
	
	fmt.Println("=== Welcome to Baji Restaurant ===")
	center.BukaToko()
	center.BuatPesanan("Nasi Goreng")
	fmt.Println(center.Name)
	fmt.Println(center.TotalChef)
	fmt.Println(center.Is24Hour)
	center.TambahKoki(20)
	// center.TambahKokii(20)
	fmt.Println(center.TotalChef)

	fmt.Println("")

	cirebonCity.BukaToko()
	cirebonCity.BuatPesanan("Nasi Lengko")
	fmt.Println(cirebonCity.Name)
	fmt.Println(cirebonCity.TotalChef)
	fmt.Println(cirebonCity.Is24Hour)
	cirebonCity.TambahKoki(20)
	// cirebonCity.TambahKokii(20)
	fmt.Println(cirebonCity.TotalChef)
}