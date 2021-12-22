FROM golang:1.16-buster AS build
WORKDIR /app


# prevent the re-installation of vendors at every change in the source code
COPY ./go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

RUN go build -o /service ./main.go

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /


COPY --from=build /service /service

USER nonroot:nonroot
ENTRYPOINT ["/service"]

