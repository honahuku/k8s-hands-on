FROM alpine:3.19

COPY ./test.sh /test.sh
RUN apk --update --no-cache add bash=5.2.21-r0 curl=8.5.0-r0 stress-ng=0.17.01-r0

CMD ["/test.sh"]
