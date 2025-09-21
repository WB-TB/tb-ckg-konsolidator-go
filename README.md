# CKG TB Konsolidator (Golang)

<p align="center">
<img src="https://go.dev/blog/go-brand/Go-Logo/SVG/Go-Logo_Blue.svg" height="120">
<img src="https://echo.labstack.com/img/logo-light.svg" height="80"><br>
Komponen Server untuk mendukung interoperabilitas data Cek Kesehatan Gratis (CKG) dengan Sistem Informasi Tuberkulosi (SITB). Menjembatani pertukaran data dua arah terkait kasus TB melalui ekosistem SATUSEHAT/ASIK (Aplikasi Sehat IndonesiaKu), memungkinkan skrining, pemeriksaan dan pemanfaatan data melalui Dashboard dapat terkonsolidasi dalam satu jembatan pertukaran data yang digunakan bersama.
  
  Sistem ini merupakan draft awal infrastruktur pertukaran data yang disiapkan menggunkan bahasa pemrograman <a href="https://go.dev/" target="_blank">Golang</a> untuk mendapatkan performa tinggi, efesiensi, skalabilitas dan fleksibilitas aplikasi server nyesuaikan kebutuhan pengembangan pertukaran data SATUSEHAT dimasa depan.</p>

## Desain Interoperabilitas Data
![Integrasi0](https://github.com/WB-TB/tb-ckg-konsolidator/raw/main/documentation/assets/images/ckg-tb-0.png)
Pengembangan sistem interoperabilitas ini seperti tergambar di atas adalah untuk mengadakan pertukaran data antara ASIK dengan SITB. Ada tiga objektif yang ingin dicapai dari desain konsolidasi data TB ini.
### 1. Pertukaran data (dua arah) Skrining TB antara ASIK-CKG dengan Aplikasi SITB
![Integrasi1](https://github.com/WB-TB/tb-ckg-konsolidator/raw/main/documentation/assets/images/ckg-tb-1.png)
### 2. Penyajian Dashboard terkait Program TB yang akurat
![Integrasi2](https://github.com/WB-TB/tb-ckg-konsolidator/raw/main/documentation/assets/images/ckg-tb-2.png)
### 3. Menjembatani pertukaran data Rekam Medis Elektronik (RME) dari sistem di Fasilitas  Kesehatan Kesehatan dengan ekosistem pencatatan TB
![Integrasi3](https://github.com/WB-TB/tb-ckg-konsolidator/raw/main/documentation/assets/images/ckg-tb-3.png)
  <p></p>

# Pengembagan
## Tahapan Pengembangan Jangka Pendek
Pengembangan ini akan dibagi menjadi 3 tahap:
1. *Tahap 1*: Mewujudkan terjadinya interoperabilitas antara SITB dengan ASIK-CKG menggunakan metadata yang ada saat ini.
2. *Tahap 2*: Membangun sistem notifikasi yang data TB yang tersedia di SITB dan ASIK-CKG untuk kebutuhan pemberitahuan kepada pasien maupun komunitas.
3. *Tahap 3*: Membangun dashboard pelaporan Tuberkuloasis yang akurat berdasarkan hasil konsolidasi ASIK-CKG/SATUSEHAT Mobile dan SITB.
![Integrasi4](https://github.com/WB-TB/tb-ckg-konsolidator/raw/main/documentation/assets/images/ckg-tb-4.png)

## Tahapan Pengembangan Jangka Panjang
Sesuai dengan objektif ketiga di atas dengan membangun pertukaran data Rekam Medis Elektronik (RME) dari sistem di Fasilitas Kesehatan terkait TB.

## Spesifikasi Sistem
Framework [Echo](https://echo.labstack.com/) digunakan dalam project ini untuk mengoptimalkan interoperabilitas data layanan skrining TB.

## Setup Web Service
Sebelum menjalankan aplikasi pastikan <a href="https://go.dev/dl/" target="_blank">Go</a> sudah terinstall. Lalu jalankan
```bash
go mod tidy
```

Pastikan variabel environment telah tersedia baik melalui export langsung di shell environment maupun memuat file .env. Berikut variabel yang diharapkan tersedia
```bash
environment=dev
HTTP_PORT=6767
HTTP_SERVER_TIMEOUT=180
MONGO_CONNECTION_STRING=mongodb://*****:ckg_tb@localhost:27017/
MONGO_DBNAME=konsolidator
MONGO_COLLECTION_NAME=abc
UNDER_MAINTENANCE=false
RATE_LIMIT=60
ENTRYPOINT_MSG=ABCDEF
API_KEY_SECRET=******

# Konfigurasi untuk CKG TB
CKG_TB_USE_CACHE=true
CKG_TB_MONGO_DBNAME=ckgtb
CKG_TB_MONGO_COLLECTION_TRANSACTION=transaction
CKG_TB_GET_DATA_PAGE_SIZE=1000
CKG_TB_GET_DATA_TIMEOUT=60
CKG_TB_POST_DATA_TIMEOUT=180
```

## Jalankan mode development
```bash
go run main.go
```

# Deployment
Deployment disesuaikan dengan jenis infrastruktur yang digunakan
