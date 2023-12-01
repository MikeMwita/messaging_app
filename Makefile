
postgres:
	docker  run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root expenseapp



.PHONY : postgres createdb