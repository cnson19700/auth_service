gorun:
	go run ./cmd/main.go

dynorestart:
    heroku dyno:restart --app authent-service
    
logstail:
    heroku logs --tail --app authent-service