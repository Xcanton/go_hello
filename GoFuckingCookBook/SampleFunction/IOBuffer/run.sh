go_file=$1

if [ ! -n $go_file ]; then
	go_file="RepeatLines.go"
fi

echo $go_file
if [ $go_file != "RepeatLines_Streaming.go" ]; then
	cat text.txt | go run $go_file
else
	go run $go_file text.txt
fi

