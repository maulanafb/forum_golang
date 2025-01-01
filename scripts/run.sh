#!/bin/bash

# Jalankan aplikasi Go dengan logging
echo "Menjalankan aplikasi Go..."
nohup go run cmd/main.go > log.txt 2>&1 &

# Ambil PID dari proses utama
APP_PID=$!

# Tunggu beberapa detik agar aplikasi memulai server
sleep 2

# Cari proses yang mendengarkan port 8080
PORT_PID=$(lsof -ti :8080)

# Simpan PID ke file (prioritaskan PORT_PID jika ada)
if [ -n "$PORT_PID" ]; then
    echo $PORT_PID > ./pid.file
    echo "Aplikasi berjalan dengan PID $PORT_PID. Lihat log.txt untuk output."
else
    echo $APP_PID > ./pid.file
    echo "Aplikasi berjalan dengan PID $APP_PID. Lihat log.txt untuk output."
fi
