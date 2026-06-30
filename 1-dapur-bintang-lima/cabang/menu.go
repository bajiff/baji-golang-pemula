// ! cabang/menu.go
package cabang

import "fmt"

type Menu struct {
	availableItem int
	Menu          string
	worker        int
	IsSoldOut     bool
}

func NewMenu(menuName string) *Menu {
	newMenu := Menu{
		availableItem: 1000,
		Menu:          menuName,
		worker:        50,
		IsSoldOut:     false,
	}

	return &newMenu
}

func (m Menu) BukaMenu() {
	fmt.Println("Perhatian!, Menu:", m.Menu, " dibuka!")
}

func (m Menu) BuatPesanan(jumlahPesanan int) {
	fmt.Println("Menu: ", m.Menu, " menerima pesanan: ", jumlahPesanan)
}
func (m Menu) GetTotalMenu() int {
	return m.availableItem
}

func (m *Menu) TambahMenu(tambahan int) int {
	m.availableItem = m.availableItem + tambahan
	fmt.Printf("Menu sekarang: %d + Menu Tambahan: %d = %d", m.availableItem-tambahan, tambahan, m.availableItem)
	fmt.Println("")
	return m.availableItem
}
func (m Menu) GetTotalWorker() int {
	return m.worker
}

func (m *Menu) AddWorker(worker int) int {
	m.worker = m.worker + worker
	fmt.Printf("Worker sekarang: %d + Worker Tambahan: %d = %d", (m.worker-worker), worker, m.worker)
	fmt.Println("")
	return m.worker
}

func (m Menu) Tutup(){
	fmt.Printf("Perhatian Menu: %s telah ditututp, tidak bisa memesan, semuanya sudah ditarik", m.Menu)
}