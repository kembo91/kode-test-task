CREATE ROLE postgres;
ALTER ROLE postgres WITH PASSWORD 'postgres';
ALTER ROLE postgres WITH LOGIN;
CREATE DATABASE postgres;
CREATE DATABASE test;
GRANT ALL PRIVILIGES ON DATABASE postgres to postgres;
GRANT ALL PRIVILIGES ON DATABASE test to postgres;
