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
			showBukuFunc(dataBuku, jumBuku)
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
		}
	}
}

func inputBukuFunc(databuku *arrBuku, jumBuku *int) {

	var action string

	for action != "STOP" {

		databuku[*jumBuku].id = *jumBuku + 1

		fmt.Println("Masukan Judul:")
		fmt.Scan(&databuku[*jumBuku].judul)

		fmt.Println("Masukkan penulis:")
		fmt.Scan(&databuku[*jumBuku].penulis)

		fmt.Println("Masukkan kategori:")
		fmt.Scan(&databuku[*jumBuku].kategori)

		fmt.Println("Masukkan tahun terbit:")
		fmt.Scan(&databuku[*jumBuku].tahunTerbit)

		fmt.Println("Masukkan penerbit:")
		fmt.Scan(&databuku[*jumBuku].penerbit)

		fmt.Println("Masukkan stok:")
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

func showBukuFunc(databuku arrBuku, jumBuku int) {
	var i int
	if jumBuku == 0 {
		fmt.Println("Tidak ada buku yang tersedia")
	} else {
		fmt.Printf("%--25s %--25s %--25s %--25s %--25s %--25s, %--25s %--25s\n", "No", "Judul", "Penulis", "Kategori", "Tahun Terbit", "Penerbit", "Stok", "Tersedia")
		for i = range jumBuku {
			fmt.Printf("%--25d %--25s %--25s %--25s %--25d %--25s %--25d %--25s\n", databuku[i].id, databuku[i].judul, databuku[i].penulis, databuku[i].kategori, databuku[i].tahunTerbit, databuku[i].penerbit, databuku[i].stok, databuku[i].tersedia)
		}
	}
}

func manageBuku(dataBuku *arrBuku, jumBuku *int) {
	var action string
	var idx int
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
				showBukuFunc(*dataBuku, *jumBuku)
				fmt.Printf("Masukkan ID: ")
				fmt.Scan(&idx)

				if idx < 1 || idx > *jumBuku {
					fmt.Printf("Buku dengan id %d tidak ada\n", idx)
				} else {
					var indexEdit int
					indexEdit = idx - 1
					fmt.Printf("Pada id %d ingin mengubah bagian mana?\n(judul/penulis/kategori/tahunterbit/penerbit/stok/statusketersediaan)\nPilihan: ", idx)
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
					case "statusketersediaan":
						if (*dataBuku)[indexEdit].stok == 0 {
							(*dataBuku)[indexEdit].tersedia = "Tidak Tersedia"
							fmt.Println("Stok 0, otomatis diset 'Tidak Tersedia'.")
						} else {
							fmt.Print("Kenapa ingin diubah menjadi tidak tersedia walau buku masih tersedia?: ")
							fmt.Scan(&updatePart.tersedia)
							(*dataBuku)[indexEdit].tersedia = updatePart.tersedia
						}
					default:
						fmt.Println("Pilihan bagian tidak valid!")
					}
					fmt.Println("Data buku berhasil diperbarui!")
				}

			case "delete":
				var indexHapus, indexTerakhir int
				fmt.Println("\nIngin hapus data buku dengan id berapa?")
				showBukuFunc(*dataBuku, *jumBuku)
				fmt.Printf("Masukkan ID: ")
				fmt.Scan(&idx)

				if idx < 1 || idx > *jumBuku {
					fmt.Printf("Buku dengan id %d tidak ada\n", idx)
				} else {
					indexHapus = idx - 1
					indexTerakhir = *jumBuku - 1

					(*dataBuku)[indexHapus] = (*dataBuku)[indexTerakhir]

					(*dataBuku)[indexTerakhir] = buku{}

					*jumBuku--

					fmt.Println("Buku berhasil dihapus dari database!")

					if *jumBuku == 0 {
						fmt.Println("Database sekarang kosong.")
						return
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

		for j >= 0 && dataBuku[j].id > key.id {
			dataBuku[j+1] = dataBuku[j]
			j = j - 1 
		}

		dataBuku[j+1] = key
	}
}

func asortingDesc(dataBuku *arrBuku, jumBuku int) {
	var i, j int
	var key buku

	for i = 1; i < jumBuku; i++ {
		key = dataBuku[i]
		j = i - 1

		for j >= 0 && dataBuku[j].id < key.id {
			dataBuku[j+1] = dataBuku[j]
			j = j - 1 
		}
		
		dataBuku[j+1] = key
}

}

func searchBuku(dataBuku *arrBuku, jumBuku int, target string) {
	var left, right int 
	var mid int 
	left, right = 0, jumBuku-1
	mid = (left + right) / 2

	for left <= right && (*dataBuku)[mid].judul != target {
		if (*dataBuku)[mid].judul > (*dataBuku)[left].judul {
			left = mid + 1
		} else {
			right = mid - 1
		}
		mid = (left + right) / 2
	}

	if (*dataBuku)[mid].judul == target {
		fmt.Printf("Data Ditemukan\n%--25s%--25s%--25s%--25s%--25s%--25s%--25s%--25s\n", "ID", "Judul", "Penulis", "Kategori", "Tahun Terbit", "Penerbit", "Stok", "Tersedia")
		fmt.Printf("%--25d%--25s%--25s%--25s%--25d%--25s%--25d%--25s", dataBuku[mid].id, dataBuku[mid].judul, dataBuku[mid].penulis, dataBuku[mid].kategori, dataBuku[mid].tahunTerbit, dataBuku[mid].penerbit, dataBuku[mid].stok, dataBuku[mid].tersedia)
	} else {
		fmt.Printf("Buku dengan judul %s tidak ada", target)
	}
}
