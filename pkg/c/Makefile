run-vendor:
	go mod edit -replace github.com/petermetz/go-monorepo-test/pkg/a=../a
	go mod edit -replace github.com/petermetz/go-monorepo-test/pkg/b=../b
	go mod vendor

undo-vendor:
	rm -rf vendor
	go mod edit -dropreplace github.com/petermetz/go-monorepo-test/pkg/a
	go mod edit -dropreplace github.com/petermetz/go-monorepo-test/pkg/b
	
build:
	go build -v .
	
build-local: run-vendor build undo-vendor

clean:
	rm -rf vendor