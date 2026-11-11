package main

import "fmt"

type Buku struct {
	   id		  string
	   judul	  string
	   tersedia   bool
}

func ()  {
	daftarBuku := []Buku{
			{id: "B01", judul: "Pemrograman Go untuk Pemula", tersedia: true},
			{id: "B02", judul: "Struktur Data dan Algoritma", tersedia: false},
			{id: "B03", judul: "Basis Data Relasional", tersedia: true}
}

var lanjut string = "y"

for lanjut == "y" {
		fmt.Println("\n==========================================")
		fmt.Println("	APLIKASI PEMINJAMAN PERPUSTAKAAN DIGITAL")
		fmt.Println("==========================================")
		fmt.Println("Daftar Buku Tersedia / Tidak Tersedia:")
		for i, b := range daftarBuku {
				status := "Kosong (Dipinjam)"
				if b.tersedia {
						status = "Tersedia"
				}
				fmt.Println("%d. [%s] %s (%s)\n", i+1, b.id, b.judul, status)
		}

		var pilihan int
		fmt.Println("\nMasukkan nomorbuku yang ingin dipinjam:  ")
		fmt.Scan(&pilihan)

		if pilihan >= 1 && pilihan <= len(daftarBuku) {
				bukuPilihan := daftarBuku[pilihan-1]

				if bukuPilihan.tersedia {
						fmt.Printf("/n[INFO] Buku '%s' tersedia!\n", bukuPilihan.judul)
						var konfirmasi string
						fmt.Print("Apakah Anda jadi meminjam buku ini? (y/n):  ")
						fmt.Scan(&konfirmasi)

						if konfirmasi == "y" {
								fmt.Println(">> Berhasil! Buku berhasil dipinjam. Selamat membaca!")
						} else {
								fmt.Println(">> Peminjaman buku dibatalkan.")
						}
				} else {
						fmt.Println("\n[INFO] Maaf, buku '%s' sedang kosong/dipinjam orang lain.\n", bukuPilihan.judul)
						var opsilain string
						fmt.Print("Apakah Anda ingin meminjam buku lain? (y/n): ")
						fmt.Scan(&opsilain)

						if opsilain == "y"{
								fmt.Println(">> Silakan pilih kembali buku yang lain pada menu berikutnya.")
						} else {
								fmt.Println(">> Baik, transaksi diakhiri.")
						}
				}
		} else {
				fmt.Println("[ERROR] Nomor buku tidak valid! Silakan coba lagi.")
		}

		fmt.Print("\nIngin melakukan transaksi lain? (y/n): ")
		fmt.Scan(&lanjut)
	}
	fmt.Println("\nTerima kasih telah menggunakan Perpustakaan Digital!")
}