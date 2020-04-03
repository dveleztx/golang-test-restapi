/******************************************************************************
 Script      : sample_db_create.sql
 Author      : David Velez
 Date        : 03/16/2020
 Description : Setup password for postgres user, create table, and insert data
 *****************************************************************************/

-- SET PASSWORD (THIS IS ONLY FOR DEMONSTRATION PURPOSES)
ALTER USER postgres WITH PASSWORD 'postgres';

-- GRANT PRIVILEGES
GRANT ALL PRIVILEGES ON DATABASE sample_db TO postgres;

-- CREATE TABLE
CREATE TABLE users (
	id		serial primary key,
	name		varchar(50) not null,
	age		int,
	created_at	timestamp with time zone default current_timestamp,
	updated_at	timestamp with time zone default current_timestamp
);

-- INSERT DATA
COPY users (name, age)
FROM '/tmp/users_db.csv' DELIMITER ',' CSV HEADER;
