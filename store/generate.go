package store

//go:generate buf generate

//go:generate mockery --quiet --dir ./api/store/v1 -r --all --inpackage --case underscore
//go:generate mockery --quiet --dir ./internal -r --all --inpackage --case underscore

//go:generate swagger generate client -q -f ./docs/api.swagger.json -c storesclient -m storesclient/models --with-flatten=remove-unused
