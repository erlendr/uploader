FROM google/golang

WORKDIR /gopath/src/github.com/erlendr/uploader
ADD . /gopath/src/github.com/erlendr/uploader/

#temp dir for uploads
RUN mkdir -p /gopath/src/github.com/erlendr/uploader/temp
RUN chmod 777 /gopath/src/github.com/erlendr/uploader/temp

# get dependencies
RUN go get github.com/erlendr/uploader/

EXPOSE 3000
CMD []
ENTRYPOINT ["/gopath/bin/uploader"]