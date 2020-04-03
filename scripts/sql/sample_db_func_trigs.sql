/******************************************************************************
 Script      : sample_db_func_trigs.sql
 Author      : David Velez
 Date        : 04/03/2020
 Description : Create function and trigger to update modify times on records
 *****************************************************************************/

-- CREATE FUNCTIONS
CREATE OR REPLACE FUNCTION update_modified_at_col()
RETURNS TRIGGER AS $$
BEGIN
	NEW.modified_at = now();
	RETURN NEW;
END;
$$ language 'plpgsql';

-- CREATE TRIGGERS
CREATE TRIGGER update_user_modtime
BEFORE UPDATE ON users
FOR EACH ROW
	EXECUTE PROCEDURE update_modified_at_col();
