# Database for Supernatural

## Description
Sets up database service for Supernatural with Docker. Creates database with a user, migrates models from Django application into created database.

## Steps
- First, create database/.env and database/my.cnf files.
- Then run `docker compose up --build`

## Files content
- .env: `SECRET_KEY='*'
DB_NAME='supernatural_db'
DB_USER='*'
DB_PASS='*'
DB_HOST='database'
`
- my.cnf: `[client]
database = supernatural_db
user = *
password = *
host = db
port = 3306`
- Get SECRET_KEY, DB_USER/user, DB_PASS/password values from AWS Secrets Manager.
