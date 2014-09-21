FROM google/golang

WORKDIR /gopath/src/github.com/erlendr/uploader
ADD . /gopath/src/github.com/erlendr/uploader/

#temp dir for uploads
RUN mkdir /gopath/src/github.com/erlendr/uploader/temp
RUN chmod 777 /gopath/src/github.com/erlendr/uploader/temp

# get dependencies
RUN go get github.com/erlendr/store
RUN go get github.com/go-martini/martini
RUN go get github.com/erlendr/uploader

CMD []
ENTRYPOINT ["/gopath/bin/uploader"]