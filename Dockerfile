FROM golang:1.13 as builder
#RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
ENV USER=appuser
ENV UID=10001
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"
RUN mkdir /app
ADD . /app/
WORKDIR /app
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o main .

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /app/main /app/main
WORKDIR /app
USER appuser:appuser
EXPOSE 8080/tcp
CMD ["/app/main"]
