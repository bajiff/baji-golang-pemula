package main

import "fmt"

type Restaurant struct {
	Name string
	TotalChef int
	Is24Hour bool
}

func (r Restaurant) BukaToko() {
	fmt.Println("Perhatian!!!, Restoran ", r.Name, "telah dibuka!!!")
}

func (r Restaurant) BuatPesanan(namaMenu string) {
	fmt.Println("Chef di ", r.Name, " sedang memasak ", namaMenu)
}



func main () {
	center := Restaurant{
		Name: "Baji Restaurant",
		TotalChef: 10,
		Is24Hour: false,
	}

	cirebonCity := Restaurant{
		Name: "Baji Restaurant Cirebon",
		TotalChef: 40,
		Is24Hour: true,
	}
	
	cirebonCity.BukaToko()
	cirebonCity.BuatPesanan("Seblak Cihuy")
	
	fmt.Println("====== Welcome To Baji Restaurant ======")
	fmt.Println("=== Pusat ===")
	fmt.Println("Data: ", center)
	fmt.Println("Nama: ", center.Name)
	fmt.Println("TotalChef: ", center.TotalChef)
	fmt.Println("Is 24 Hour?: ", center.Is24Hour)

	fmt.Println("=== Cirebon City ===")
	fmt.Println("Data: ", cirebonCity.Name)
	fmt.Println("Nama: ", cirebonCity.Name)
	fmt.Println("TotalChef: ", cirebonCity.TotalChef)
	fmt.Println("Is 24 Hour?: ", cirebonCity.Is24Hour)
}