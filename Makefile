export MYSQL_URL='mysql://root:Docker123!@tcp(localhost:3306)/forumdb?parseTime=true'

migrate-create:
	@ migrate create -ext sql -dir scripts/migrations -seq $(name)

migrate-up:
	@ migrate -database ${MYSQL_URL} -path scripts/migrations up

migrate-down:
	@ migrate -database ${MYSQL_URL} -path scripts/migrations down