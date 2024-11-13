package customer

//go:generate buf generate

//go:generate mockery --quiet --dir ./api/customer/v1 -r --all --inpackage --case underscore
//go:generate mockery --quiet --dir ./internal -r --all --inpackage --case underscore
////go:generate mockery --quiet --dir ./internal -r --all --case underscore --output ./mocks

//go:generate swagger generate client -q -f ./docs/api.swagger.json -c customersclient -m customersclient/models --with-flatten=remove-unused
