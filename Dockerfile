FROM golang:1.12

# Prepare env
ENV GOPATH /go
ENV PATH ${GOPATH}/bin:$PATH

# Install DB migration tool
RUN go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate

# Install dependency tool
RUN go get -u github.com/golang/dep/cmd/dep

# Copy src
RUN mkdir -p /go/src/github.com/gusta-project/go-api
WORKDIR /go/src/github.com/gusta-project/go-api
COPY . .

# avoid problems with wrong line endings that may have been introduced with
# git autocrlf=true
RUN tr -d '\r' < bin/start.sh > bin/tmp.sh
RUN mv bin/tmp.sh bin/start.sh
RUN chmod +x bin/start.sh

RUN dep ensure
RUN go build

CMD ["./bin/start.sh"]
