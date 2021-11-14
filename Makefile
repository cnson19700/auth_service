gorun:
    go run ./cmd/main.go

dynorestart:
    heroku dyno:restart --app authent-service

logstail:
    heroku logs --tail --app authent-service

gen-protoc:
    @protoc --go_out=plugins=grpc:./proto --proto_path=proto --go_opt=paths=source_relative proto/auth_service.proto