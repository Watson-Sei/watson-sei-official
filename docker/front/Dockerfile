FROM node:15.0.1-alpine3.10

ENV LANG=C.UTF-8 TZ=Asia/Tokyo

WORKDIR /app

RUN apk update && \
    apk upgrade && \
    apk add --no-cache make gcc g++ python

COPY ./docker/front/package*.json ./

RUN yarn install

COPY ./docker/front .

ENV HOST 0.0.0.0
EXPOSE 3000

RUN yarn build