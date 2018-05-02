# Full SDK version
FROM golang:1.10-alpine AS build

RUN apk update && apk upgrade \
    && apk add curl git

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go/src/github.com/sasaken555/ponz_goecho_server
COPY . .

RUN dep ensure
RUN go build -o ponz_goecho_server

# Final Output
FROM golang:1.10-alpine
COPY --from=build /go/src/github.com/sasaken555/ponz_goecho_server/ponz_goecho_server /bin/ponz_goecho_server
CMD /bin/ponz_goecho_server
