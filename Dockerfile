FROM golang:1.21-alpine

# membuat direktori folder
RUN mkdir /app

# set working directory
WORKDIR /app

COPY ./ /app

RUN go mod tidy

# create executable
RUN go build -o beapi

CMD [ "./beapi" ]