FROM golang
ADD . /go
WORKDIR /go
RUN go install github.com/tools/godep@latest
RUN godep save
RUN godep restore
#RUN go install github.com/telkomdev/indihome
ENTRYPOINT /go/bin/indihome
EXPOSE 80

#FROM golang
#ADD . /go/src/github.com/telkomdev/indihome/backend
#WORKDIR /go/src/github.com/telkomdev/indihome
#RUN go get github.com/tools/godep
#RUN godep restore
#RUN go install github.com/telkomdev/indihome
#ENTRYPOINT /go/bin/indihome
#LISTEN 80
