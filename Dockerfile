FROM golang:1.23.3-alpine

WORKDIR /app

RUN apk add --no-cache curl

COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o pb .

EXPOSE 8090
ENTRYPOINT ["/bin/sh", "-c", "./pb superuser upsert \"$PB_ADMIN_EMAIL\" \"$PB_ADMIN_PASSWORD\" && exec ./pb serve --http=0.0.0.0:8090"]