# courier-service-problem2

The project is written with Go.

You can run the program by opening the executable main.exe.

If Go is installed in your machine, you can run below command to execute the program:
go run ./main.go

At the main_test.go file, function getTestData() is where I put the test input and output list.
Run 'go test -v' to execute the test function.

I also attached a sample screenshot of the program. (SampleOutput.PNG)

Sample Input:
100 5
PKG1 50 30 OFR001
PKG2 75 125 OFR008
PKG3 175 100 OFR003
PKG4 110 60 OFR002
PKG5 155 95 NA
2 70 200

Sample Output:
PKG1 0 750 3.98
PKG2 0 1475 1.78
PKG3 0 2350 1.42
PKG4 105 1395 0.85
PKG5 0 2125 4.19

Thank you~

