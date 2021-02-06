
function genAll {
  env CGO_ENABLED=0 GO111MODULE=on go get github.com/vektra/mockery/v2/
  env CGO_ENABLED=0 GO111MODULE=on go install github.com/vektra/mockery/v2/cmd
  $GOPATH/bin/mockery --name=Client --output=. --filename=client_mock.go --inpackage --dir=http
}

# __main__
genAll
