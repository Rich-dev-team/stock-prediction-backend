# stock-prediction-backend

## Start steps

> `$ go mod download `

> `$ go get -u github.com/gin-gonic/gin`

> `$ go run main.go`

## Remove unused package and install required package

> `$ go mod tidy` 

## Debug postgres lock
> https://wiki.postgresql.org/wiki/Lock_Monitoring

## Add go env to path
> `export PATH=$PATH:$HOME/go/bin`

## trouble shooting in mockgen
> if you cannot find go/src/xxx, then add `-build_flags=--mod=mod` to command

## show go testcase coverage settings
> https://dev.to/vuong/golang-in-vscode-show-code-coverage-of-after-saving-test-8g0