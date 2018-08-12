# wasm-lifegame

Very simple lifegame by wasm(golang).

### Build

```bash
$ GOOS=js GOARCH=wasm go build -o test.wasm main.go
// or
$ make build_wasm
```

### Run

```bash
$ make serve
// or
$ make run // Build & Run
```
