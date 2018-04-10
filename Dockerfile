FROM scratch

COPY conf /conf
COPY static /static
COPY go-admin /go-admin
COPY version /version
COPY views /views

#ENV GOPATH /
WORKDIR  /

EXPOSE 8080

ENTRYPOINT ["/go-admin"]