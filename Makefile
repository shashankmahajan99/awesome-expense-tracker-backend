DB_URL=mysql://root:password@tcp(localhost:3306)/awesome_expense_tracker?multiStatements=true

generate_from_protos:   
	rm -rf generated
	mkdir generated
	cd api; buf generate; cd ..
	make convert_swagger_v2_to_v3
	cp -r generated/* api/

convert_swagger_v2_to_v3:
	for file in generated/*swagger.json; do \
		base=$$(basename $$file .swagger.json); \
		if [[ $$base != *_v3 ]]; then \
			echo "Converting $$file"; \
			swagger2openapi $$file -o "generated/$${base}_v3.swagger.json"; \
		fi \
	done

local-server-run:
	go run server/main.go

podman-server-build:
	podman build -t awesome-expense-tracker-backend -f Dockerfile .

podman-server-run:
	podman run -d -p 8080:8080 -p 8081:8081 -e PORT=8080 -e GRPC_PORT=8081 -e MYSQLUSER=root -e MYSQLPASSWORD=password -e MYSQLADDR=0.0.0.0:3306 -e MYSQLDBNAME=awesome_expense_tracker --name awesome-expense-tracker-backend --rm awesome-expense-tracker-backend

mysql-build:
	podman build -t mysql -f Dockerfile.mysql .

check_mysql_server_running:
	@if podman inspect -f '{{.State.Running}}' mysql-server >/dev/null 2>&1; then \
		echo "Stopping existing mysql-server container..."; \
		podman stop mysql-server; \
	fi

mysql-run: check_mysql_server_running
	podman run -d -p 3306:3306 --name mysql-server --rm -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=awesome_expense_tracker mysql

mysql-client:
	podman exec -it mysql-server mysql -h localhost -u root -ppassword awesome_expense_tracker

migrateup:
	migrate -path pkg/db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path pkg/db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path pkg/db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path pkg/db/migration -database "$(DB_URL)" -verbose down 1

sqlc-install:
	podman pull sqlc/sqlc

sqlc-generate:
	podman run --rm -v $(shell pwd):/src -w /src sqlc/sqlc generate
	
.PHONY: check_mysql_server_running run_mysql_server mysql-client migrateup migrateup1 migratedown migratedown1 sqlc-install sqlc-generate
