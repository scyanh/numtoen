FROM golang:1.14.3

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN make build
EXPOSE 8080
CMD ["./numtoen"]