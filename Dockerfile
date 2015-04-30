FROM golang:onbuild
EXPOSE 8080

# set the number of os threads go will be using
ENV GOMAXPROCS 2

