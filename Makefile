createdb:
	docker exec -it postgres14 createdb --username=root --owner=root crowdfunding

dropdb:
	docker exec -it postgres14 dropdb crowdfunding

.PHONY: createdb dropdb