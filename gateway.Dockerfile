FROM golang:1.16-buster AS build
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -o /gateway ./gateway/gateway.go


##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /gateway /gateway
USER nonroot:nonroot
ENTRYPOINT ["/gateway"]



