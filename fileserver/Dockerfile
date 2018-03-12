FROM golang as builder
WORKDIR /home/go/
COPY ./fileserver.go .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o fileserver .

FROM busybox
COPY --from=builder /home/go/fileserver /usr/sbin

CMD ["fileserver", "-d", "/usr/apps/web/"]
