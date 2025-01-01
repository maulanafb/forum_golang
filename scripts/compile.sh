#!/bin/bash

# Compile aplikasi Go
echo "Meng-compile aplikasi Go..."
go build -o app cmd/main.go

if [ $? -eq 0 ]; then
  echo "Aplikasi berhasil di-compile menjadi 'app'."
else
  echo "Gagal meng-compile aplikasi. Periksa kode Anda."
  exit 1
fi
