FROM golang:1.11 as builder

ARG APP_NAME=vista-demo
ARG LOG_DIR=/${APP_NAME}/logs

RUN mkdir -p ${LOG_DIR}

ENV LOG_FILE_LOCATION=${LOG_DIR}/app.log 

WORKDIR $GOPATH/src/gitlab.com/YottaDB/Demo/vista-kubernetes/
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/vista-kubernetes .

## Build second image
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /go/bin/vista-kubernetes .
COPY ./static/ ./static/


EXPOSE 8080

CMD ["./vista-kubernetes"]
