FROM golang:1.15.0-alpine AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest AS production
COPY --from=builder /app .
CMD ["./main"]




# To run the container Do
# docker build -t differ .
# docker run -it -p 3000:3000
