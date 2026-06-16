package main

import (
	"fmt"
)

type buku struct {
	id          int
	judul       string
	penulis     string
	kategori    string
	tahunTerbit int
	penerbit    string
	stok        int
	tersedia    string
}

type arrBuku [999]buku

func tampilMenu() {
	fmt.Println("\n+------------------------------------+")
	fmt.Println("|         +++ SiPerpus +++           |")
	fmt.Println("+------------------------------------+")
	fmt.Println("| Pilih aksi yang ingin dilakukan:   |")
	fmt.Println("| 1. Tambah buku baru                |")
	fmt.Println("| 2. Tampilkan semua buku            |")
	fmt.Println("| 3. Update / Hapus buku             |")
	fmt.Println("| 4. Cari buku                       |")
	fmt.Println("| 5. Urutkan buku                    |")
	fmt.Println("| 6. Statistik koleksi               |")
	fmt.Println("| 0. Keluar program                  |")
	fmt.Println("+------------------------------------+")
	fmt.Print("Pilihan: ")
}

func main() {

	var dataBuku arrBuku
	var action string
	var jumBuku int = 0

	fmt.Println("\n+++ SiPerpus +++")
	fmt.Println("Selamat Datang di Sistem Manajemen Perpustakaan Digital!")
	tampilMenu()
	fmt.Scan(&action)

	for action != "0" {
		switch action {
		case "1":
			fmt.Println("\nKetik 'STOP' jika ingin berhenti input buku")
			inputBukuFunc(&dataBuku, &jumBuku)
		case "2":
			showBukuFunc(&dataBuku, jumBuku)
		case "3":
			manageBuku(&dataBuku, &jumBuku)
		case "4":
			searchMenu(&dataBuku, jumBuku)
		case "5":
			sortMenu(&dataBuku, &jumBuku)
		case "6":
			statsBuku(&dataBuku, jumBuku)
		default:
			fmt.Println("\nPilihan tidak dikenal! Silakan coba lagi.")
		}
		tampilMenu()
		fmt.Scan(&action)
	}

	fmt.Println("\nTerima kasih telah menggunakan +++ SiPerpus +++")
}

func inputBukuFunc(databuku *arrBuku, jumBuku *int) {
	var action string

	for action != "STOP" {
		if *jumBuku >= len(*databuku) {
			fmt.Println("\nKapasitas penuh! Tidak dapat menambahkan buku lagi.")
			action = "STOP"
		} else {
			databuku[*jumBuku].id = *jumBuku + 1

			fmt.Println("\n----------------------------------------------")
			fmt.Println("Catatan: Ganti spasi dengan '_' (contoh: Harry_Potter)")
			fmt.Println("----------------------------------------------")
			
			fmt.Print("Masukan Judul          : ")
			fmt.Scan(&databuku[*jumBuku].judul)
			
			fmt.Print("Masukkan Penulis       : ")
			fmt.Scan(&databuku[*jumBuku].penulis)
			
			fmt.Print("Masukkan Kategori      : ")
			fmt.Scan(&databuku[*jumBuku].kategori)
			
			fmt.Print("Masukkan Tahun Terbit  : ")
			fmt.Scan(&databuku[*jumBuku].tahunTerbit)
			
			fmt.Print("Masukkan Penerbit      : ")
			fmt.Scan(&databuku[*jumBuku].penerbit)
			
			fmt.Print("Masukkan Stok          : ")
			fmt.Scan(&databuku[*jumBuku].stok)

			if databuku[*jumBuku].stok == 0 {
				databuku[*jumBuku].tersedia = "Tidak_Tersedia"
			} else {
				databuku[*jumBuku].tersedia = "Tersedia"
			}

			*jumBuku++
			
			fmt.Println("\n✓ Buku berhasil ditambahkan!")
			fmt.Println("----------------------------------------------")
			fmt.Print("Tambah buku lagi? (LANJUT/STOP): ")
			fmt.Scan(&action)
		}
	}
}

func showBukuFunc(databuku *arrBuku, jumBuku int) {
	var i int
	if jumBuku == 0 {
		fmt.Println("\nBelum ada data buku.")
	} else {
		fmt.Println()
		fmt.Printf("%-4s %-25s %-18s %-12s %-6s %-18s %-5s %-15s\n",
			"ID", "Judul", "Penulis", "Kategori", "Tahun", "Penerbit", "Stok", "Tersedia")
		fmt.Println("----------------------------------------------------------------------------------------------------")
		for i = 0; i < jumBuku; i++ {
			fmt.Printf("%-4d %-25s %-18s %-12s %-6d %-18s %-5d %-15s\n",
				databuku[i].id, databuku[i].judul, databuku[i].penulis,
				databuku[i].kategori, databuku[i].tahunTerbit,
				databuku[i].penerbit, databuku[i].stok, databuku[i].tersedia)
		}
		fmt.Println("----------------------------------------------------------------------------------------------------")
	}
}

func cetakDetailBuku(dataBuku *arrBuku, index int) {
	fmt.Println()
	fmt.Printf("%-4s %-25s %-18s %-12s %-6s %-18s %-5s %-15s\n",
		"ID", "Judul", "Penulis", "Kategori", "Tahun", "Penerbit", "Stok", "Tersedia")
	fmt.Println("----------------------------------------------------------------------------------------------------")
	fmt.Printf("%-4d %-25s %-18s %-12s %-6d %-18s %-5d %-15s\n",
		dataBuku[index].id, dataBuku[index].judul, dataBuku[index].penulis,
		dataBuku[index].kategori, dataBuku[index].tahunTerbit,
		dataBuku[index].penerbit, dataBuku[index].stok, dataBuku[index].tersedia)
}

func manageBuku(dataBuku *arrBuku, jumBuku *int) {
	var action string
	var id int
	var update string
	var updatePart buku

	if *jumBuku == 0 {
		fmt.Println("\nTidak ada buku yang tersedia, tidak bisa manage buku.")
		return
	}

	fmt.Println("\n----------------------------------------------")
	fmt.Println("              Menu Manage Buku")
	fmt.Println("----------------------------------------------")
	fmt.Println("update - Ubah data buku")
	fmt.Println("delete - Hapus buku")
	fmt.Println("STOP   - Kembali ke menu")
	fmt.Println("----------------------------------------------")
	fmt.Print("Pilihan: ")
	fmt.Scan(&action)

	for action != "STOP" {
		switch action {
		case "update":
			showBukuFunc(dataBuku, *jumBuku)
			fmt.Println()
			fmt.Print("Masukkan ID buku yang ingin diubah: ")
			fmt.Scan(&id)

			if id < 1 || id > *jumBuku {
				fmt.Printf("\nBuku dengan ID %d tidak ada.\n", id)
			} else {
				var indexEdit int = id - 1
				fmt.Println("\n----------------------------------------------")
				fmt.Println("Bagian yang ingin diubah?")
				fmt.Println("  judul | penulis | kategori | tahunterbit")
				fmt.Println("  penerbit | stok | statusketersediaan")
				fmt.Println("----------------------------------------------")
				fmt.Print("Pilihan: ")
				fmt.Scan(&update)
				fmt.Println()

				switch update {
				case "judul":
					fmt.Print("Judul baru          : ")
					fmt.Scan(&updatePart.judul)
					(*dataBuku)[indexEdit].judul = updatePart.judul
				case "penulis":
					fmt.Print("Penulis baru        : ")
					fmt.Scan(&updatePart.penulis)
					(*dataBuku)[indexEdit].penulis = updatePart.penulis
				case "kategori":
					fmt.Print("Kategori baru       : ")
					fmt.Scan(&updatePart.kategori)
					(*dataBuku)[indexEdit].kategori = updatePart.kategori
				case "tahunterbit":
					fmt.Print("Tahun terbit baru   : ")
					fmt.Scan(&updatePart.tahunTerbit)
					(*dataBuku)[indexEdit].tahunTerbit = updatePart.tahunTerbit
				case "penerbit":
					fmt.Print("Penerbit baru       : ")
					fmt.Scan(&updatePart.penerbit)
					(*dataBuku)[indexEdit].penerbit = updatePart.penerbit
				case "stok":
					fmt.Print("Stok baru           : ")
					fmt.Scan(&updatePart.stok)
					(*dataBuku)[indexEdit].stok = updatePart.stok
					// memperbarui status tersedia otomatis setelah stok diubah
					if (*dataBuku)[indexEdit].stok > 0 {
						(*dataBuku)[indexEdit].tersedia = "Tersedia"
					} else {
						(*dataBuku)[indexEdit].tersedia = "Tidak_Tersedia"
					}
				case "statusketersediaan":
					fmt.Print("Status baru (Tersedia/Tidak_Tersedia): ")
					fmt.Scan(&updatePart.tersedia)
					if updatePart.tersedia == "Tersedia" || updatePart.tersedia == "Tidak_Tersedia" {
						(*dataBuku)[indexEdit].tersedia = updatePart.tersedia
					} else {
						fmt.Println("\nStatus tidak valid! Gunakan: Tersedia atau Tidak_Tersedia")
					}
				default:
					fmt.Println("\nPilihan bagian tidak valid!")
				}
				fmt.Println("\n✓ Data buku berhasil diperbarui!")
			}

		case "delete":
			var indexHapus, indexTerakhir, i int
			showBukuFunc(dataBuku, *jumBuku)
			fmt.Println()
			fmt.Print("Masukkan ID buku yang ingin dihapus: ")
			fmt.Scan(&id)

			if id < 1 || id > *jumBuku {
				fmt.Printf("\nBuku dengan ID %d tidak ada.\n", id)
			} else {
				indexHapus = id - 1
				indexTerakhir = *jumBuku - 1

				(*dataBuku)[indexHapus] = (*dataBuku)[indexTerakhir]
				(*dataBuku)[indexTerakhir] = buku{}
				*jumBuku--

				fmt.Println("\n✓ Buku berhasil dihapus!")

				if *jumBuku == 0 {
					fmt.Println("Database sekarang kosong.")
					return
				}
				for i = indexHapus; i < *jumBuku; i++ {
					(*dataBuku)[i].id = i + 1
				}
			}

		case "STOP":
			return

		default:
			fmt.Println("\nPilihan tidak dikenal. Gunakan: update, delete, atau STOP")
		}

		fmt.Println("\n----------------------------------------------")
		fmt.Print("Mau manage lagi? (update/delete/STOP): ")
		fmt.Scan(&action)
	}
}

// BAGIAN ATTALA — SORTING & SEARCHING

func selectionSortAsc(dataBuku *arrBuku, jumBuku int) {
	var i, j, minIdx int
	var temp buku

	for i = 0; i < jumBuku-1; i++ {
		minIdx = i
		for j = i + 1; j < jumBuku; j++ {
			if dataBuku[j].tahunTerbit < dataBuku[minIdx].tahunTerbit {
				minIdx = j
			}
		}
		temp = dataBuku[minIdx]
		dataBuku[minIdx] = dataBuku[i]
		dataBuku[i] = temp
	}
}

func selectionSortDesc(dataBuku *arrBuku, jumBuku int) {
	var i, j, maxIdx int
	var temp buku

	for i = 0; i < jumBuku-1; i++ {
		maxIdx = i
		for j = i + 1; j < jumBuku; j++ {
			if dataBuku[j].tahunTerbit > dataBuku[maxIdx].tahunTerbit {
				maxIdx = j
			}
		}
		temp = dataBuku[maxIdx]
		dataBuku[maxIdx] = dataBuku[i]
		dataBuku[i] = temp
	}
}

func insertionSortAsc(dataBuku *arrBuku, jumBuku int) {
	var i, j int
	var key buku

	for i = 1; i < jumBuku; i++ {
		key = dataBuku[i]
		j = i - 1
		for j >= 0 && dataBuku[j].tahunTerbit > key.tahunTerbit {
			dataBuku[j+1] = dataBuku[j]
			j = j - 1
		}
		dataBuku[j+1] = key
	}
}

func insertionSortDesc(dataBuku *arrBuku, jumBuku int) {
	var i, j int
	var key buku

	for i = 1; i < jumBuku; i++ {
		key = dataBuku[i]
		j = i - 1
		for j >= 0 && dataBuku[j].tahunTerbit < key.tahunTerbit {
			dataBuku[j+1] = dataBuku[j]
			j = j - 1
		}
		dataBuku[j+1] = key
	}
}

func sortMenu(dataBuku *arrBuku, jumBuku *int) {
	if *jumBuku == 0 {
		fmt.Println("\nTidak ada buku untuk diurutkan.")
		return
	}

	var jenis, urutan string
	fmt.Println("\n----------------------------------------------")
	fmt.Println("           Pilih Metode Pengurutan")
	fmt.Println("----------------------------------------------")
	fmt.Println("1. Selection Sort")
	fmt.Println("2. Insertion Sort")
	fmt.Println("----------------------------------------------")
	fmt.Print("Pilihan: ")
	fmt.Scan(&jenis)

	fmt.Println("\n----------------------------------------------")
	fmt.Println("             Pilih Arah Urutan")
	fmt.Println("----------------------------------------------")
	fmt.Println("1. Ascending  (Terlama → Terbaru)")
	fmt.Println("2. Descending (Terbaru → Terlama)")
	fmt.Println("----------------------------------------------")
	fmt.Print("Pilihan: ")
	fmt.Scan(&urutan)
	fmt.Println()

	// menjalankan fungsi sort sesuai pilihan metode dan urutan (attala)
	if jenis == "1" && urutan == "1" {
		selectionSortAsc(dataBuku, *jumBuku)
		fmt.Println("✓ Berhasil diurutkan dengan Selection Sort Ascending!")
	} else if jenis == "1" && urutan == "2" {
		selectionSortDesc(dataBuku, *jumBuku)
		fmt.Println("✓ Berhasil diurutkan dengan Selection Sort Descending!")
	} else if jenis == "2" && urutan == "1" {
		insertionSortAsc(dataBuku, *jumBuku)
		fmt.Println("✓ Berhasil diurutkan dengan Insertion Sort Ascending!")
	} else if jenis == "2" && urutan == "2" {
		insertionSortDesc(dataBuku, *jumBuku)
		fmt.Println("✓ Berhasil diurutkan dengan Insertion Sort Descending!")
	} else {
		fmt.Println("Pilihan tidak valid!")
		return
	}

	showBukuFunc(dataBuku, *jumBuku)
}

func sequentialSearchJudul(dataBuku *arrBuku, jumBuku int, target string) {
	var i int
	var found int = -1

	for i = 0; i < jumBuku && found == -1; i++ {
		if (*dataBuku)[i].judul == target {
			found = i
		}
	}

	if found == -1 {
		fmt.Printf("\nBuku dengan judul '%s' tidak ditemukan.\n", target)
	} else {
		fmt.Println("\n✓ Buku ditemukan!")
		cetakDetailBuku(dataBuku, found)
	}
}

func sequentialSearchID(dataBuku *arrBuku, jumBuku int, targetID int) {
	var i int
	var found int = -1

	for i = 0; i < jumBuku && found == -1; i++ {
		if (*dataBuku)[i].id == targetID {
			found = i
		}
	}

	if found == -1 {
		fmt.Printf("\nBuku dengan ID %d tidak ditemukan.\n", targetID)
	} else {
		fmt.Println("\n✓ Buku ditemukan!")
		cetakDetailBuku(dataBuku, found)
	}
}

func binarySearchID(dataBuku *arrBuku, jumBuku int, targetID int) {
	var i, j, minIdx int
	var temp buku
	for i = 0; i < jumBuku-1; i++ {
		minIdx = i
		for j = i + 1; j < jumBuku; j++ {
			if dataBuku[j].id < dataBuku[minIdx].id {
				minIdx = j
			}
		}
		temp = dataBuku[minIdx]
		dataBuku[minIdx] = dataBuku[i]
		dataBuku[i] = temp
	}

	var low, high, mid int
	low = 0
	high = jumBuku - 1
	var found int = -1

	for low <= high && found == -1 {
		mid = (low + high) / 2
		if (*dataBuku)[mid].id == targetID {
			found = mid
		} else if (*dataBuku)[mid].id < targetID {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if found == -1 {
		fmt.Printf("\nBuku dengan ID %d tidak ditemukan.\n", targetID)
	} else {
		fmt.Println("\n✓ Buku ditemukan!")
		cetakDetailBuku(dataBuku, found)
	}
}

func searchMenu(dataBuku *arrBuku, jumBuku int) {
	if jumBuku == 0 {
		fmt.Println("\nTidak ada buku untuk dicari.")
		return
	}

	var jenis string
	fmt.Println("\n----------------------------------------------")
	fmt.Println("          Pilih Metode Pencarian")
	fmt.Println("----------------------------------------------")
	fmt.Println("1. Sequential Search by Judul")
	fmt.Println("2. Sequential Search by ID")
	fmt.Println("3. Binary Search by ID")
	fmt.Println("----------------------------------------------")
	fmt.Print("Pilihan: ")
	fmt.Scan(&jenis)
	fmt.Println()

	switch jenis {
	case "1":
		var target string
		fmt.Print("Masukkan judul buku: ")
		fmt.Scan(&target)
		sequentialSearchJudul(dataBuku, jumBuku, target)
	case "2":
		var targetID int
		fmt.Print("Masukkan ID buku: ")
		fmt.Scan(&targetID)
		sequentialSearchID(dataBuku, jumBuku, targetID)
	case "3":
		var targetID int
		fmt.Print("Masukkan ID buku: ")
		fmt.Scan(&targetID)
		fmt.Println("(Data otomatis diurutkan berdasarkan ID sebelum pencarian)")
		binarySearchID(dataBuku, jumBuku, targetID)
	default:
		fmt.Println("Pilihan tidak valid!")
	}
}

func statsBuku(dataBuku *arrBuku, jumBuku int) {
	if jumBuku == 0 {
		fmt.Println("\nTidak ada data untuk statistik.")
		return
	}

	var tersedia, tidakTersedia, totalStok int
	var totalTahun, stok, tahun int
	var tahunTertua, tahunTerbaru int
	var judulTertua, judulTerbaru, judul string
	var stokTerbanyak, stokTersedikit int
	var judulTerbanyak, judulTersedikit string
	var rataStok, rataTahun, persenTersedia float64
	var i int

	var kategoriList [100]string
	var kategoriCount [100]int
	var jumKategori int = 0

	tahunTertua = (*dataBuku)[0].tahunTerbit
	tahunTerbaru = (*dataBuku)[0].tahunTerbit
	judulTertua = (*dataBuku)[0].judul
	judulTerbaru = (*dataBuku)[0].judul
	stokTerbanyak = (*dataBuku)[0].stok
	stokTersedikit = (*dataBuku)[0].stok
	judulTerbanyak = (*dataBuku)[0].judul
	judulTersedikit = (*dataBuku)[0].judul

	for i = 0; i < jumBuku; i++ {
		stok = (*dataBuku)[i].stok
		tahun = (*dataBuku)[i].tahunTerbit
		judul = (*dataBuku)[i].judul

		totalStok += stok
		totalTahun += tahun

		if stok > 0 {
			tersedia++
		} else {
			tidakTersedia++
		}

		if tahun < tahunTertua {
			tahunTertua = tahun
			judulTertua = judul
		}
		if tahun > tahunTerbaru {
			tahunTerbaru = tahun
			judulTerbaru = judul
		}
		if stok > stokTerbanyak {
			stokTerbanyak = stok
			judulTerbanyak = judul
		}
		if stok < stokTersedikit {
			stokTersedikit = stok
			judulTersedikit = judul
		}

		var kat string = (*dataBuku)[i].kategori
		var ketemu bool = false
		var k int
		for k = 0; k < jumKategori; k++ {
			if kategoriList[k] == kat {
				kategoriCount[k]++
				ketemu = true
			}
		}
		if !ketemu {
			kategoriList[jumKategori] = kat
			kategoriCount[jumKategori] = 1
			jumKategori++
		}
	}

	rataStok = float64(totalStok) / float64(jumBuku)
	rataTahun = float64(totalTahun) / float64(jumBuku)
	persenTersedia = float64(tersedia) / float64(jumBuku) * 100

	fmt.Println("\n+++ SiPerpus +++")
	fmt.Println("============================================")
	fmt.Println("       STATISTIK PERPUSTAKAAN DIGITAL      ")
	fmt.Println("============================================")
	fmt.Printf("Total Buku              : %d\n", jumBuku)
	fmt.Printf("Buku Tersedia           : %d (%.1f%%)\n", tersedia, persenTersedia)
	fmt.Printf("Buku Tidak Tersedia     : %d\n", tidakTersedia)
	fmt.Printf("Total Stok              : %d eksemplar\n", totalStok)
	fmt.Printf("Rata-rata Stok/Buku     : %.1f\n", rataStok)
	fmt.Printf("Tahun Terbit Tertua     : %d (%s)\n", tahunTertua, judulTertua)
	fmt.Printf("Tahun Terbit Terbaru    : %d (%s)\n", tahunTerbaru, judulTerbaru)
	fmt.Printf("Rata-rata Tahun Terbit  : %.0f\n", rataTahun)
	fmt.Printf("Stok Terbanyak          : %d (%s)\n", stokTerbanyak, judulTerbanyak)
	fmt.Printf("Stok Tersedikit         : %d (%s)\n", stokTersedikit, judulTersedikit)
	fmt.Println("\n-------- Jumlah Buku per Kategori --------")

	var k int
	for k = 0; k < jumKategori; k++ {
		fmt.Printf("  %-20s : %d buku\n", kategoriList[k], kategoriCount[k])
	}
	fmt.Println("============================================")
	fmt.Println("+++ SiPerpus +++")
}