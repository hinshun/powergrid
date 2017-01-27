vendor:
	glide update -v
	glide-vc --use-lock-file

run:
	go run cmd/powergrid/powergrid.go

.PHONY: vendor run
