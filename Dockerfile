# build stage
FROM golang:alpine AS builder
ADD . /go/src/github.com/feiskyer/openai-copilo