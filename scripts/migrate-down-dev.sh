#!/bin/bash

# Konfigurasi database untuk lingkungan pengembangan
DATABASE_USERNAME="root"
DATABASE_PASSWORD="Docker123!"
DATABASE_NAME="forumdb"
DATABASE_HOST="localhost" # Digunakan jika Anda ingin mengubah ke TCP
DATABASE_PORT="3306"      # Digunakan jika Anda ingin mengubah ke TCP

DATABASE_URL="mysql://${DATABASE_USERNAME}:${DATABASE_PASSWORD}@/${DATABASE_NAME}?parseTime=true"

# Export URL ke environment
# export DATABASE_URL=$DATABASE_URL

# Jalankan migrasi ke bawah
echo "Membatalkan migrasi terakhir (development)..."
dbmate --url $DATABASE_URL down

if [ $? -eq 0 ]; then
  echo "Migrasi terakhir berhasil dibatalkan!"
else
  echo "Gagal membatalkan migrasi. Periksa konfigurasi atau file migrasi Anda."
  exit 1
fi
