vendor:
	glide update -v
	glide-vc --use-lock-file

run:
	go run gowergrid.go

.PHONY: vendor
