FROM golang:1.12

# Prepare env
ENV GOPATH /go
ENV PATH ${GOPATH}/bin:$PATH

RUN go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate

# Install dep (Note: run this locally to speed up image build)
#RUN go get -u github.com/golang/dep/cmd/dep

RUN mkdir -p /go/src/github.com/pscn/flavor2go
WORKDIR /go/src/github.com/pscn/flavor2go
COPY . .

RUN chmod +x bin/start.sh

#RUN dep ensure
RUN go build

CMD ["./bin/start.sh"]
