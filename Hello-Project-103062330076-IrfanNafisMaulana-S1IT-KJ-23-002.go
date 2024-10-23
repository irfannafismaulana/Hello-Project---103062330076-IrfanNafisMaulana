package main

// Fachrul Rozi - 103062330074
import (
	"fmt"
)

// Struktur data untuk Mahasiswa
type Mahasiswa struct {
	Nama        string
	NIM         int
	Kelas       string
	MataKuliahs [4]MataKuliah // Setiap mahasiswa memiliki maksimal 4 mata kuliah
}

// Struktur data untuk MataKuliah
type MataKuliah struct {
	Nama  string
	SKS   int
	UTS   float64
	UAS   float64
	Quiz  float64
	Total float64
	Grade string
}

var dataMahasiswa [100]Mahasiswa // Array untuk menyimpan data 100 mahasiswa

// Fungsi untuk menambahkan data mahasiswa
func tambahDataMahasiswa() {
	var mhs Mahasiswa // Membuat variabel mhs untuk menyimpan data mahasiswa sementara
	fmt.Print("Nama: ") // Meminta pengguna memasukkan nama mahasiswa
	fmt.Scan(&mhs.Nama) // Membaca input nama dan menyimpannya ke variabel mhs.Nama
	fmt.Print("NIM: ") // Meminta pengguna memasukkan NIM mahasiswa
	fmt.Scan(&mhs.NIM) // Membaca input NIM dan menyimpannya ke variabel mhs.NIM
	fmt.Print("Kelas: ") // Meminta pengguna memasukkan kelas mahasiswa
	fmt.Scan(&mhs.Kelas) // Membaca input kelas dan menyimpannya ke variabel mhs.Kelas

	index := cariIndeksKosong() // Mencari indeks kosong di array dataMahasiswa untuk menyimpan data baru
	if index == -1 { // Jika tidak ada indeks kosong ditemukan (array penuh)
		fmt.Println("=== Kapasitas data mahasiswa sudah penuh. ===") // Tampilkan pesan bahwa kapasitas sudah penuh
		return // Keluar dari fungsi
	}

	dataMahasiswa[index] = mhs // Menyimpan data mahasiswa baru di indeks kosong yang ditemukan
	fmt.Println("\n=== Data Mahasiswa berhasil ditambahkan ===") // Tampilkan pesan bahwa data mahasiswa berhasil ditambahkan
}

// Fungsi tambahan yang diperlukan (misalnya, untuk mencari indeks kosong) perlu ditambahkan di sini
// Contoh fungsi cariIndeksKosong
func cariIndeksKosong() int {
	for i := 0; i < len(dataMahasiswa); i++ {
		if dataMahasiswa[i].NIM == 0 { // Asumsi NIM 0 menunjukkan slot kosong
			return i
		}
	}
	return -1 // Mengembalikan -1 jika tidak ada slot kosong ditemukan
}


func updateDataMahasiswa() {
	var nim int // Deklarasi variabel nim untuk menyimpan input NIM dari pengguna
	fmt.Print("Masukkan NIM, Untuk Mengupdate Data Mahasiswa: ") // Menampilkan pesan untuk meminta pengguna memasukkan NIM
	fmt.Scan(&nim) // Membaca input NIM dari pengguna dan menyimpannya ke variabel nim

	// Loop melalui seluruh elemen dalam array dataMahasiswa
	for i := 0; i < len(dataMahasiswa); i++ {
		// Memeriksa apakah NIM mahasiswa pada indeks i cocok dengan NIM yang dimasukkan pengguna
		if dataMahasiswa[i].NIM == nim {
			fmt.Print("Nama: ") // Menampilkan pesan untuk meminta pengguna memasukkan nama baru
			fmt.Scan(&dataMahasiswa[i].Nama) // Membaca input nama dan menyimpannya ke variabel Nama mahasiswa pada indeks i
			fmt.Print("Kelas: ") // Menampilkan pesan untuk meminta pengguna memasukkan kelas baru
			fmt.Scan(&dataMahasiswa[i].Kelas) // Membaca input kelas dan menyimpannya ke variabel Kelas mahasiswa pada indeks i
			fmt.Println("=== Data Mahasiswa berhasil diupdate ===") // Menampilkan pesan bahwa data mahasiswa berhasil diupdate
			return // Keluar dari fungsi
		}
	}
	// Jika NIM tidak ditemukan dalam loop, menampilkan pesan bahwa NIM tidak ditemukan
	fmt.Println("=== Maaf NIM tidak ditemukan ===")
}


func deleteDataMahasiswa() {
	var nim int // Deklarasi variabel nim untuk menyimpan input NIM dari pengguna
	fmt.Print("Masukkan NIM, Untuk Menghapus Data Mahasiswa: ") // Menampilkan pesan untuk meminta pengguna memasukkan NIM
	fmt.Scan(&nim) // Membaca input NIM dari pengguna dan menyimpannya ke variabel nim

	// Loop melalui seluruh elemen dalam array dataMahasiswa
	for i := 0; i < len(dataMahasiswa); i++ {
		// Memeriksa apakah NIM mahasiswa pada indeks i cocok dengan NIM yang dimasukkan pengguna
		if dataMahasiswa[i].NIM == nim {
			// Menggeser data mahasiswa setelah indeks i ke kiri untuk mengisi slot yang dihapus
			for j := i; j < len(dataMahasiswa)-1; j++ {
				dataMahasiswa[j] = dataMahasiswa[j+1] // Menggeser elemen ke kiri
			}
			// Mengosongkan data mahasiswa pada slot terakhir karena telah digeser
			dataMahasiswa[len(dataMahasiswa)-1] = Mahasiswa{} // Menetapkan nilai default Mahasiswa
			fmt.Println("=== Data Mahasiswa berhasil dihapus ===") // Menampilkan pesan bahwa data mahasiswa berhasil dihapus
			return // Keluar dari fungsi
		}
	}
	// Jika NIM tidak ditemukan dalam loop, menampilkan pesan bahwa NIM tidak ditemukan
	fmt.Println("=== Maaf NIM tidak ditemukan ===")
}
func deleteDataMataKuliah() {
	var nim int // Deklarasi variabel nim untuk menyimpan input NIM dari pengguna
	fmt.Print("Masukkan NIM, Untuk Menghapus Mata Kuliah Mahasiswa: ") // Menampilkan pesan untuk meminta pengguna memasukkan NIM
	fmt.Scan(&nim) // Membaca input NIM dari pengguna dan menyimpannya ke variabel nim

	// Loop melalui seluruh elemen dalam array dataMahasiswa
	for i := 0; i < len(dataMahasiswa); i++ {
		// Memeriksa apakah NIM mahasiswa pada indeks i cocok dengan NIM yang dimasukkan pengguna
		if dataMahasiswa[i].NIM == nim {
			fmt.Print("Nama Mata Kuliah yang akan dihapus: ") // Menampilkan pesan untuk meminta pengguna memasukkan nama mata kuliah yang akan dihapus
			var namaMK string // Deklarasi variabel namaMK untuk menyimpan input nama mata kuliah
			fmt.Scan(&namaMK) // Membaca input nama mata kuliah dari pengguna dan menyimpannya ke variabel namaMK

			// Loop melalui seluruh mata kuliah mahasiswa pada indeks i
			for j := 0; j < len(dataMahasiswa[i].MataKuliahs); j++ {
				// Memeriksa apakah nama mata kuliah pada indeks j cocok dengan nama mata kuliah yang dimasukkan pengguna
				if dataMahasiswa[i].MataKuliahs[j].Nama == namaMK {
					// Menggeser data mata kuliah setelah indeks j ke kiri untuk mengisi slot yang dihapus
					for k := j; k < len(dataMahasiswa[i].MataKuliahs)-1; k++ {
						dataMahasiswa[i].MataKuliahs[k] = dataMahasiswa[i].MataKuliahs[k+1] // Menggeser elemen ke kiri
					}
					// Mengosongkan data mata kuliah pada slot terakhir karena telah digeser
					dataMahasiswa[i].MataKuliahs[len(dataMahasiswa[i].MataKuliahs)-1] = MataKuliah{} // Menetapkan nilai default MataKuliah
					fmt.Println("=== Mata Kuliah berhasil dihapus ===") // Menampilkan pesan bahwa mata kuliah berhasil dihapus
					return // Keluar dari fungsi
				}
			}
			// Jika nama mata kuliah tidak ditemukan dalam loop, menampilkan pesan bahwa nama mata kuliah tidak ditemukan
			fmt.Println("\n=== Maaf Nama Mata Kuliah tidak ditemukan ===")
			return // Keluar dari fungsi
		}
	}
	// Jika NIM tidak ditemukan dalam loop, menampilkan pesan bahwa NIM tidak ditemukan
	fmt.Println("=== Maaf NIM tidak ditemukan ===")
}
func tambahDataMataKuliah() {
	var nim int // Deklarasi variabel nim untuk menyimpan input NIM dari pengguna
	fmt.Print("Masukkan NIM, Untuk Menambah Mata Kuliah Mahasiswa: ") // Menampilkan pesan untuk meminta pengguna memasukkan NIM
	fmt.Scan(&nim) // Membaca input NIM dari pengguna dan menyimpannya ke variabel nim

	// Loop melalui seluruh elemen dalam array dataMahasiswa
	for i := 0; i < len(dataMahasiswa); i++ {
		// Memeriksa apakah NIM mahasiswa pada indeks i cocok dengan NIM yang dimasukkan pengguna
		if dataMahasiswa[i].NIM == nim {
			// Loop melalui seluruh mata kuliah mahasiswa pada indeks i
			for j := 0; j < len(dataMahasiswa[i].MataKuliahs); j++ {
				// Memeriksa apakah ada slot kosong untuk menambahkan mata kuliah baru
				if dataMahasiswa[i].MataKuliahs[j].Nama == "" {
					var matkul MataKuliah // Deklarasi variabel matkul untuk menyimpan data mata kuliah baru
					fmt.Print("Nama Mata Kuliah: ") // Menampilkan pesan untuk meminta pengguna memasukkan nama mata kuliah
					fmt.Scan(&matkul.Nama) // Membaca input nama mata kuliah dan menyimpannya ke variabel Nama dalam struct matkul
					fmt.Print("SKS: ") // Menampilkan pesan untuk meminta pengguna memasukkan jumlah SKS
					fmt.Scan(&matkul.SKS) // Membaca input jumlah SKS dan menyimpannya ke variabel SKS dalam struct matkul
					fmt.Print("Nilai UTS: ") // Menampilkan pesan untuk meminta pengguna memasukkan nilai UTS
					fmt.Scan(&matkul.UTS) // Membaca input nilai UTS dan menyimpannya ke variabel UTS dalam struct matkul
					fmt.Print("Nilai UAS: ") // Menampilkan pesan untuk meminta pengguna memasukkan nilai UAS
					fmt.Scan(&matkul.UAS) // Membaca input nilai UAS dan menyimpannya ke variabel UAS dalam struct matkul
					fmt.Print("Nilai Quiz: ") // Menampilkan pesan untuk meminta pengguna memasukkan nilai Quiz
					fmt.Scan(&matkul.Quiz) // Membaca input nilai Quiz dan menyimpannya ke variabel Quiz dalam struct matkul

					hitungTotalDanGrade(&matkul) // Memanggil fungsi hitungTotalDanGrade untuk menghitung total dan grade mata kuliah
					dataMahasiswa[i].MataKuliahs[j] = matkul // Menyimpan data mata kuliah baru pada slot yang ditemukan
					fmt.Println("=== Mata Kuliah berhasil ditambahkan ===") // Menampilkan pesan bahwa mata kuliah berhasil ditambahkan
					return // Keluar dari fungsi
				}
			}
			// Jika tidak ada ruang kosong untuk menambahkan mata kuliah baru
			fmt.Println("=== Maaf, sudah tidak ada ruang kosong untuk menambahkan mata kuliah ===")
			return // Keluar dari fungsi
		}
	}
	// Jika NIM tidak ditemukan dalam loop, menampilkan pesan bahwa NIM tidak ditemukan
	fmt.Println("=== Maaf NIM tidak ditemukan ===")
}

func updateDataMataKuliah() {
	var nim int // Deklarasi variabel nim untuk menyimpan input NIM dari pengguna
	fmt.Print("Masukkan NIM, Untuk Mengubah Mata Kuliah Mahasiswa: ") // Menampilkan pesan untuk meminta pengguna memasukkan NIM
	fmt.Scan(&nim) // Membaca input NIM dari pengguna dan menyimpannya ke variabel nim

	// Loop melalui seluruh elemen dalam array dataMahasiswa
	for i := 0; i < len(dataMahasiswa); i++ {
		// Memeriksa apakah NIM mahasiswa pada indeks i cocok dengan NIM yang dimasukkan pengguna
		if dataMahasiswa[i].NIM == nim {
			fmt.Print("Nama Mata Kuliah yang akan diubah: ") // Menampilkan pesan untuk meminta pengguna memasukkan nama mata kuliah yang akan diubah
			var oldNama string // Deklarasi variabel oldNama untuk menyimpan input nama mata kuliah lama
			fmt.Scan(&oldNama) // Membaca input nama mata kuliah lama dari pengguna dan menyimpannya ke variabel oldNama

			// Loop melalui seluruh mata kuliah mahasiswa pada indeks i
			for j := 0; j < len(dataMahasiswa[i].MataKuliahs); j++ {
				// Memeriksa apakah nama mata kuliah pada indeks j cocok dengan nama mata kuliah lama yang dimasukkan pengguna
				if dataMahasiswa[i].MataKuliahs[j].Nama == oldNama {
					var matkul MataKuliah // Deklarasi variabel matkul untuk menyimpan data mata kuliah baru
					fmt.Print("Nama Mata Kuliah Baru: ") // Menampilkan pesan untuk meminta pengguna memasukkan nama mata kuliah baru
					fmt.Scan(&matkul.Nama) // Membaca input nama mata kuliah baru dan menyimpannya ke variabel Nama dalam struct matkul
					fmt.Print("SKS: ") // Menampilkan pesan untuk meminta pengguna memasukkan jumlah SKS baru
					fmt.Scan(&matkul.SKS) // Membaca input jumlah SKS baru dan menyimpannya ke variabel SKS dalam struct matkul
					fmt.Print("Nilai UTS: ") // Menampilkan pesan untuk meminta pengguna memasukkan nilai UTS baru
					fmt.Scan(&matkul.UTS) // Membaca input nilai UTS baru dan menyimpannya ke variabel UTS dalam struct matkul
					fmt.Print("Nilai UAS: ") // Menampilkan pesan untuk meminta pengguna memasukkan nilai UAS baru
					fmt.Scan(&matkul.UAS) // Membaca input nilai UAS baru dan menyimpannya ke variabel UAS dalam struct matkul
					fmt.Print("Nilai Quiz: ") // Menampilkan pesan untuk meminta pengguna memasukkan nilai Quiz baru
					fmt.Scan(&matkul.Quiz) // Membaca input nilai Quiz baru dan menyimpannya ke variabel Quiz dalam struct matkul

					hitungTotalDanGrade(&matkul) // Memanggil fungsi hitungTotalDanGrade untuk menghitung total dan grade mata kuliah baru
					dataMahasiswa[i].MataKuliahs[j] = matkul // Menyimpan data mata kuliah baru pada slot yang ditemukan
					fmt.Println("=== Mata Kuliah berhasil diubah ===") // Menampilkan pesan bahwa mata kuliah berhasil diubah
					return // Keluar dari fungsi
				}
			}
			fmt.Println("=== Maaf Nama Mata Kuliah tidak ditemukan ===") // Menampilkan pesan bahwa nama mata kuliah lama tidak ditemukan
			return // Keluar dari fungsi
		}
	}
	fmt.Println("=== Maaf NIM tidak ditemukan ===") // Menampilkan pesan bahwa NIM tidak ditemukan
}

func cariMahasiswaByMataKuliah() {
	var mataKuliah string // Variabel untuk menyimpan nama mata kuliah yang akan dicari
	fmt.Print("Masukkan Nama Mata Kuliah: ") // Menampilkan pesan untuk meminta pengguna memasukkan nama mata kuliah
	fmt.Scan(&mataKuliah) // Membaca input nama mata kuliah dari pengguna dan menyimpannya ke variabel mataKuliah
	fmt.Printf("\nDaftar Mahasiswa yang Mengambil Mata Kuliah %s:\n", mataKuliah) // Menampilkan pesan dengan nama mata kuliah yang dicari
	for _, mahasiswa := range dataMahasiswa { // Loop melalui setiap mahasiswa dalam array dataMahasiswa
		for _, mk := range mahasiswa.MataKuliahs { // Loop melalui setiap mata kuliah mahasiswa saat ini
			if mk.Nama == mataKuliah { // Memeriksa apakah nama mata kuliah cocok dengan mata kuliah yang dicari
				fmt.Printf("Nama: %s, NIM: %d, Kelas: %s\n", mahasiswa.Nama, mahasiswa.NIM, mahasiswa.Kelas) // Menampilkan informasi mahasiswa yang mengambil mata kuliah tersebut
				break // Keluar dari loop untuk mata kuliah mahasiswa saat ini
			}
		}
	}
}

// Insertion Sort Descending
func urutMahasiswaByNilaiSKSInsertion() {
	for i := 1; i < len(dataMahasiswa); i++ { // Loop dari indeks kedua hingga akhir array dataMahasiswa
		key := dataMahasiswa[i] // Simpan data mahasiswa pada indeks i sebagai kunci
		j := i - 1 // Inisialisasi indeks j sebagai indeks sebelumnya

		// Bandingkan berdasarkan kriteria Descending
		for j >= 0 && (hitungTotalSKS(dataMahasiswa[j]) < hitungTotalSKS(key) || 
			hitungTotalSKS(dataMahasiswa[j]) == hitungTotalSKS(key) && hitungTotalNilai(dataMahasiswa[j]) < hitungTotalNilai(key)) { // Bandingkan total SKS dan total nilai
			dataMahasiswa[j+1] = dataMahasiswa[j] // Geser data ke kanan
			j = j - 1 // Pindah ke indeks sebelumnya
		}
		dataMahasiswa[j+1] = key // Letakkan kunci di posisi yang benar
	}
}

// Fungsi untuk menghitung total nilai dari mata kuliah yang diambil oleh seorang mahasiswa
func hitungTotalNilai(m Mahasiswa) float64 {
	var totalNilai float64 // Variabel untuk menyimpan total nilai
	for _, mk := range m.MataKuliahs { // Loop melalui setiap mata kuliah mahasiswa
		totalNilai += mk.Total // Tambahkan nilai total mata kuliah saat ini ke total nilai
	}
	return totalNilai // Kembalikan total nilai
}

// Fungsi untuk menghitung total SKS dari mata kuliah yang diambil oleh seorang mahasiswa
func hitungTotalSKS(m Mahasiswa) int {
	var totalSKS int // Variabel untuk menyimpan total SKS
	for _, mk := range m.MataKuliahs { // Loop melalui setiap mata kuliah mahasiswa
		totalSKS += mk.SKS // Tambahkan SKS mata kuliah saat ini ke total SKS
	}
	return totalSKS // Kembalikan total SKS
}

// Fungsi untuk menghitung total nilai dan menentukan grade dari suatu mata kuliah
func hitungTotalDanGrade(mk *MataKuliah) {
	// Dalam fungsi ini, UTS diberi bobot 30%, UAS 40%, dan Quiz 30%.
	mk.Total = mk.UTS*0.3 + mk.UAS*0.4 + mk.Quiz*0.3 // Hitung total nilai berdasarkan bobot
	if mk.Total >= 85 { // Jika total nilai lebih besar atau sama dengan 85
		mk.Grade = "A" // Set grade menjadi A
	} else if mk.Total >= 70 { // Jika total nilai lebih besar atau sama dengan 70
		mk.Grade = "B" // Set grade menjadi B
	} else if mk.Total >= 55 { // Jika total nilai lebih besar atau sama dengan 55
		mk.Grade = "C" // Set grade menjadi C
	} else if mk.Total >= 40 { // Jika total nilai lebih besar atau sama dengan 40
		mk.Grade = "D" // Set grade menjadi D
	} else { // Jika total nilai kurang dari 40
		mk.Grade = "E" // Set grade menjadi E
	}
}

// Fungsi untuk menampilkan data mahasiswa beserta mata kuliah yang diambilnya
func tampilkanDataMahasiswa() {
	fmt.Println("\nData Mahasiswa:") // Menampilkan header untuk data mahasiswa
	for _, mahasiswa := range dataMahasiswa { // Loop melalui setiap mahasiswa dalam array dataMahasiswa
		if mahasiswa.NIM != 0 { // Memeriksa apakah NIM mahasiswa tidak kosong
			fmt.Printf("Nama: %s, NIM: %d, Kelas: %s\n", mahasiswa.Nama, mahasiswa.NIM, mahasiswa.Kelas) // Menampilkan informasi mahasiswa
			for _, mk := range mahasiswa.MataKuliahs { // Loop melalui setiap mata kuliah mahasiswa
				if mk.Nama != "" { // Memeriksa apakah nama mata kuliah tidak kosong
					fmt.Printf("  Mata Kuliah: %s, SKS: %d, UTS: %.2f, UAS: %.2f, Quiz: %.2f, Total: %.2f, Grade: %s\n",
						mk.Nama, mk.SKS, mk.UTS, mk.UAS, mk.Quiz, mk.Total, mk.Grade) // Menampilkan informasi mata kuliah
				}
			}
		}
	}
}

func main() {
	var pilihan int // Variabel untuk menyimpan pilihan menu yang dimasukkan oleh pengguna
	for { // Loop tak terbatas
		fmt.Println("\nMenu:") // Menampilkan menu opsi
		fmt.Println("1. Tambah Data Mahasiswa")
		fmt.Println("2. Update Data Mahasiswa")
		fmt.Println("3. Hapus Data Mahasiswa")
		fmt.Println("4. Tambah Mata Kuliah Mahasiswa")
		fmt.Println("5. Update Mata Kuliah Mahasiswa")
		fmt.Println("6. Hapus Mata Kuliah Mahasiswa")
		fmt.Println("7. Tampilkan Daftar Mahasiswa yang Mengambil Mata Kuliah")
		fmt.Println("8. Tampilkan Daftar Mahasiswa Terurut berdasarkan Nilai dan Jumlah SKS (Descending)")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ") // Meminta pengguna untuk memilih menu
		fmt.Scan(&pilihan) // Membaca pilihan menu yang dimasukkan oleh pengguna

		switch pilihan { // Memproses pilihan menu
		case 1: // Jika pilihan adalah 1
			tambahDataMahasiswa() // Panggil fungsi untuk menambah data mahasiswa
		case 2: // Jika pilihan adalah 2
			updateDataMahasiswa() // Panggil fungsi untuk memperbarui data mahasiswa
		case 3: // Jika pilihan adalah 3
			deleteDataMahasiswa() // Panggil fungsi untuk menghapus data mahasiswa
		case 4: // Jika pilihan adalah 4
			tambahDataMataKuliah() // Panggil fungsi untuk menambah mata kuliah mahasiswa
		case 5: // Jika pilihan adalah 5
			updateDataMataKuliah() // Panggil fungsi untuk memperbarui mata kuliah mahasiswa
		case 6: // Jika pilihan adalah 6
			deleteDataMataKuliah() // Panggil fungsi untuk menghapus mata kuliah mahasiswa
		case 7: // Jika pilihan adalah 7
			cariMahasiswaByMataKuliah() // Panggil fungsi untuk mencari mahasiswa berdasarkan mata kuliah
		case 8: // Jika pilihan adalah 8
			urutMahasiswaByNilaiSKSInsertion() // Panggil fungsi untuk mengurutkan mahasiswa berdasarkan nilai dan jumlah SKS (Descending)
			tampilkanDataMahasiswa() // Panggil fungsi untuk menampilkan data mahasiswa
		case 0: // Jika pilihan adalah 0
			return // Keluar dari program
		default: // Jika pilihan tidak valid
			fmt.Println("Pilihan tidak valid") // Tampilkan pesan kesalahan
		}
	}
}
