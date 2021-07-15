FROM golang:1.14.3

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/scyanh/numtoen

COPY go.mod .
COPY go.sum .

# Download all the dependencies
RUN go mod download

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

RUN go build main.go

RUN go build -o numtoen main.go

EXPOSE 8080
CMD ["./numtoen"]