RD /s /Q %cd%\server

echo "Create server"
MD server

echo "Start..."

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=mipsle
SET GOMIPS=softfloat


echo now the CGO_ENABLED:
 go env CGO_ENABLED

echo now the GOOS:
 go env GOOS

echo now the GOARCH:
 go env GOARCH
go build -ldflags="-w -s" main.go

MOVE main %cd%\server


SET CGO_ENABLED=1
SET GOOS=windows
SET GOARCH=amd64


echo now the CGO_ENABLED:
 go env CGO_ENABLED

echo now the GOOS:
 go env GOOS

echo now the GOARCH:
 go env GOARCH

echo "Finished":