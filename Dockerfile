FROM postgres:16

COPY ./init.sql /docker-entrypoint-initdb.d/

ENV POSTGRES_DB=mdg
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres

EXPOSE 5432

CMD ["postgres"]
