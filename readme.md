Submission_Backend_VocaGames/ ├── main.go # File utama aplikasi ├── input.txt # File input untuk testing ├── go.mod # Go module configuration └── README.md # Dokumentasi proyek

plaintext

## 🛠️ Instalasi dan Setup### Prerequisites- **Go 1.19+** harus terinstall di sistem- Terminal/Command Prompt### Langkah Instalasi#### 1. Clone Repository```bashgit clone https://github.com/[username]/Submission_Backend_VocaGames.gitcd Submission_Backend_VocaGames
2. Verifikasi Go Installation
bash
Run
go version
3. Initialize Go Module (jika belum ada)
bash
Run
go mod init parking-lot
4. Download Dependencies
bash
Run
go mod tidy
🚀 Cara Menjalankan
Menjalankan dengan File Input
bash
Run
go run main.go input.txt
Menjalankan dengan Input Custom
Buat file input baru (misal: custom_input.txt)
Jalankan:
bash
Run
go run main.go custom_input.txt
Build Executable
bash
Run
go build -o parking-lot main.go./parking-lot input.txt
📝 Format Input
File input harus berisi perintah-perintah berikut:

plaintext

create_parking_lot 6park KA-01-HH-1234 Whitepark KA-01-HH-9999 Redleave KA-01-HH-3141 4status
Perintah yang Tersedia:
create_parking_lot <capacity> - Membuat parkir dengan kapasitas N
park <registration_number> <color> - Parkir kendaraan
leave <registration_number> <hours> - Keluarkan kendaraan
status - Tampilkan status parkir

## 🤝 Contributing
1. Fork repository
2. Buat feature branch ( git checkout -b feature/AmazingFeature )
3. Commit changes ( git commit -m 'Add some AmazingFeature' )
4. Push ke branch ( git push origin feature/AmazingFeature )
5. Buat Pull Request
## 📄 License
Project ini dibuat untuk submission Backend VocaGames.

## 👨‍💻 Author
Dimas Mulya

- Email: [ your-email@example.com ]
- GitHub: [@your-username]
## 📞 Support
Jika ada pertanyaan atau issue, silakan buat issue di repository ini atau hubungi melalui email.

Made with ❤️ for VocaGames Backend Submission