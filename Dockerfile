FROM golang:onbuild
RUN go install github.com/MOZGIII/dphx/cmd/dphx
EXPOSE 1080
