build_wasm:
	GOOS=js GOARCH=wasm go1.11beta3 build -o test.wasm main.go

serve:
	go1.11beta3 run server.go

run:
	make build_wasm
	make serve