FROM golang
RUN apt-get update && apt-get install -y postgresql-client
COPY ./migrations ./migrations
COPY ./parts ./bin
COPY ./migrate ./bin

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
RUN chmod +x /wait

CMD /wait && ./bin/migrate -database postgres://postgres:123@postgres:5432/parts?sslmode=disable -path ./migrations/ up && chmod +x ./bin/parts && ./bin/parts