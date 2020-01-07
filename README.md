# gorsvp

Go RSVP Tracking System.

It uses the [Echo](https://echo.labstack.com/) Go web framework.

This app is only the API Server (backend). The frontend app is [WingDing](https://github.com/elkrammer/wingding).

# Requirements
* Go 1.12+
* Postgres 10+
* (optional) SendGrid account for being able to send out the email invites

# Running gorsvp
1. Edit .env file
2. `go build`
3. `./gorsvp`

# Sample .env File
```
DB_USERNAME="dbuser"
DB_PASSWORD="dbpass"
DB_NAME="gorsvp"
DB_HOST="localhost"
DB_PORT="5432"
JWT_SECRET="5UP3R S3CR3T SH1T!" # secret string for signing the jwt token
SENDGRID_API_KEY="KEY" # used for sending email invites
RUN_MIGRATIONS="true" # run database migrations on init
```
