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

func main() {

	var dataBuku arrBuku
	var action string
	var jumBuku int = 0

	fmt.Println("Selamat Datang di Library")
	fmt.Println("Mau Melakukan apa hari ini?")
	fmt.Printf("Mau Input?(input)\nMau manage data buku?(manage)\nMau cek status ketersediaan?(cek)\nCek statistik buku?(stats)\n")
	fmt.Scan(&action)

	for action != "stop" {
		switch action {
		case "input":
			fmt.Println("Ketik 'STOP' jika ingin stop input buku")
			inputBukuFunc(&dataBuku, &jumBuku)
			fmt.Println("Mau Melakukan aksi apa lagi? ")
			fmt.Printf("Mau Input?(input)\nMau manage data buku?(manage)\nMau cek status ketersediaan?(cek)\nCek statistik buku?(stats)\n")
			fmt.Scan(&action)
		case "cek":
			showBukuFunc(&dataBuku, jumBuku)
			fmt.Println("Mau Melakukan aksi apa lagi? ")
			fmt.Printf("Mau Input?(input)\nMau manage data buku?(manage)\nMau cek status ketersediaan?(cek)\nCek statistik buku?(stats)\n")
			fmt.Scan(&action)
		case "manage":
			manageBuku(&dataBuku, &jumBuku)
			fmt.Println("Mau Melakukan aksi apa lagi? ")
			fmt.Printf("Mau Input?(input)\nMau manage data buku?(manage)\nMau cek status ketersediaan?(cek)\nCek statistik buku?(stats)\n")
			fmt.Scan(&action)
		case "search":
			fmt.Println("Buku apa yang ingin anda cari?")
			var target string
			fmt.Scan(&target)
			searchBuku(&dataBuku, jumBuku, target)
			fmt.Println("Mau Melakukan aksi apa lagi? ")
			fmt.Printf("Mau Input?(input)\nMau manage data buku?(manage)\nMau cek status ketersediaan?(cek)\nCek statistik buku?(stats)\n")
			fmt.Scan(&action)
		case "stats":
			statsBuku(&dataBuku, jumBuku)
			fmt.Println("Mau Melakukan aksi apa lagi? ")
			fmt.Printf("Mau Input?(input)\nMau manage data buku?(manage)\nMau cek status ketersediaan?(cek)\nCek statistik buku?(stats)\n")
			fmt.Scan(&action)
		}
	}
}

func inputBukuFunc(databuku *arrBuku, jumBuku *int) {

	var action string

	for action != "STOP" {

		if *jumBuku >= len(*databuku) {
			fmt.Println("Tidak dapat menambahkan buku lagi kapasitas sudah penuh!")
			action = "STOP"
		} else {

			databuku[*jumBuku].id = *jumBuku + 1

			fmt.Println("Jangan menggunakan spasi, jika ingin ganti dengan '_' saja")
			fmt.Print("Masukan Judul:")
			fmt.Scan(&databuku[*jumBuku].judul)

			fmt.Print("Masukkan penulis:")
			fmt.Scan(&databuku[*jumBuku].penulis)

			fmt.Print("Masukkan kategori:")
			fmt.Scan(&databuku[*jumBuku].kategori)

			fmt.Print("Masukkan tahun terbit:")
			fmt.Scan(&databuku[*jumBuku].tahunTerbit)

			fmt.Print("Masukkan penerbit:")
			fmt.Scan(&databuku[*jumBuku].penerbit)

			fmt.Print("Masukkan stok:")
			fmt.Scan(&databuku[*jumBuku].stok)

			if databuku[*jumBuku].stok == 0 {
				databuku[*jumBuku].tersedia = "Tidak Tersedia"
			} else {
				databuku[*jumBuku].tersedia = "Tersedia"
			}

			*jumBuku++

			fmt.Println("\nInput Buku selesai, apakah ingin menambahkan lagi? (LANJUT/STOP)")
			fmt.Scan(&action)
		}
	}
}

func showBukuFunc(databuku *arrBuku, jumBuku int) {
	var i int
	if jumBuku == 0 {
		fmt.Println("Tidak ada buku yang tersedia")
	} else {
		fmt.Printf("%-25s %-25s %-25s %-25s %-25s %-25s, %-25s %-25s\n", "No", "Judul", "Penulis", "Kategori", "Tahun Terbit", "Penerbit", "Stok", "Tersedia")
		for i = range jumBuku {
			fmt.Printf("%-25d %-25s %-25s %-25s %-25d %-25s %-25d %-25s\n", databuku[i].id, databuku[i].judul, databuku[i].penulis, databuku[i].kategori, databuku[i].tahunTerbit, databuku[i].penerbit, databuku[i].stok, databuku[i].tersedia)
		}
	}
}

func manageBuku(dataBuku *arrBuku, jumBuku *int) {
	var action string
	var id int
	var update string
	var updatePart buku

	if *jumBuku == 0 {
		fmt.Println("Tidak ada buku yang tersedia, tidak bisa manage buku")
		return
	} else {

		for action != "STOP" {
			fmt.Printf("\nMau Manage apa di database kita?\n(update)\n(delete)\n")
			fmt.Printf("Pilihan Anda: ")
			fmt.Scan(&action)

			switch action {
			case "update":
				fmt.Println("\nIngin update data buku dengan id berapa?")
				showBukuFunc(dataBuku, *jumBuku)
				fmt.Printf("Masukkan ID: ")
				fmt.Scan(&id)

				if id < 1 || id > *jumBuku {
					fmt.Printf("Buku dengan id %d tidak ada\n", id)
				} else {
					var indexEdit int
					indexEdit = id - 1
					fmt.Printf("Pada id %d ingin mengubah bagian mana?\n(judul/penulis/kategori/tahunterbit/penerbit/stok/statusketersediaan)\nPilihan: ", id)
					fmt.Scan(&update)

					switch update {
					case "judul":
						fmt.Print("Masukan Judul baru: ")
						fmt.Scan(&updatePart.judul)
						(*dataBuku)[indexEdit].judul = updatePart.judul
					case "penulis":
						fmt.Print("Masukan penulis baru: ")
						fmt.Scan(&updatePart.penulis)
						(*dataBuku)[indexEdit].penulis = updatePart.penulis
					case "kategori":
						fmt.Print("Masukan kategori baru: ")
						fmt.Scan(&updatePart.kategori)
						(*dataBuku)[indexEdit].kategori = updatePart.kategori
					case "tahunterbit":
						fmt.Print("Masukan tahunterbit baru: ")
						fmt.Scan(&updatePart.tahunTerbit)
						(*dataBuku)[indexEdit].tahunTerbit = updatePart.tahunTerbit
					case "penerbit":
						fmt.Print("Masukan penerbit baru: ")
						fmt.Scan(&updatePart.penerbit)
						(*dataBuku)[indexEdit].penerbit = updatePart.penerbit
					case "stok":
						fmt.Print("Masukan stok baru: ")
						fmt.Scan(&updatePart.stok)
						(*dataBuku)[indexEdit].stok = updatePart.stok
						if (*dataBuku)[indexEdit].stok > 0 {
							(*dataBuku)[indexEdit].tersedia = "Tersedia"
						} else {
							(*dataBuku)[indexEdit].tersedia = "Tidak Tersedia"
						}
					case "statusketersediaan":
						fmt.Print("Masukan status baru: ")
						fmt.Scan(&updatePart.tersedia)

						if updatePart.tersedia == "Tersedia" || updatePart.tersedia == "Tidak Tersedia" {
							(*dataBuku)[indexEdit].tersedia = updatePart.tersedia
						} else {
							fmt.Println("Status ketersediaan tidak valid, input salah satu\nTersedia atau Tidak tersedia")
						}

					default:
						fmt.Println("Pilihan bagian tidak valid!")
					}
					fmt.Println("Data buku berhasil diperbarui!")
				}

			case "delete":
				var indexHapus, indexTerakhir, i int
				fmt.Println("\nIngin hapus data buku dengan id berapa?")
				showBukuFunc(dataBuku, *jumBuku)
				fmt.Printf("Masukkan ID: ")
				fmt.Scan(&id)

				if id < 1 || id > *jumBuku {
					fmt.Printf("Buku dengan id %d tidak ada\n", id)
				} else {
					indexHapus = id - 1
					indexTerakhir = *jumBuku - 1

					(*dataBuku)[indexHapus] = (*dataBuku)[indexTerakhir]

					(*dataBuku)[indexTerakhir] = buku{}

					*jumBuku--

					fmt.Println("Buku berhasil dihapus dari database!")

					if *jumBuku == 0 {
						fmt.Println("Database sekarang kosong.")
						return
					} else {
						for i = indexHapus; i < *jumBuku; i++ {
							(*dataBuku)[i].id = i + 1
						}
					}

				}

			default:
				fmt.Println("Pilihan action tidak dikenal. Pilih (update), (delete), atau (STOP).")
			}

			fmt.Printf("\nMau manage lagi? (update/delete) atau ketik 'STOP' untuk keluar: ")
			fmt.Scan(&action)
		}
	}
}

func sortingAsc(dataBuku *arrBuku, jumBuku int) {
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

func sortingDesc(dataBuku *arrBuku, jumBuku int) {
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

func searchBuku(dataBuku *arrBuku, jumBuku int, target string) {
	var i int = 0
	var found int = -1

	for found == -1 && i < jumBuku {
		if (*dataBuku)[i].judul == target {
			found = i
		}
		i++
	}

	if found == -1 {
		fmt.Printf("Buku dengan judul %s tidak ada", target)
	} else {
		fmt.Printf("Data Ditemukan\n%-25s%-25s%-25s%-25s%-25s%-25s%-25s%-25s\n", "ID", "Judul", "Penulis", "Kategori", "Tahun Terbit", "Penerbit", "Stok", "Tersedia")
		fmt.Printf("%-25d%-25s%-25s%-25s%-25d%-25s%-25d%-25s", dataBuku[found].id, dataBuku[found].judul, dataBuku[found].penulis, dataBuku[found].kategori, dataBuku[found].tahunTerbit, dataBuku[found].penerbit, dataBuku[found].stok, dataBuku[found].tersedia)
	}
}

func statsBuku(dataBuku *arrBuku, jumBuku int) {
	if jumBuku == 0 {
		fmt.Println("Tidak ada data untuk statistik.")
		return
	}

	var tersedia, tidakTersedia, totalStok int
	var totalTahun, stok, tahun int
	var tahunTertua, tahunTerbaru int
	var judulTertua, judulTerbaru, judul string
	var stokTerbanyak, stokTersedikit int
	var judulTerbanyak, judulTersedikit string
	var rataStok, rataTahun, persenTersedia float64

	tahunTertua = (*dataBuku)[0].tahunTerbit
	tahunTerbaru = (*dataBuku)[0].tahunTerbit
	judulTertua = (*dataBuku)[0].judul
	judulTerbaru = (*dataBuku)[0].judul
	stokTerbanyak = (*dataBuku)[0].stok
	stokTersedikit = (*dataBuku)[0].stok
	judulTerbanyak = (*dataBuku)[0].judul
	judulTersedikit = (*dataBuku)[0].judul

	for i := 0; i < jumBuku; i++ {
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
	}

	rataStok = float64(totalStok) / float64(jumBuku)
	rataTahun = float64(totalTahun) / float64(jumBuku)
	persenTersedia = float64(tersedia) / float64(jumBuku) * 100

	fmt.Println("\n========== STATISTIK PERPUSTAKAAN ==========")
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
	fmt.Println("============================================")
}
