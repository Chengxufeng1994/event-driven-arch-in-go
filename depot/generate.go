package depot

//go:generate buf generate

//go:generate mockery --quiet --dir ./api/depot/v1 -r --all --inpackage --case underscore
//go:generate mockery --quiet --dir ./internal -r --all --inpackage --case underscore

//go:generate swagger generate client -q -f ./docs/api.swagger.json -c depotclient -m depotclient/models --with-flatten=remove-unused
