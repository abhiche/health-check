# pulling a lightweight version of golang
FROM golang:1.11-alpine
RUN apk --update add --no-cache git

ENV GOPATH /go
# Copy the local package files to the container's workspace.

ADD . /go/src/health-check
WORKDIR /go/src/health-check

# copy crontabs for root user
RUN echo "$WORKDIR"
COPY $WORKDIR/internal/prober/cronjobs /etc/crontabs/root

RUN git config --global url."git://".insteadOf https://

RUN go get github.com/abhiche/health-check/pkg/waitgroup

RUN go build ./internal/prober

RUN chmod +x ./prober

# start crond with log level 8 in foreground, output to stderr
CMD ["crond", "-f", "-d", "8"]