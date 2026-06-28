// ! main.go
package main

import (
	"dapur-bintang-lima/cabang"
	"fmt"
)

type Facilities interface {
	Tutup()
}

func main() {
	center := cabang.NewRestaurant("Baji Restaurant Pusat", 100, false)

	cirebonCity := cabang.NewRestaurant("Baji Restaurant Cirebon City", 70, true)

	bedulanCity := cabang.NewMenu("Nasi Rebus")

	fmt.Println("=== Welcome to Baji Restaurant ===")
	center.BukaToko()
	center.BuatPesanan("Nasi Goreng")
	fmt.Println(center.Name)
	fmt.Println(center.Is24Hour)
	fmt.Println("Total koki: ", center.GetTotalChef())
	center.TambahKoki(20)
	fmt.Println("Total Koki saat ini: ", center.GetTotalChef())

	fmt.Println("")

	cirebonCity.BukaToko()
	cirebonCity.BuatPesanan("Nasi Lengko")
	fmt.Println(cirebonCity.Name)
	fmt.Println(cirebonCity.Is24Hour)
	fmt.Println("Total koki: ", cirebonCity.GetTotalChef())
	cirebonCity.TambahKoki(50)
	fmt.Println("Total Koki saat ini: ", cirebonCity.GetTotalChef())

	fmt.Println("")

	fmt.Println("=== Welcome to Baji Menu ===")
	bedulanCity.BukaMenu()
	bedulanCity.BuatPesanan(12)
	fmt.Println(bedulanCity.Menu)
	fmt.Println(bedulanCity.IsSoldOut)
	fmt.Println("Total Menu: ", bedulanCity.GetTotalMenu())
	bedulanCity.TambahMenu(50)
	fmt.Println("Total Menu: ", bedulanCity.GetTotalMenu())
	fmt.Println("Total Worker saat ini: ", bedulanCity.GetTotalWorker())
	bedulanCity.AddWorker(10)
	fmt.Println("Total Worker saat ini: ", bedulanCity.GetTotalWorker())
	

	fmt.Println("\n ===== INSTRUKSI DARI BOS BESAR: TUTUP SEMUANYA =====")
	
	daftarFasilitas := []Facilities{center,cirebonCity,bedulanCity}
	
	for _, facility := range daftarFasilitas {
		facility.Tutup()
	}
}

