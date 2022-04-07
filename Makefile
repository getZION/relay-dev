COVERAGE_DIR ?= .coverage

deploy:
	make build-aws && eb deploy

build-aws:
	GOOS=linux GOARCH=amd64 go build -o bin/application

generate:
	rm -rf gen
	buf generate

mysql-migration:
	cd api/storage/sql/mysql/migrations && go-bindata -pkg migrations .

test:
	@-rm -r $(COVERAGE_DIR)
	@mkdir $(COVERAGE_DIR)
	go test -v -race -covermode atomic -coverprofile $(COVERAGE_DIR)/combined.txt -bench=. -benchmem -timeout 20m ./...

build-docker:
	docker compose up -d --remove-orphans --force-recreate

clean-docker:
	docker compose down -v --remove-orphans

html-coverage:
	go tool cover -html=$(COVERAGE_DIR)/combined.txt
