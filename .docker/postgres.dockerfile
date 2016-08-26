FROM postgres:latest
ENV POSTGRES_PASSWORD=admin
ENV POSTGRES_DB=goapp
ADD scripts/sql-queries/ /docker-entrypoint-initdb.d/
