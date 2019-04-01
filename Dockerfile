FROM golang

WORKDIR $GOPATH/src/github.com/rof20004/vuttr
COPY . .
RUN export GO111MODULE=on && go build .

CMD [ "./vuttr" ]