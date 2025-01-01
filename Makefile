# Makefile untuk Proyek Anda

# Migrasi database ke atas untuk lingkungan production
migrate-up:
	./scripts/migrate-up.sh

# Migrasi database ke atas untuk lingkungan development
migrate-up-dev:
	./scripts/migrate-up-dev.sh

# Membatalkan migrasi terakhir untuk production
migrate-down:
	./scripts/migrate-down.sh

# Membatalkan migrasi terakhir untuk development
migrate-down-dev:
	./scripts/migrate-down-dev.sh

# Menjalankan aplikasi Go di mode development
run-d:
	./scripts/run.sh

# Menjalankan aplikasi yang sudah di-compile
run:
	./app

# Menghentikan aplikasi yang sedang berjalan
stop:
	./scripts/stop.sh

# Meng-compile aplikasi Go menjadi binary
compile:
	./scripts/compile.sh

# Membantu pengguna melihat daftar perintah yang tersedia
help:
	@echo "Available targets:"
	@echo "  migrate-up        - Apply all up database migrations (production)"
	@echo "  migrate-up-dev    - Apply all up database migrations (development)"
	@echo "  migrate-down      - Roll back the last database migration (production)"
	@echo "  migrate-down-dev  - Roll back the last database migration (development)"
	@echo "  run-d             - Run the Go application in development mode"
	@echo "  run               - Run the compiled application"
	@echo "  stop              - Stop the running application"
	@echo "  compile           - Compile the Go application into a binary"
