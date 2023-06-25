.PHONEY build:
build:
	go build -ldflags="-s -w" -trimpath

.PHONEY out:
out:
	./go_ray_tracing.git > out.ppm

.PHONEY run:
run:
	make build && make out

.PHONEY run_dev:
run_dev:
	go run main.go

