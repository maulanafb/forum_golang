# Forum Golang Application

Aplikasi forum ini dibangun menggunakan Golang dengan berbagai fitur seperti signup, login, komentar pada postingan, dan fitur CRUD untuk manajemen postingan. Aplikasi ini juga menggunakan **JWT (JSON Web Tokens)** untuk autentikasi dan **Refresh Token** untuk menjaga sesi pengguna.

## Fitur

- **Signup**: Pengguna dapat mendaftar dengan email dan password.
- **Login**: Pengguna dapat login dengan menggunakan email dan password.
- **CRUD Post**: Pengguna dapat membuat, membaca, memperbarui, dan menghapus postingan.
- **Komentar**: Pengguna dapat memberikan komentar pada setiap postingan.
- **JWT Authentication**: Menggunakan JSON Web Token untuk autentikasi dan sesi.
- **Refresh Token**: Menggunakan refresh token untuk memperpanjang masa aktif session.

## Teknologi

- **Golang**: Bahasa pemrograman utama untuk aplikasi ini.
- **GORM**: ORM untuk komunikasi dengan database (PostgreSQL).
- **JWT**: Untuk autentikasi dan sesi.
- **Gin Gonic**: Web framework untuk Golang.
- **PostgreSQL**: Database yang digunakan untuk menyimpan data.
- **Docker**: Untuk containerization aplikasi.