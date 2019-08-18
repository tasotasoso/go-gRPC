FROM golang:1.12
ENV PKG_PATH /go/src/
ENV GO111MODULE on
WORKDIR $PKG_PATH
RUN go mod download
COPY ./ $PKG_PATH
RUN ls
RUN go install .

FROM grpc/go
COPY --from=0 /go/bin /go/bin/server 
ENV PORT 50051
RUN ls /go/bin/server
CMD ["/go/bin/server/main"]
