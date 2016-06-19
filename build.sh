cd server
$GOROOT/bin/go test -coverprofile=c.out -logtostderr=true
#$GOROOT/bin/go tool cover -html=c.out
$GOROOT/bin/go build

cd ..
