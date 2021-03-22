FROM mysql:latest
COPY ./initialize.sql /docker-entrypoint-initdb.d/1-initialize.sql
COPY ./inittable.sql /docker-entrypoint-initdb.d/2-initdb.sql
COPY ./my.custom.cnf /etc/mysql/conf.d/

ENV LC_ALL C.UTF-8
ENV TIMEZONE Asia/Tokyo
RUN echo "${TIMEZONE}" > /etc/timezone
RUN rm /etc/localtime && ln -s /usr/share/zoneinfo/${TIMEZONE} /etc/localtime
