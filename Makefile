# ===================
# VARIABLES
# include .envrc

.PHONY: kill
kill: 
	@kill -9 `lsof -t -i :5000`

# =======================
# DEVELOPMENT 
# ======================+
.PHONY: it 
it: 
	@go run . -srv-port=5000 -srv-env=development -srv-read-timeout=45s -srv-write-timeout=90s -srv-idle-timeout=120s

