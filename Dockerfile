FROM node:alpine AS build-stage

WORKDIR /
COPY ./frontend/package*.json .

RUN npm install

COPY ./frontend .

RUN npm run build

FROM golang:alpine

WORKDIR /
COPY server .

COPY --from=build-stage /dist ./dist

RUN go get
RUN go build -o /build/server

CMD [ "/build/server" ]

EXPOSE 8080