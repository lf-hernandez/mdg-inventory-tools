#!/bin/bash

psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -c "\copy item(item_id, description, quantity, price) FROM '/docker-entrypoint-initdb.d/cleaned-source.csv' DELIMITER ',' CSV HEADER"
