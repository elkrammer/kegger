# gorsvp

Go RSVP Tracking System.

It uses the [Echo](https://echo.labstack.com/) Go web framework.

This app is only the API Server (backend). The frontend app is [WingDing](https://github.com/elkrammer/wingding).

# Requirements
* Go 1.12+
* Postgres 10+
* (optional) SendGrid account for being able to send out the email invites

# Sample .env File
```
DB_USERNAME="dbuser"
DB_PASSWORD="dbpass"
DB_NAME="gorsvp"
DB_HOST="localhost"
DB_PORT="5432"
JWT_SECRET="5UP3R S3CR3T SH1T!"
SENDGRID_API_KEY="KEY"
```
