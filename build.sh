{
    rm ./main
    go build -o ./main
    ./main
} || {
    echo "This is the error"
}