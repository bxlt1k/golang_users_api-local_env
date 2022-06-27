start:
	cd docker && docker-compose up -d

down:
	cd docker && docker-compose down

build:
	make down && cd docker && docker-compose up --build -d

restart: down start

migrate_create:
	./docker/utils/migrate create -ext sql -dir ./services/mysql/migrations -seq ${name}

migrate_up_all:
	./docker/utils/migrate -path ./services/mysql/migrations -database "mysql://root:root@tcp(first_mysql:3306)/first" up

migrate_down_all:
	./docker/utils/migrate -path ./services/mysql/migrations -database "mysql://root:root@tcp(localhost:3306)/first" down