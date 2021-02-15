FROM golang:alpine3.13

RUN  mkdir /combined_data_mt
COPY ./src /src
WORKDIR /src
# db connection string should be passed as arg to docker run
ENTRYPOINT ["go", "run", "main.go", "-data=/combined_data_mt", "-db="]