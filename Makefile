deploy:
	cd ui && yarn export && cd .. && make build-aws && eb deploy

build-aws:
	GOOS=linux GOARCH=amd64 go build -o bin/application

generate:
	rm -rf gen
	buf generate

test:
	cd did && go test -v
	cd api && ginkgo -v

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
