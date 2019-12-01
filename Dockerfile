FROM  golang:alpine as BUILD
RUN go version
RUN apk add --update make
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
WORKDIR /src
COPY . /src
RUN go build -o ./bin/slackjob

FROM alpine:latest
WORKDIR /src
COPY --from=BUILD /src/bin /src/bin
RUN ls /src/bin
RUN chmod 777 /src/bin/slackjob
#ENV PATH = $PATH:/helm/bin
ENTRYPOINT   ["/src/bin/slackjob"]