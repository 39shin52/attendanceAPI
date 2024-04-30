.PHONY: run
run:
	MYSQL_USER=root \
	MYSQL_PASSWORD=root \
	MYSQL_HOST=127.0.0.1 \
	MYSQL_PORT=3306 \
	MYSQL_DATABASE=attendance \
	go run cmd/main.go