{
    rm ./main
    go build ./main.go
    ./main
} || {
    echo "This is the error"
}