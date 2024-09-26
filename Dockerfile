# GO image
FROM golang:1.23-alpine AS builder

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

# checking if exe is in container
RUN ls -l /app/

# use a smaller image for final exe
FROM alpine:latest

# set working directory in final image
WORKDIR /root/

# copy compiled binary
COPY --from=builder /app/main .

# Conditional copy based on build argument
ARG COPY_ENV=false
RUN if [ "$COPY_ENV" = "true" ]; then cp /app/.env .; fi

# make it executable
RUN chmod +x ./main

# command to run exe
CMD ["./main"]
