docker build -f Dockerfile -t gorsvp .
docker container run --rm -p 8080:8080 --name gorsvp gorsvp
