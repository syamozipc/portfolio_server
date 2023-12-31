.PHONY: migrate-new
migrate-new:
ifdef name
	sql-migrate new $(name)
else
	@echo "nameパラメータを指定してください\n例:\"make migrate-new name=create_user_table\""
endif

.PHONY: run
run:
	go run main.go serve

.PHONY: migrate
migrate:
	go run main.go migrate

# EC2で使用しているAmazon Linuxで動かせるよう環境を指定
.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build main.go