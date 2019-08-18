FROM golang:1.12
ENV PKG_PATH /go/src/
ENV GO111MODULE on
WORKDIR $PKG_PATH
RUN go mod download
COPY ./ $PKG_PATH
RUN ls
RUN go install .

FROM grpc/go
COPY --from=0 /go/bin /go/bin/client 
COPY --from=0 /go/src/postMenu.html /go/bin/client 
EXPOSE 50050
RUN ls /go/bin/client
ENTRYPOINT ["/go/bin/client/main"]