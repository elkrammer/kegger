FROM node:14.0.0-alpine as base

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

ENV NODE_ENV=production

RUN npm run build

# production stage
FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
