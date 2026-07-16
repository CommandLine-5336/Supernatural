# Database for Supernatural

## Description
Sets up database service for Supernatural with Docker. Creates database with a user, migrates models from Django application into created database.

## Steps
- First, change SECRET_KEY, DB_USER, DB_PASS in .env to values from AWS Secrets Manager. Also change user and
password in my.cnf to same values from AWS Secrets.
- Then run `docker compose up --build`
