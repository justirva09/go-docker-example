ARG BASE_IMAGE=golang:1.19-alpine
FROM ${BASE_IMAGE}

WORKDIR /src

COPY go.* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -a -o main


FROM alpine:3.14

RUN apk add --update --no-cache tzdata ca-certificates

WORKDIR /app

COPY --from=0 /src/main .

EXPOSE 8000

ENTRYPOINT ["./main"]
