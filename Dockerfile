# Dockerfile Example
# https://medium.com/@petomalina/using-go-mod-download-to-speed-up-golang-docker-builds-707591336888
# Based on this image: https:/hub.docker.com/_/golang/
FROM golang:latest as builder


RUN mkdir -p /go/github.com/bilalislam/todo-app-api
WORKDIR /go/github.com/bilalislam/todo-app-api

# Force the go compiler to use modules
ENV GO111MODULE on
# <- COPY go.mod and go.sum files to the workspace
COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .


# Compile application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./todo-app-api main.go

RUN chmod +x /go/github.com/bilalislam/todo-app-api

# Execite application when container is started(NOT SCRATCH)
#ENTRYPOINT /go/src/git.hepsiburada.com/Checkout/murtlap/checkout_order



#Image Diff
#(Not Scratch)co-agreement-api 1.23GB
#(Scratch    )co-agreement-api 34.3MB
# <- Second step to build minimal image
FROM scratch
WORKDIR /root/
COPY --from=builder /go/github.com/bilalislam/todo-app-api .
# Execite application when container is started
CMD ["./todo-app-api"]


EXPOSE 8080
