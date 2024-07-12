FROM golang:1.22

COPY . /happiness
WORKDIR /happiness/src

RUN go build -o /bin/main .

EXPOSE 5000

CMD ["/bin/main"]