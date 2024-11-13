package depot

//go:generate buf generate

//go:generate mockery --quiet --dir ./api/depot/v1 -r --all --inpackage --case underscore
//go:generate mockery --quiet --dir ./internal -r --all --inpackage --case underscore
