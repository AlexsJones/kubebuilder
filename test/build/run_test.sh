
child_pid1=0
child_pid2=0

other_commands() {
    printf "\rSIGINT caught\n"
    kill -9 $child_pid1
    kill -9 $child_pid2

}

trap 'other_commands' SIGINT

echo "Starting subscriber..."
go run main.go &
child_pid1=$!
sleep 3
echo "Starting publisher..."
go run test/build/main.go
child_pid2=$!
