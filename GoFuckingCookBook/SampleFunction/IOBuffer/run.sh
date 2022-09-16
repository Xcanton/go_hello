go_file=$1

if [ -n $go_file ]; then
	go_file="RepeatLines.go"
fi

echo $go_file
cat text.txt | go run $go_file
