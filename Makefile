# PHONY means that it doesn't correspond to a file; it always runs the build commands.

.PHONY: build
build:
	@cd lib && cargo build --release
	@cp lib/target/release/libtiktoken_ffi.a lib/
	go build .

.PHONY: test
test:
	go test -v ./... -count=1