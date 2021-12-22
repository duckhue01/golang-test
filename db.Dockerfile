
FROM mysql:8.0.23

COPY ./*.sql /docker-entrypoint-initdb.d/