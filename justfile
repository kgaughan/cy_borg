build:
	goreleaser build --auto-snapshot --single-target --clean

clean:
	rm -rf dist
