# pulling a lightweight version of golang
FROM golang:1.11-alpine
RUN apk --update add --no-cache git

ENV GOPATH /go
# Copy the local package files to the container's workspace.
ADD . /go/src/health-check/prober
WORKDIR /go/src/health-check/prober

RUN git config --global url."git://".insteadOf https://

RUN go build

RUN chmod +x ./prober

For those who wants to use a simple and lightweight image:

FROM alpine:3.6

# copy crontabs for root user
COPY config/cronjobs /etc/crontabs/root

# start crond with log level 8 in foreground, output to stderr
CMD ["crond", "-f", "-d", "8"]


Where cronjobs is the file that contains your cronjobs, in this form:


# Run the command by default when the container starts.
ENTRYPOINT ["./prober"]