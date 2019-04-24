FROM golang as builder
WORKDIR /go/src/github.com/jakubknejzlik/go-wait
COPY . .
RUN go get ./... 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /tmp/wait

FROM alpine:3.5

WORKDIR /app

COPY --from=builder /tmp/wait /usr/local/bin/wait

# RUN apk --update add docker

# https://serverfault.com/questions/772227/chmod-not-working-correctly-in-docker
RUN chmod +x /usr/local/bin/wait

ENTRYPOINT []
CMD [ "wait", "-h" ]