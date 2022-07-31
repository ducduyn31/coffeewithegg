##
## Build
##
FROM golang:1.18.4-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY apps/adam apps/adam

WORKDIR /app/apps/adam

RUN go build -o /adam-server -modfile /app/go.mod

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /adam-server /adam

EXPOSE 4001

USER nonroot:nonroot

ENTRYPOINT ["/adam"]
