generate:
	rm -rf gen
	buf generate

test:
	cd did && go test -v
