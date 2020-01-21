![Kegger](./frontend/src/assets/logo.png)

# Kegger - RSVP Manager

Kegger is an Event / RSVP manager to keep track of guests attendance.
The application is split in two: the API (Go) and the frontend(Vue.js).

## Features

- Create parties with guests
- Send invite by email (only SendGrid supported as of now)
- Create new / customize invitation email template
- Create multiple administrators (hosts)
- Keep track of
   - Guests Attendance
   - Invitations sent / not sent
   - Invitation was opened by guest

## Running the project using Docker

You can quickly get started using the project with Docker:

```sh
# edit .env files for in api/ and frontend/
docker-compose up
```

Then open your browser and head to http://127.0.0.1:8081/
The default email to login is admin@admin.com and the password is admin.

# Building the project locally

## Backend / API

Kegger's API is built using the [Echo](https://echo.labstack.com/) Go web framework.

### Requirements

- Go 1.12+
- Go Echo v4.1.5 (later versions are [broken](https://github.com/labstack/echo/issues/1356) in the way they handle [bingings & params](https://github.com/labstack/echo/issues/1466)
- Postgres 10+
- SendGrid account for being able to send out the email invites (optional)

### Build

```sh
cd api
# edit file api/.env
go build
./gorsvp
```

## Frontend

Kegger's Frontend is built with [Vue.js](https://vuejs.org/).

### Requirements

- Node.js 6+
- Vue 2.6
- Vuex 3.1

### Build

```sh
cd frontend
npm install
# edit file frontend/.env
npm run serve
```

Then open your browser and head to http://127.0.0.1:8080/
The default email to login is admin@admin.com and the password is admin.
