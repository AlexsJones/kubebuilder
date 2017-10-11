
child_pid1=0
child_pid2=0

other_commands() {
    printf "\rSIGINT caught\n"
    kill -9 $child_pid1
    kill -9 $child_pid2

}

trap 'other_commands' SIGINT

echo "Starting consumer..."
go run main.go &
child_pid1=$!
sleep 3
echo "Starting publiser..."
go run test/main.go
child_pid2=$!
