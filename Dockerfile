# builder image
FROM golang:1.20-alpine3.17 as build-env
#STEP 1 build the image
#RUN apk --no-cache add git
ENV GO111MODULE=on
RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .

# Get dependenciesgo
RUN go mod download
# COPY the source code as the last step
COPY . .
RUN GO111MODULE=on GOOS=linux CGO_ENABLED=0 go build /app/main.go

FROM alpine

RUN apk update && apk add ca-certificates
RUN rm -rf /var/cache/apk/*

RUN mkdir /app
WORKDIR /

COPY --from=build-env /app/main /app/
RUN mv /app/main /app/referentielijsten

COPY docs/openapi.yaml /app/docs/openapi.yaml
COPY docs/redoc.html /app/docs/redoc.html

ENV TMPDIR=/tmp
ENV GOMAXPROCS=8

EXPOSE 8000
ENTRYPOINT [ "sh", "-c", "/app/referentielijsten" ]