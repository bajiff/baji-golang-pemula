// ! cabang/restaurant.go
package cabang

import (
	"fmt"
	"sync"
	"time"
)

// ? Cara membuat struct atau blueprint gampangnya
type Restaurant struct {
	Name      string
	totalChef int
	Is24Hour  bool
}

func NewRestaurant(restaurantName string, totalChef int, is24Hour bool) *Restaurant {
	newRestaurant := Restaurant{
		Name:      restaurantName,
		totalChef: totalChef,
		Is24Hour:  is24Hour,
	}

	return &newRestaurant
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
	fmt.Println("Total Chef ", r.totalChef-tambahan, " + Chef Tambahan ", tambahan, " = ", r.totalChef)
}

func (r Restaurant) GetTotalChef() int {
	return r.totalChef
}

func (r Restaurant) Tutup(){
	fmt.Printf("Perhatian Restaurant: %s telah ditutup, mohon untuk mempersiapkan keluar dari Restaurant\n", r.Name)
}
func Masak(pesanan string, waktuDetik time.Duration, wg *sync.WaitGroup){
	defer wg.Done()

	fmt.Printf("Koki Mulai memasak %s \n", pesanan)
	fmt.Printf("Dengan durasi: %d\n", waktuDetik)
	time.Sleep(time.Duration(waktuDetik) * time.Second)
	fmt.Printf("BAM!!!. Pesanan %s selesai\n", pesanan)
}