FROM node:10 as node-build

WORKDIR /app/clientsrc

COPY clientsrc/package.json clientsrc/package-lock.json /app/clientsrc/

RUN npm ci

COPY clientsrc /app/clientsrc

RUN npm run build

FROM golang:alpine as go-build

RUN apk add --no-cache git

WORKDIR /go/src/app
COPY . /go/src/app

RUN rm -r /go/src/app/clientsrc \
    && go get -d -v \
    && go install -v \
    && go build -o app

FROM scratch as scratch

ENV CLIENT_PATH=/go/src/app/client

COPY --from=node-build /app/client /go/src/app/client
COPY --from=go-build /go/bin/app /

ENTRYPOINT ["/app"]

# FROM gcr.io/distroless/base as distroless
FROM alpine as distroless

ENV CLIENT_PATH=/go/src/app/client

RUN apk add --no-cache bash

COPY --from=node-build /app/client /go/src/app/
COPY --from=go-build /go/bin/app /

ENTRYPOINT ["/app"]
