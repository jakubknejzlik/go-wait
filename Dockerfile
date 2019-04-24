FROM golang as builder
WORKDIR /go/src/github.com/jakubknejzlik/go-wait-for
COPY . .
RUN go get ./... 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /tmp/wait-for

FROM alpine:3.5

WORKDIR /app

COPY --from=builder /tmp/wait-for /usr/local/bin/wait-for

# RUN apk --update add docker

# https://serverfault.com/questions/772227/chmod-not-working-correctly-in-docker
RUN chmod +x /usr/local/bin/wait-for

ENTRYPOINT []
CMD [ "wait-for", "-h" ]