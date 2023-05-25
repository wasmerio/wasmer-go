cargo build --target=wasm32-unknown-unknown --release

#go run simple.go

go build -ldflags "-s -w" simple.go

./simple -count=$1
