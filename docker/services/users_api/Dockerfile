FROM golang:1.18

COPY ./docker/utils/air /usr/bin/air
COPY ./services/users_api /users_api

WORKDIR /users_api

CMD ["air", "-d"]