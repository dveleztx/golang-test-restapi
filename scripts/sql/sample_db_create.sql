/******************************************************************************
 Script      : sample_db_create.sql
 Author      : David Velez
 Date        : 03/16/2020
 Description : Setup password for postgres user, create table, and insert data
 *****************************************************************************/

-- SET PASSWORD (THIS IS ONLY FOR DEMONSTRATION PURPOSES)
ALTER USER postgres WITH PASSWORD 'postgres';

-- CREATE TABLE
CREATE TABLE users (
	id		int primary key,
	name		varchar(30) not null,
	age		int,
	created_at	char(10),
	updated_at	char(10)
);

-- INSERT DATA
COPY users (id, name, age, created_at, updated_at)
FROM '/tmp/users_db.csv' DELIMITER ',' CSV HEADER;

-- GRANT PRIVILEGES
GRANT ALL PRIVILEGES ON DATABASE sample_db TO postgres;
