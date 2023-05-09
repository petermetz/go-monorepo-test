run-vendor:
	go mod edit -replace github.com/petermetz/go-monorepo-test/pkg/a/v2=../a
	go mod edit -replace github.com/petermetz/go-monorepo-test/pkg/b/v2=../b
	go mod vendor

undo-vendor:
	rm -rf vendor
	go mod edit -dropreplace github.com/petermetz/go-monorepo-test/pkg/a/v2
	go mod edit -dropreplace github.com/petermetz/go-monorepo-test/pkg/b/v2
	
build:
	go build -v .
	
build-local: run-vendor build undo-vendor

clean:
	rm -rf vendor