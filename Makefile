vendor:
	glide update -v
	glide-vc --use-lock-file

run:
	go run powergrid.go

.PHONY: vendor run
