run-vendor:
	go mod edit -replace github.com/petermetz/go-monorepo-test/pkg/a/v2=../a
	go mod vendor

undo-vendor:
	rm -rf vendor
	go mod edit -dropreplace github.com/petermetz/go-monorepo-test/pkg/a/v2
	
test:
	go test -v .
	
test-local: run-vendor test undo-vendor

clean:
	rm -rf vendor