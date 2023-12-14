day1: buildDay1
	./bin/2023/day1

day2: buildDay2
	./bin/2023/day2

buildDay1:
	go build -o ./bin/2023/day1 ./2023/day1/*.go
	
buildDay2:
	go build -o ./bin/2023/day2 ./2023/day2/*.go