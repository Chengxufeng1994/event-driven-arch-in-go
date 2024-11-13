package basket

//go:generate buf generate

//go:generate mockery --quiet --dir ./api/basket/v1 -r --all --inpackage --case underscore
//go:generate mockery --quiet --dir ./internal -r --all --inpackage --case underscore

//go:generate swagger generate client -q -f ./docs/api.swagger.json -c basketsclient -m basketsclient/models --with-flatten=remove-unused
