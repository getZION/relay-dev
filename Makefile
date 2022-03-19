COVERAGE_DIR ?= .coverage

deploy:
	cd ui && yarn export && cd .. && make build-aws && eb deploy

build-aws:
	GOOS=linux GOARCH=amd64 go build -o bin/application

defs:
	make generate && make compile-ts && make compile-docs

generate:
	rm -rf gen
	buf generate

test:
	@-rm -r $(COVERAGE_DIR)
	@mkdir $(COVERAGE_DIR)
	go test -v -race -covermode atomic -coverprofile $(COVERAGE_DIR)/combined.txt -bench=. -benchmem -timeout 20m ./...

compile-ts:
	rm -rf ui/proto
	protoc proto/identityhub/v1/*.proto \
		--plugin=./ui/node_modules/.bin/protoc-gen-ts \
		--js_out="import_style=commonjs,binary:./ui" \
		--ts_out=service=grpc-web:./ui
	protoc proto/zion/v1/*.proto \
		--plugin=./ui/node_modules/.bin/protoc-gen-ts \
		--js_out="import_style=commonjs,binary:./ui" \
		--ts_out=service=grpc-web:./ui
	protoc proto/protoc-gen-gorm/**/*.proto \
		--plugin=./ui/node_modules/.bin/protoc-gen-ts \
		--js_out="import_style=commonjs,binary:./ui" \
		--ts_out=service=grpc-web:./ui

compile-docs:
	protoc proto/zion/v1/*.proto \
		--doc_out=./docs \
		--doc_opt=markdown,grpc-zion.md

	protoc proto/identityhub/v1/*.proto \
		--doc_out=./docs \
		--doc_opt=markdown,grpc-identityhub.md

build-docker:
	docker compose up -d --remove-orphans --force-recreate

clean-docker:
	docker compose down -v --remove-orphans

html-coverage:
	go tool cover -html=$(COVERAGE_DIR)/combined.txt
