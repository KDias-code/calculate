FROM golang:1.21.0
WORKDIR /source
COPY . /source
RUN go build -o server ./main.go
EXPOSE 50051
CMD ["./server"]