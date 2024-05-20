FROM golang:1.23 as build-env
WORKDIR /go/src/github.com/adelowo/app

COPY ./go.mod /go/src/github.com/adelowo/app
COPY ./go.sum /go/src/github.com/adelowo/app

ENV GONOPROXY=github.com/adelowo/*
ENV GONOSUMDB=github.com/adelowo/*

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download && go mod verify
# COPY the source code as the last step
COPY . .


ENV CGO_ENABLED=0 

RUN go install ./cmd

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/cmd /
CMD ["/cmd"]

