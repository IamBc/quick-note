cd server
$GOROOT/bin/go test -coverprofile=c.out -logtostderr=true
$GOROOT/bin/go build

cd ..
