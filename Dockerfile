FROM golang:1.14.3 as builder

RUN apt-get update
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the Current Working Directory inside the container
WORKDIR /go/src

COPY go.mod .
COPY go.sum .

# Download all the dependencies
RUN go mod download

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .
RUN go build main.go

EXPOSE 5000

FROM scratch
COPY --from=builder /go/src .
ENTRYPOINT  ["./main"]