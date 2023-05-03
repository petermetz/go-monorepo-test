echo "\n### Running A test ###\n"
cd pkg/a
make test
cd -
echo "\n\n### Running B test ###\n"
cd pkg/b
make test
cd -
echo "\n\n### Running C ###\n"
cd pkg/c
make build
./c
cd -
