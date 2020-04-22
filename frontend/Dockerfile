FROM node:13.13.0-alpine as base

ENV NODE_ENV=development

WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

# install and cache app dependencies
COPY package.json /app/package.json
RUN npm install
RUN npm install -g @vue/cli@4.3.1

# install frontend
COPY . .

# build stage
FROM base as build-stage

ARG VUE_APP_API_SERVER
ARG VUE_APP_EMAIL_FROM_NAME
ARG VUE_APP_EMAIL_FROM
ARG VUE_APP_EMAIL_SUBJECT
ARG VUE_APP_EMAIL_TO
ARG VUE_APP_EMAIL_TO_NAME

ENV VUE_APP_API_SERVER=$VUE_APP_API_SERVER \
    VUE_APP_EMAIL_FROM_NAME=$VUE_APP_EMAIL_FROM_NAME \
    VUE_APP_EMAIL_FROM=$VUE_APP_EMAIL_FROM \
    VUE_APP_EMAIL_SUBJECT=$VUE_APP_EMAIL_SUBJECT \
    VUE_APP_EMAIL_TO=$VUE_APP_EMAIL_TO \
    VUE_APP_EMAIL_TO_NAME=$VUE_APP_EMAIL_TO_NAME \
    NODE_ENV=production

RUN npm run build

# production stage
FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

# start app
CMD ["nginx", "-g", "daemon off;"]
