// ! main.go
package main

import (
	"dapur-bintang-lima/cabang"
	"fmt"
)

func main () {
	center := cabang.Restaurant{
		Name: "Baji Restaurant Pusat",
		Is24Hour: true,
	}

	cirebonCity := cabang.Restaurant{
		Name: "Baji Restaurant Cirebon City",
		Is24Hour: true,
	}
	
	fmt.Println("=== Welcome to Baji Restaurant ===")
	center.BukaToko()
	center.BuatPesanan("Nasi Goreng")
	fmt.Println(center.Name)
	fmt.Println(center.Is24Hour)
	fmt.Println(center.GetTotalChef())
	center.TambahKoki(20)
	fmt.Println("Total Koki saat ini: ", center.GetTotalChef())

	fmt.Println("")

	cirebonCity.BukaToko()
	cirebonCity.BuatPesanan("Nasi Lengko")
	fmt.Println(cirebonCity.Name)
	fmt.Println(cirebonCity.Is24Hour)
	fmt.Println(cirebonCity.GetTotalChef())
	cirebonCity.TambahKoki(50)
	fmt.Println("Total Koki saat ini: ",cirebonCity.GetTotalChef())
}