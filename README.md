cpuburn lets you use 100% of all available cores, useful when http-based stress-testing.

## Installation

```bash
go get github.com/ilyakaznacheev/cpuburn-web
```

Or download a build for Linux from release.

## Usage

```bash
# turn on with all the cores
curl http://127.0.0.1:8080/on
# turn on with 2 cores
curl http://127.0.0.1:8080/on?n=2
# turn off
curl http://127.0.0.1:8080/off
```
