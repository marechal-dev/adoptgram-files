FROM golang:alpine AS build

WORKDIR /usr/app

COPY . .

RUN make build

EXPOSE 3333

CMD [ "./main" ]
