FROM python:3.9 as data-processing
WORKDIR /app
COPY source.csv /app/
COPY clean_csv_source.py /app/
COPY requirements.txt /app/
RUN pip install -r requirements.txt
RUN python clean_csv_source.py

FROM postgres:16
COPY --from=data-processing /app/cleaned-source.csv /docker-entrypoint-initdb.d/
COPY init.sql /docker-entrypoint-initdb.d/
COPY load_data.sh /docker-entrypoint-initdb.d/
RUN chmod +x /docker-entrypoint-initdb.d/load_data.sh

ENV POSTGRES_DB=mdg
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres

EXPOSE 5432
CMD ["postgres"]
