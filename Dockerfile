FROM node:10 as node-build

WORKDIR /app/clientsrc

COPY clientsrc/package.json clientsrc/package-lock.json /app/clientsrc/

RUN npm ci

COPY clientsrc /app/clientsrc

RUN npm run build

FROM golang as go-build

WORKDIR /go/src/app
COPY . /go/src/app

RUN rm -r /go/src/app/clientsrc \
    && go get -d -v \
    && go install -v \
    && go build

FROM gcr.io/distroless/base

ENV CLIENT_PATH=/go/src/app/client

COPY --from=node-build /app/client /go/src/app/client
COPY --from=go-build /go/bin/app /

CMD ["/app"]
