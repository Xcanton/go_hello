go_file=$1

if [ -n $go_file]; then
	go_file="RepeatLines.go"
fi

cat text.txt | go run $go_file
