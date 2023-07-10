FROM golang:alpine AS build

WORKDIR /usr/app

COPY . .

EXPOSE 3333

CMD [ "make", "dev" ]
