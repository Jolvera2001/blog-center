# GO image
FROM golang:1.23 AS builder

# set current working directory
WORKDIR /app

# copy go.mod and go.sum files first
COPY go.mod go.sum ./
RUN go mod download

# copy entire cmd directory to working directory
COPY cmd ./cmd/
COPY internal/ ./internal/

# Build the go application
RUN go build -o main ./cmd/blog-center/main.go

# use a smaller image for final exe
FROM alpine:latest

# set working directory in final image
WORKDIR /root/

# copy compiled binary
COPY --from=builder /app/main .

# command to run exe
CMD [ "./main" ]
