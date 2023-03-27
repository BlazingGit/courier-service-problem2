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

Thought Process:
The first problem require us to calculate the delivery cost based on the given weight, distance, and discount coupon. Weight and distance are straight forward.
To calculate correctly based on the coupon offerId, I store the information of the coupon into a Map so I can retrieve it easily based on the user's input.
Then I check if the given weight/distance matched the coupon's criteria, if yes then I applied the discount to the delivery cost calculation.

For the second problem, vehicle, maximum speed, and maximum carriable weight are introduced to calculate the time packages would be delivered. In addition to the problem1 calculation, I first find out what could be the possible combination of packages that falls within the max carriable weight. I've decided to set the optimum number of packages per combination/trip is (Number of Packages/Number of vehicle). Example if there are 9 packages and 2 vehicle available, at most 1 vehicle can carry 5 packages, given that the weight is within the max carriable weight.

After that, I will sort the combination by weight in descending order, if the weights are the same, then I will sort them by distance in ascending order.
Then I will loop the package combination to calculate the delivery time. Need to take note that the time taken for the vehicle to be available is getting the furthest destination devide by the max speed then multiply by 2. If 2 packages located at the same distance, I will assume they are delivered at the same time.

Thank you~

