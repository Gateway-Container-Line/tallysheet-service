FROM golang:alpine

#RUN apk update && apk add --no-cache git

# additional private repo
RUN apk update && apk add --no-cache git ca-certificates
#RUN mkdir -p $GOROOT/src/github.com/Gateway-Container-Line
#RUN git clone -b go https://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:x-oauth-basic@github.com/Gateway-Container-Line/tallysheet-service.git $GOPATH/src/github.com/Gateway-Container-Line/tallysheet-service
##ENV PATH /go/bin:$PATH

#WORKDIR /go/src/github.com/Gateway-Container-Line/tallysheet-service
#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
#RUN mkdir ~/.ssh
#RUN touch ~/.ssh/known_hosts
#RUN ssh-keyscan -t rsa github.com >> ~/.ssh/known_hosts
#
## allow private repo pull
#RUN git config --global url."https://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:x-oauth-basic@github.com/".insteadOf "https://github.com/"

#RUN apk update && apk add --no-cache ca-certificates git-core ssh
#
#RUN echo "[url \"git@github.com:\"]\n\tinsteadOf = https://github.com/" >> /root/.gitconfig
#RUN mkdir /root/.ssh && echo "StrictHostKeyChecking no " > /root/.ssh/config
#ADD keys/my_key_rsa /root/.ssh/id_rsa
#RUN chmod 700 /root/.ssh/id_rsa
#RUN echo "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config
#RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/
# until here



# additional private repo
ARG GITHUB_PAT
#ARG GITHUB_TOKEN
#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
ENV CGO_ENABLED=0 GO111MODULE=on GOOS=linux TOKEN=$GITHUB_PAT
#
RUN go env -w GOPRIVATE=github.com/Gateway-Container-Line/*
#RUN export GOPRIVATE="github.com/Gateway-Container-Line"
#
RUN git config --global url."https://${GITHUB_PAT}:x-oauth-basic@github.com/".insteadOf "https://github.com/"
#RUN git config --global  url."ssh://git@github.com/Gateway-Container-Line".insteadOf "https://github.com/Gateway-Container-Line"
#RUN git config --global  \
#                          url."ssh://git@github.com/Gateway-Container-Line".insteadOf \
#                          "https://github.com/Gateway-Container-Line"
#until here

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -ldflags '-w -s' -o tallysheet-service
#RUN go build -o tallysheet-service

RUN export GO111MODULE=on

ENTRYPOINT ["/app/tallysheet-service"]

CMD ["go", "run" , "main.go"]