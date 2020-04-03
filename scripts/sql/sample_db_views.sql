/******************************************************************************
 Script      : sample_db_views.sql
 Author      : David Velez
 Date        : 04/03/2020
 Description : Create View for Users Table and sort by ID as the order may
               change when updating records
 *****************************************************************************/

-- CREATE VIEWS
CREATE VIEW user_by_id
	SELECT *
	FROM users
	ORDER BY id;
