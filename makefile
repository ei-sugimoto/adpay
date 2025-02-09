.PHONY: migrate
migrate:
	@echo "Migrating database..."
	@cd apps/backend && atlas schema apply -u "postgresql://user:password@localhost:5432/adpay?sslmode=disable" --to file://./adpay.hcl