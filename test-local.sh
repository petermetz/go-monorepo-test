echo "\n### Running A test ###\n"
cd pkg/a
make test
cd -
echo "\n\n### Running B test (local) ###\n"
cd pkg/b
make test-local
cd -
echo "\n\n### Running C (local) ###\n"
cd pkg/c
make build-local
./c
cd -
