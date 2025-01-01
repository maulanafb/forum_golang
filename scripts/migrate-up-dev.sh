#!/bin/bash

# Konfigurasi database untuk lingkungan pengembangan
DATABASE_USERNAME="root"
DATABASE_PASSWORD="Docker123!"
DATABASE_NAME="forumdb"
DATABASE_HOST="localhost" # Digunakan jika Anda ingin mengubah ke TCP
DATABASE_PORT="3306"      # Digunakan jika Anda ingin mengubah ke TCP

DATABASE_URL="mysql://${DATABASE_USERNAME}:${DATABASE_PASSWORD}@/${DATABASE_NAME}?parseTime=true"

# Debugging: Cetak URL
echo "URL: $DATABASE_URL"

# Export URL ke environment
export DATABASE_URL=$DATABASE_URL

# Jalankan migrasi ke atas
echo "Menjalankan migrasi database (development)..."
dbmate --url "$DATABASE_URL" up

if [ $? -eq 0 ]; then
  echo "Migrasi berhasil!"
else
  echo "Migrasi gagal. Periksa konfigurasi atau file migrasi Anda."
  exit 1
fi
