FROM golang:1.22.5-alpine

WORKDIR /src

COPY . .

WORKDIR "/src/server app"

RUN go build -o ./build/server

CMD [ "./build/server" ]