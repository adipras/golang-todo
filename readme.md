# Readme, please.

## Instalasi golang

Ikuti instruksi ini untuk melakukan instalasi golang: [Tutorial instalasi golang](https://go.dev/doc/install)

Versi golang yang dipakai pada boilerplate ini: **go version go1.16.4 darwin/amd64**. Disarankan menggunakan versi golang yang sama atau lebih baru.

Bila ingin memperbarui versi golang. Berikut cara memperbarui versi golang:

```bash
git clone https://github.com/udhos/update-golang
cd update-golang
sudo ./update-golang.sh
```

Ketika cloning dan error, jalankan perintah:

```bash
go mod tidy
```

## Menjalankan service secara lokal

Jika kamu ingin menjalankan service ini secara lokal dengan database lokal, kamu tinggal copy **run_example.sh** ke **run_local.sh**.

```bash
cp run_example.sh run_local.sh
```

Ubah variabel dalam **run_local.sh**.

```bash
export DB_USERNAME="USERNAME_MYSQL_KAMU"
export DB_PASSWORD="PASSWORD_MYSQL_KAMU"
export DB_NAME="NAMA_DATABASE_MYSQL_KAMU"
```

Jalankan shell script untuk menjalankan service

```bash
sh ./run_local.sh
```

Jangan lupa, kamu perlu melakukan instalasi **nodemon** terlebih dahulu. Lihat [Tutorial instalasi nodemon](https://www.npmjs.com/package/nodemon) ini bila kamu belum mengetahui cara instalasi nodemon.

```bash
npm install -g nodemon
```