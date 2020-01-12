![Kegger](./frontend/src/assets/logo.png)

Kegger is an Event / RSVP manager to keep track of guests attendance.
The application is split into the API(Go) and the frontend(Vue.js).

## Features

- Create parties with guests
- Send invite by email (only SendGrid supported as of now)
- Create new / customize invitation email template
- Create multiple administrators (hosts)
- Keep track of
   - Guests Attendance
   - Invitations sent / not sent
   - Invitation was opened by guest

## Backend / API

Kegger's API is built using the [Echo](https://echo.labstack.com/) Go web framework.

### Requirements

- Go 1.12+
- Postgres 10+
- SendGrid account for being able to send out the email invites (optional)

### Build

```sh
cd api
edit .env file as per instructions in api/README.md)
go build
./gorsvp
```

## Frontend

Kegger's Frontend is built with Vue.js.

### Requirements

- Node.js 6+

### Build

```sh
cd frontend
npm install # to install dependencies
edit .env file as per instructions in frontend/README.md)
npm run serve
```
