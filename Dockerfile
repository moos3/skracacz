FROM golang:alpine

RUN apk --update add git curl && rm -rf /var/cache/apk/*
ENV APP_DIR $GOPATH/src/github.com/moos3/skracacz

RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 && chmod +x /usr/local/bin/dep

RUN mkdir /app
ADD . $APP_DIR
WORKDIR $APP_DIR
RUN /usr/local/bin/dep ensure -vendor-only
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/shorty .

EXPOSE 3000

CMD ["/app/shorty"]