# 🚗 Parking Lot System - VocaGames Backend Submission

Sistem manajemen parkir sederhana yang diimplementasikan menggunakan Go untuk submission Backend VocaGames.

## 📁 Struktur Proyek

```
Submission_Backend_VocaGames/
├── main.go          # File utama aplikasi
├── input.txt        # File input untuk testing
├── go.mod           # Go module configuration
└── README.md        # Dokumentasi proyek
```

## 🛠️ Instalasi dan Setup

### Prerequisites
- **Go 1.19+** harus terinstall di sistem
- Terminal/Command Prompt

### Langkah Instalasi

#### 1. Clone Repository
```bash
git clone https://github.com/[username]/Submission_Backend_VocaGames.git
cd Submission_Backend_VocaGames
```

#### 2. Verifikasi Go Installation
```bash
go version
```

#### 3. Initialize Go Module (jika belum ada)
```bash
go mod init parking-lot
```

#### 4. Download Dependencies
```bash
go mod tidy
```

## 🚀 Cara Menjalankan

### Menjalankan dengan File Input
```bash
go run main.go input.txt
```

### Menjalankan dengan Input Custom
1. Buat file input baru (misal: `custom_input.txt`)
2. Jalankan:
```bash
go run main.go custom_input.txt
```

### Build Executable
```bash
go build -o parking-lot main.go
./parking-lot input.txt
```

## 📝 Format Input

File input harus berisi perintah-perintah berikut:

```
create_parking_lot 6
park KA-01-HH-1234 White
park KA-01-HH-9999 Red
leave KA-01-HH-3141 4
status
```

### Perintah yang Tersedia:

| Perintah | Deskripsi |
|----------|-----------|
| `create_parking_lot <capacity>` | Membuat parkir dengan kapasitas N |
| `park <registration_number> <color>` | Parkir kendaraan |
| `leave <registration_number> <hours>` | Keluarkan kendaraan |
| `status` | Tampilkan status parkir |

## 💡 Contoh Penggunaan

```bash
# Buat parking lot dengan kapasitas 6
create_parking_lot 6

# Parkir beberapa kendaraan
park KA-01-HH-1234 White
park KA-01-HH-9999 Red
park KA-01-BB-0001 Black

# Cek status
status

# Keluarkan kendaraan setelah 4 jam
leave KA-01-HH-1234 4
```

## 🤝 Contributing

1. Fork repository
2. Buat feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit changes (`git commit -m 'Add some AmazingFeature'`)
4. Push ke branch (`git push origin feature/AmazingFeature`)
5. Buat Pull Request

## 📄 License

Project ini dibuat untuk submission Backend VocaGames.

## 👨‍💻 Author

**Dimas Mulya**
- Email: [your-email@example.com]
- GitHub: [@your-username]

## 📞 Support

Jika ada pertanyaan atau issue, silakan:
- Buat issue di repository ini
- Hubungi melalui email

---

*Made with ❤️ for VocaGames Backend Submission*
