FROM node:15.0.1-alpine3.10

ENV LANG=C.UTF-8 TZ=Asia/Tokyo

WORKDIR /front

RUN apk update

COPY ./package.json ./

RUN yarn install

COPY ./ .

RUN yarn run build

EXPOSE 3000
