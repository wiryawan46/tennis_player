# Tennis Player

Project ini merupakan project yang dibuat untuk melakukan pre tes di Evermos

## Mulai

Berikut ini project sederhana untuk kasus tennis player. Bahasa yang digunakan adalah Go.

### Persiapan

Beberapa yang perlu dilakukan sebelum menjalankan service ini

```
1. Buat USER pada postgresql dengan nama 'postgres' password 'Bismillah123'
2. Import Database sql yang sudah disediakan
3. Jika ingin merubah config db cukup dirubah di file .env yang telah disediakan
4. Lanjut ke step berikutnya
```

### Menjalankan service

1. Silahkan mengikuti intruksi menginstall Go kedalam OS anda pada laman berikut [ini] (https://golang.org/doc/install)

2. Masuk ke direktori project lalu jalankan binary file dengan cara 
```
go run tennis-player
```

### Cara lain
Jika anda ingin mencoba menjalankan projectnya tanpa binary file lakukan hal berikut ini

1. Masuk ke folder project jalankan perintah berikut ini
```
go mod
```

2. Setelah selesai mengunduh dependency yang diperlukan selanjutnya compile project dengan perintah berikut
```
go build
```

3. File binary dari nama project sudah terbentuk, selanjunya jalankan file binary tersebut
```
go run nama_binary_file
```

**NB** : nama binary file sesuai dengan nama project

## Mencoba service

Untuk mencoba service silahkan import file json postman-collection ke aplikasi postman
Id player yang disediakan 1-10 dan untuk container ball 1-5

## Dibangun dengan

* [Echo Framework](https://echo.labstack.com/) - Web framework yang digunakan

## Penulis

* **Manggala Pramuditya Wiryawan** - *Inisial* - [Wiryawan46](https://github.com/wiryawan46)
