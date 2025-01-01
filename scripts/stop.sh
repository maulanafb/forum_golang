#!/bin/bash

# Hentikan aplikasi berdasarkan PID atau port
if [ -f ./pid.file ]; then
  PID=$(cat ./pid.file)
  echo "Menghentikan aplikasi dengan PID $PID..."

  # Coba hentikan proses
  if kill $PID > /dev/null 2>&1; then
    rm -f ./pid.file
    echo "Aplikasi telah dihentikan."
  else
    echo "Gagal menghentikan aplikasi. Proses mungkin sudah tidak aktif."
  fi
else
  echo "PID file tidak ditemukan. Mencoba menghentikan proses di port 8080..."
fi

# Hentikan proses berdasarkan port 8080
if lsof -i :8080 > /dev/null 2>&1; then
  sudo fuser -k 8080/tcp
  echo "Proses yang menggunakan port 8080 telah dihentikan."
else
  echo "Tidak ada proses yang menggunakan port 8080."
fi
