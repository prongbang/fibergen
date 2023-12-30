gen_postgresql:
	dbml2sql authors/schema.dbml --postgres -o authors/schema.sql

gen_mysql:
	dbml2sql authors/schema.dbml --mysql -o authors/schema.sql

gen_crud:
	sqlc generate -f authors/sqlc.yaml