FROM golang:alpine
COPY . .
RUN go mod download
CMD ["go", "run", "cmd/main.go"]
