generate:
	rm -rf gen
	buf generate

test:
	cd did && go test -v

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
