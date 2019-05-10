FROM golang:1.11

# Prepare env
ENV GOPATH /go
ENV PATH ${GOPATH}/bin:$PATH
# Install dep (Note: run this locally to speed up image build)
#RUN go get -u github.com/golang/dep/cmd/dep

RUN mkdir -p /go/src/github.com/pscn/flavor2go
WORKDIR /go/src/github.com/pscn/flavor2go
COPY . .
#RUN dep ensure
RUN go build
CMD ["./flavor2go"]
