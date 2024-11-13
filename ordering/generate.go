package ordering

//go:generate buf generate

//go:generate mockery --quiet --dir ./api/order/v1 -r --all --inpackage --case underscore
//go:generate mockery --quiet --dir ./internal -r --all --inpackage --case underscore

//go:generate swagger generate client -q -f ./docs/api.swagger.json -c orderingclient -m orderingclient/models --with-flatten=remove-unused
