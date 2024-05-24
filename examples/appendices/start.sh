cargo build --target=wasm32-unknown-unknown --release

#go run simple.go

go build -ldflags "-s -w" simple.go

num=1000000

if [ -n "$1" ];then
  num=$1
fi 

./simple -count=$num
