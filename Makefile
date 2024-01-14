migrate: silent
	sqlite3 forum.sqlite3 < migrations/add_user_table.sql
	sqlite3 forum.sqlite3 < migrations/add_post_table.sql
