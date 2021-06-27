FROM golang:1.16.5-alpine AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


# Production image to run our app
FROM alpine:latest AS production
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache git make musl-dev go
COPY --from=builder /app/main .


# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH


RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin
EXPOSE 3000
CMD ["./main"]


# To run the container Do

# docker build -t differ .
# docker run -it -p 3000:3000
