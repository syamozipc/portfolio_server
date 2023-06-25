.PHONY: migrate-new
migrate-new:
ifdef name
	sql-migrate new $(name)
else
	@echo "nameパラメータを指定してください\n例:\"make migrate-new name=create_user_table\""
endif
