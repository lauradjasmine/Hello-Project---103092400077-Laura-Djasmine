package main

import (
	"fmt"
)

const maxAkun = 100 // maksimal akun yang bisa dimasukkan
const maxTransaksi = 1000

type Akun struct {
	id        int
	nama      string
	saldo     float64
	disetujui bool
}

type Transaksi struct {
	pengirim int
	penerima int
	jumlah   float64
	jenis    string
}

var daftarAkun [maxAkun]Akun
var jumlahAkun int
var daftarTransaksi [maxTransaksi]Transaksi
var jumlahTransaksi int

func registrasiAkun(id int, nama string, saldo float64) {
	if jumlahAkun >= maxAkun {
		fmt.Println("Kuota akun penuh")
		return
	}
	if sequentialSearchAkun(id) != -1 {
		fmt.Println("ID akun sudah terdaftar")
		return
	}
	var akunBaru Akun
	akunBaru.id = id
	akunBaru.nama = nama
	akunBaru.saldo = saldo
	akunBaru.disetujui = false
	daftarAkun[jumlahAkun] = akunBaru
	jumlahAkun++
	fmt.Println("Akun berhasil diregistrasi:", id)
}

func setujuiAkun(id int, setuju bool) {
	var idx int = sequentialSearchAkun(id)
	if idx != -1 {
		daftarAkun[idx].disetujui = setuju
		var status string = "ditolak"
		if setuju {
			status = "disetujui"
		}
		fmt.Println("Akun", id, status)
	} else {
		fmt.Println("Akun tidak ditemukan")
	}
}

func kirimUang(pengirimID int, penerimaID int, jumlah float64) {
	idxPengirim := sequentialSearchAkun(pengirimID)
	idxPenerima := sequentialSearchAkun(penerimaID)
	if idxPengirim == -1 || idxPenerima == -1 {
		fmt.Println("ID tidak valid")
		return
	}
	if !daftarAkun[idxPengirim].disetujui || !daftarAkun[idxPenerima].disetujui {
		fmt.Println("Akun belum disetujui")
		return
	}
	if daftarAkun[idxPengirim].saldo < jumlah {
		fmt.Println("Saldo tidak cukup")
		return
	}
	daftarAkun[idxPengirim].saldo -= jumlah
	daftarAkun[idxPenerima].saldo += jumlah

	var t Transaksi
	t.pengirim = pengirimID
	t.penerima = penerimaID
	t.jumlah = jumlah
	t.jenis = "Transfer"

	daftarTransaksi[jumlahTransaksi] = t
	jumlahTransaksi++
	fmt.Println("Transfer berhasil")
}

func bayar(id int, jenis string, jumlah float64) {
	var idx int = sequentialSearchAkun(id)
	if idx == -1 || !daftarAkun[idx].disetujui {
		fmt.Println("Akun tidak valid atau belum disetujui")
		return
	}
	if daftarAkun[idx].saldo < jumlah {
		fmt.Println("Saldo tidak cukup")
		return
	}
	daftarAkun[idx].saldo -= jumlah

	var t Transaksi
	t.pengirim = id
	t.penerima = -1
	t.jumlah = jumlah
	t.jenis = jenis

	daftarTransaksi[jumlahTransaksi] = t
	jumlahTransaksi++
	fmt.Println("Pembayaran berhasil")
}

func cetakTransaksi() {
	var i int
	for i = 0; i < jumlahTransaksi; i++ {
		var t Transaksi
		t = daftarTransaksi[i]
		fmt.Printf("%d. %d -> %d Rp%.2f | %s\n", i+1, t.pengirim, t.penerima, t.jumlah, t.jenis)
	}
}

func editAkun(id int, namaBaru string) {
	var idx int = sequentialSearchAkun(id)
	if idx != -1 {
		daftarAkun[idx].nama = namaBaru
		fmt.Println("Nama akun diperbarui")
	} else {
		fmt.Println("Akun tidak ditemukan")
	}
}

func hapusAkun(id int) {
	var idx int
	idx = sequentialSearchAkun(id)
	if idx != -1 {
		var i int
		for i = idx; i < jumlahAkun-1; i++ {
			daftarAkun[i] = daftarAkun[i+1]
		}
		jumlahAkun--
		fmt.Println("Akun", id, "dihapus")
	} else {
		fmt.Println("Akun tidak ditemukan")
	}
}

func hapusTransaksi(id int) {
	var i int
	i = 0
	for i < jumlahTransaksi {
		if daftarTransaksi[i].pengirim == id {
			var j int
			for j = i; j < jumlahTransaksi-1; j++ {
				daftarTransaksi[j] = daftarTransaksi[j+1]
			}
			jumlahTransaksi--
		} else {
			i++
		}
	}
	fmt.Println("Transaksi dari", id, "dihapus")
}

func insertionSortTransaksi(urutan string) {
	var i int
	for i = 1; i < jumlahTransaksi; i++ {
		var temp Transaksi = daftarTransaksi[i]
		var j int = i - 1
		for j >= 0 && ((urutan == "asc" && daftarTransaksi[j].jumlah > temp.jumlah) || (urutan == "desc" && daftarTransaksi[j].jumlah < temp.jumlah)) {
			daftarTransaksi[j+1] = daftarTransaksi[j]
			j--
		}
		daftarTransaksi[j+1] = temp
	}
}

func selectionSortAkun(id string, nama string) {
	var i, j, idxTerpilih int
	for i = 0; i < jumlahAkun-1; i++ {
		idxTerpilih = i
		for j = i + 1; j < jumlahAkun; j++ {
			if id == "asc" && daftarAkun[j].id < daftarAkun[idxTerpilih].id {
				idxTerpilih = j
			} else if id == "desc" && daftarAkun[j].id > daftarAkun[idxTerpilih].id {
				idxTerpilih = j
			} else if nama == "asc" && daftarAkun[j].nama < daftarAkun[idxTerpilih].nama {
				idxTerpilih = j
			} else if nama == "desc" && daftarAkun[j].nama > daftarAkun[idxTerpilih].nama {
				idxTerpilih = j
			}
		}
		if idxTerpilih != i {
			var temp Akun
			temp = daftarAkun[i]
			daftarAkun[i] = daftarAkun[idxTerpilih]
			daftarAkun[idxTerpilih] = temp
		}
	}
}

func binarySearchAkun(id int) int {
	var low int = 0
	var high int = jumlahAkun - 1
	for low <= high {
		var mid int = (low + high) / 2
		if daftarAkun[mid].id == id {
			return mid
		} else if daftarAkun[mid].id < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func sequentialSearchAkun(id int) int {
	var i int
	for i = 0; i < jumlahSaldo; i++ {
		if daftarAkun[i].id == id {
			return i
		}
	}
	return -1
}

func cetakAkun() {
	if jumlahAkun == 0 {
		fmt.Println("Belum ada akun terdaftar.")
		return
	}
	fmt.Println("Daftar Akun yang Terdaftar:")
	var i int
	for i = 0; i < jumlahAkun; i++ {
		var a Akun
		a = daftarAkun[i]
		fmt.Printf("%d. ID: %d, Nama: %s, Saldo: Rp%.2f, disetujui: %v\n", i+1, a.id, a.nama, a.saldo, a.disetujui)
	}
}

func cetakSaldo(id int, nama string) {
	var idx int = sequentialSearchAkun(id)
	if idx == -1 {
		fmt.Println("Akun tidak ditemukan")
		return
	}
	if daftarAkun[idx].nama != nama {
		fmt.Println("Nama tidak sesuai dengan ID")
		return
	}
	fmt.Printf("Sisa saldo ID: %d, Nama: %s, saldo: Rp%.2f\n", daftarAkun[idx].id, daftarAkun[idx].nama, daftarAkun[idx].saldo)
}

func main() {
	var pilihan int
	var benar bool = true
	for benar {
		fmt.Println("====================")
		fmt.Println("  APLIKASI E-MONEY  ")
		fmt.Println("====================")
		fmt.Println("1. Registrasi Akun")
		fmt.Println("2. Persetujuan Admin")
		fmt.Println("3. Kirim Uang")
		fmt.Println("4. Pembayaran")
		fmt.Println("5. Cetak Transaksi")
		fmt.Println("6. Edit Akun")
		fmt.Println("7. Hapus Akun")
		fmt.Println("8. Hapus Transaksi")
		fmt.Println("9. Sorting Transaksi (Jumlah)")
		fmt.Println("10. Sorting Akun (saldo)")
		fmt.Println("11. Cari Akun (Binary / Sequential Search)")
		fmt.Println("12. Lihat Semua Akun")
		fmt.Println("13. Cek Saldo")
		fmt.Println("0. Keluar")
		fmt.Println("====================")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			var id int
			var nama string
			var saldo float64
			fmt.Print("Masukkan ID: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan Nama: ")
			fmt.Scan(&nama)
			fmt.Print("Masukkan Saldo: ")
			fmt.Scan(&saldo)
			registrasiAkun(id, nama, saldo)

		} else if pilihan == 2 {
			var id int
			var setuju int
			fmt.Print("Masukkan ID akun: ")
			fmt.Scan(&id)
			fmt.Print("Setujui (1 = ya, 0 = tidak): ")
			fmt.Scan(&setuju)
			setujuiAkun(id, setuju == 1)

		} else if pilihan == 3 {
			var pengirim, penerima int
			var jumlah float64
			fmt.Print("ID Pengirim: ")
			fmt.Scan(&pengirim)
			fmt.Print("ID Penerima: ")
			fmt.Scan(&penerima)
			fmt.Print("Jumlah: ")
			fmt.Scan(&jumlah)
			kirimUang(pengirim, penerima, jumlah)

		} else if pilihan == 4 {
			var id int
			var jenis string
			var jumlah float64
			fmt.Print("ID Akun: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan jenis pembayaran (makanan/pulsa/listrik/BPJS): ")
			fmt.Scan(&jenis)
			fmt.Print("Jumlah: ")
			fmt.Scan(&jumlah)
			bayar(id, jenis, jumlah)

		} else if pilihan == 5 {
			cetakTransaksi()

		} else if pilihan == 6 {
			var id int
			var namaBaru string
			fmt.Print("ID Akun: ")
			fmt.Scan(&id)
			fmt.Print("Nama Baru: ")
			fmt.Scan(&namaBaru)
			editAkun(id, namaBaru)

		} else if pilihan == 7 {
			var id int
			fmt.Print("ID Akun: ")
			fmt.Scan(&id)
			hapusAkun(id)

		} else if pilihan == 8 {
			var id int
			fmt.Print("ID Akun Transaksi: ")
			fmt.Scan(&id)
			hapusTransaksi(id)

		} else if pilihan == 9 {
			var urutan string
			fmt.Print("Urutan (asc/desc): ")
			fmt.Scan(&urutan)
			insertionSortTransaksi(urutan)
			fmt.Println("Transaksi sudah diurutkan.")
			cetakTransaksi()

		} else if pilihan == 10 {
			var kolom, urutan string
			fmt.Print("Urut berdasarkan apa? (saldo): ")
			fmt.Scan(&kolom)
			fmt.Print("Urutan berdasarkan? (asc/desc): ")
			fmt.Scan(&urutan)

			if kolom == "id" && (urutan == "asc" || urutan == "desc") {
				selectionSortAkun(urutan, "")
			} else if kolom == "nama" && (urutan == "asc" || urutan == "desc") {
				selectionSortAkun("", urutan)
			} else {
				fmt.Println("Input tidak valid.")
			}

			fmt.Println("Akun sudah diurutkan.")
			var i int
			for i = 0; i < jumlahAkun; i++ {
				fmt.Printf("%d. ID: %d, Nama: %s, Disetujui: %v\n", i+1, daftarAkun[i].id, daftarAkun[i].nama, daftarAkun[i].disetujui)
			}

		} else if pilihan == 11 {
			var id int
			var metode string
			fmt.Print("Masukkan ID Akun: ")
			fmt.Scan(&id)
			fmt.Print("Metode (binary/sequential): ")
			fmt.Scan(&metode)
			var idx int
			idx = -1
			if metode == "binary" {
				idx = binarySearchAkun(id)
			} else {
				idx = sequentialSearchAkun(id)
			}
			if idx != -1 {
				fmt.Printf("Ditemukan: ID: %d, Nama: %s, Saldo: %.2f, Disetujui: %v\n",
					daftarAkun[idx].id, daftarAkun[idx].nama, daftarAkun[idx].saldo, daftarAkun[idx].disetujui)
			} else {
				fmt.Println("Akun tidak ditemukan")
			}

		} else if pilihan == 12 {
			cetakAkun()

		} else if pilihan == 13 {
			var id int
			var nama string
			fmt.Print("Masukkan ID Akun: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan Nama: ")
			fmt.Scan(&nama)
			cetakSaldo(id, nama)

		} else if pilihan == 0 {
			fmt.Println("Keluar dari program.")
			benar = false

		} else {
			fmt.Println("Pilihan tidak valid")
		}
		fmt.Println()
	}
}
