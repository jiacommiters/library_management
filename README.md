# Catatan Perbaikan Proyek Library Management

## Ringkasan Proyek

Proyek ini adalah sistem manajemen perpustakaan berbasis CLI (Command Line Interface) yang dibuat dengan bahasa Go. Aplikasi memungkinkan pengguna untuk menambah, melihat, mengubah, menghapus, mencari buku, dan menampilkan statistik sederhana dari koleksi buku.

**Fitur utama:**
- Input buku (judul, penulis, kategori, tahun terbit, penerbit, stok, status otomatis)
- Lihat daftar buku dalam tabel rapi (rata kiri)
- Manajemen data: update field tertentu dan hapus buku (dengan pembaruan ID otomatis)
- Pencarian buku berdasarkan judul (linear search)
- Statistik koleksi: jumlah buku, ketersediaan, stok, tahun terbit, dll.
- Sorting (ascending/descending) siap pakai (belum terhubung ke menu)

**Status saat ini:** Aplikasi sudah berjalan untuk skenario dasar, namun masih ada beberapa kekurangan yang harus diperbaiki agar fitur berfungsi penuh dan input lebih handal.

---



  
## Daftar Masalah yang Harus Diperbaiki

### 🔴 Kritis
- [ ] **Menu "stats" tidak terhubung.** Fungsi `statsBuku` sudah dibuat, tetapi belum ada `case "stats"` di switch utama `main()`. Akibatnya, fitur statistik tidak bisa diakses sama sekali.
  - *Cara perbaiki:* Tambahkan case "stats" di `main`, panggil `statsBuku(&dataBuku, jumBuku)`, lalu lanjutkan prompt aksi. 
  - Done

### 🟡 Medium
- [ ] **Input teks tidak bisa mengandung spasi.** Semua input string menggunakan `fmt.Scan` sehingga judul/penulis/penerbit hanya terbaca satu kata (contoh: "Laskar Pelangi" → "Laskar").
  - *Opsi perbaikan:* Ganti dengan `bufio.Scanner` untuk membaca satu baris penuh, atau beri instruksi tegas kepada pengguna agar mengganti spasi dengan `_`.
- [ ] **Tidak ada error handling untuk input angka.** `fmt.Scan` untuk tahun terbit dan stok tidak diperiksa nilai kembaliannya. Jika pengguna memasukkan huruf, data menjadi 0 tanpa peringatan, dan buffer bisa kacau.
  - *Cara perbaiki:* Periksa `n, err := fmt.Scan(...)`; jika `err != nil`, beri tahu pengguna dan bersihkan buffer.

### 🟢 Low
- [ ] **Fungsi `sortingDesc` tidak pernah dipakai.** Ini menambah kode mati. Sebaiknya dihapus atau dihubungkan ke opsi tampilan (misalnya sorting di menu "cek").
- [ ] **Validasi status ketersediaan saat update kurang ketat.** Di menu update, pengguna bisa mengubah status menjadi "Tersedia" padahal stok 0, atau sebaliknya. Idealnya, aturan seperti di input diterapkan: jika stok 0, otomatis "Tidak Tersedia", dan jika stok > 0, otomatis "Tersedia" atau setidaknya dicegah.
- [ ] **Variabel di fungsi `statsBuku` bisa dirapikan.** Ada deklarasi `var stok, tahun int` di luar loop yang sebenarnya bisa langsung menggunakan `:=` di dalam loop.

### 📌 Catatan Tambahan
- Data belum persisten (hilang saat program berhenti). Untuk pengujian berulang yang lebih mudah, disarankan menambahkan penyimpanan ke file atau database sederhana (misal SQLite) di versi selanjutnya.
- Batasan 999 buku (array tetap) sudah diatasi dengan pengecekan kapasitas, tetapi jika ingin lebih fleksibel bisa diganti slice dinamis (`[]buku`).

### Note ASPRAK
  - dummy data 
  - rapihin output
  - lengkapin spesifikasi nya sorting nya sesuai denganspesisifikasi dokumen yang diberikan oleh dosen 
  - hafalin sorting searching and logic nya 
  - untuk mengakses menu nya mending pakai angka aja opsi 1, 2, 3, 4, 5.

---

**Repo:** https://github.com/jiacommiters/library_management  
**Terakhir diperbarui:** Rabu, 11 Juni 2026