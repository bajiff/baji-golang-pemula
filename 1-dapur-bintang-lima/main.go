// ! main.go
package main

import (
	"dapur-bintang-lima/cabang"
	"fmt"
	"sync"
)

type Facilities interface {
	Tutup()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go cabang.Masak("Nasi Kebuli",3, &wg)
	go cabang.Masak("Nasi Lengko",2, &wg)
	go cabang.Masak("Nasi Lengko",2, &wg)
	go cabang.Masak("Nasi Lengko",2, &wg)
	
	wg.Wait()

	fmt.Printf("=== Semua Pesanan Selesai, Restoran Di tutup ===")
}

