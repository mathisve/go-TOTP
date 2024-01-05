# go-TOTP
totp (Timed One Time Password) experiment in golang

### server
- GetSecret - returns a SecretID and Secret
- Challenge - Takes SecretID and TOTP, returns OK if TOTP is correct

### client
1. Requests SecretID and Secret
2. Computes TOTP based on the Secret and Timeblock
3. Sends Challenge request to server with SecretID and TOTP


## setup
protoc command:
```shell
protoc --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    totp/totp.proto
```

start:
```shell
go run server/server.go
go run client/client.go
```