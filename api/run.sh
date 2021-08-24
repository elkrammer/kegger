docker build -f Dockerfile -t kegger-api .
docker container run --rm -p 8080:8080 --name kegger-api --env-file ./.env kegger-api
