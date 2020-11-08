FROM golang:1.14.3

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/scyanh/numtoen

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

EXPOSE 8080
CMD ["./numtoen"]