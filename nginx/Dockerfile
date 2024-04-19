FROM node:20.10.0 AS builder

WORKDIR /app

COPY client/package.json package.json
COPY client/yarn.lock yarn.lock

RUN yarn

COPY client .

RUN yarn build

FROM nginx

COPY --from=builder /app/dist /etc/nginx/html/

COPY ./nginx/default.conf /etc/nginx/conf.d/default.conf
