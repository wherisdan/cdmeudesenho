FROM node:alpine AS build-frontend

WORKDIR /
COPY ./frontend/package*.json .

RUN npm install

COPY ./frontend .

RUN npm run build

FROM golang:alpine AS build-server

WORKDIR /
COPY server .

RUN go get
RUN go build -o /build/server

FROM alpine:latest AS development

WORKDIR /
COPY --from=build-server /build/server .

ENV BUILD_MODE=development

CMD ["./server"]

EXPOSE 80

FROM alpine:latest AS production

WORKDIR /
COPY --from=build-server /build/server .
COPY --from=build-frontend /dist .

ENV BUILD_MODE=production

CMD ["./server"]
EXPOSE 80