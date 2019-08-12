mkdir gen

swagger generate server -f swagger/swagger.yml -t gen --exclude-main -P models.Principal