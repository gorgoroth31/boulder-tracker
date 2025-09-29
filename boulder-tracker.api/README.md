## build the project for ubuntu server

```powershell
$env:GOOS="linux"
$env:GOARCH="amd64"
go build -o bouldertracker.api
```
