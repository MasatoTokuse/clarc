#!/bin/bash
set -e
psql -U admin clarc <<EOSQL
  GRANT ALL PRIVILEGES ON DATABASE clarc TO admin;
  CREATE TABLE users(
    id           SERIAL,
    user_id      VARCHAR(255) NOT NULL,
    password     VARCHAR(255) NOT NULL,
    name         VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
  );
EOSQL
