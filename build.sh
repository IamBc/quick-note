cd server

$GOROOT/bin/go test -coverprofile=c.out -logtostderr=true
if [ "$1" == "report" ]; then
    $GOROOT/bin/go tool cover -html=c.out
fi
$GOROOT/bin/go build

cd ..
