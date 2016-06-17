FROM golang:1.6

RUN mkdir -p /src
WORKDIR /src

COPY . /src/

EXPOSE 8082
CMD make run
