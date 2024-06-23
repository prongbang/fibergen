gen_authors_postgresql:
	dbml2sql authors/schema.dbml --postgres -o authors/schema.sql

gen_authors_mysql:
	dbml2sql authors/schema.dbml --mysql -o authors/schema.sql

gen_authors_crud:
	sqlc generate -f authors/sqlc.yaml

gen_users_postgresql:
	dbml2sql users/schema.dbml --postgres -o users/schema.sql

gen_users_mysql:
	dbml2sql users/schema.dbml --mysql -o users/schema.sql

gen_users_crud:
	sqlc generate -f users/sqlc.yaml