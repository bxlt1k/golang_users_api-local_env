FROM mysql:latest

COPY ./docker/services/mysql/init/init.sh /mysql/
RUN chmod 777 /mysql/init.sh

ENTRYPOINT /mysql/init.sh