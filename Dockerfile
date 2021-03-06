FROM golang:1.12.0-alpine3.9
RUN mkdir /app
ADD . /app
WORKDIR /app
# COPY openshift/request201.yml .
RUN apk add --no-cache git
RUN go build -o main .
CMD ["/app/main"]