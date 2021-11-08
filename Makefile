gorun:
	go run ./cmd/main.go

dynorestart:
    heroku dyno:restart --app authenication-service
    heroku dyno:restart --app users-servicee
    heroku dyno:restart --app community-servicee
    heroku dyno:restart --app movie-servicee