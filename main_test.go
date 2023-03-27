package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"example.com/courier-service/model"
)

func TestCalculateDeliveryCost(t *testing.T) {
	testDataList := getTestData()

	for i, testData := range testDataList {
		processTestData(testData)
		calculateDeliveryTime()
		calculateDeliveryCost()

		fmt.Print("\n***Comparing output and expected***\n", i)
		for j, expected := range testData.ExpectedOutput {
			output := pkgDetailList[j].PkgId +
				" " +
				strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.2f", pkgDetailList[j].Discount), "0"), ".") +
				" " +
				strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.2f", pkgDetailList[j].DeliveryCost), "0"), ".") +
				" " +
				strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.2f", pkgDetailList[j].DeliveryTime), "0"), ".")

			if expected != output {
				t.Errorf("Dataset %v output %v not equal to expected %v", i, output, expected)
			}
		}
	}
}

func processTestData(testData *model.TestData) {
	pkgDetailList = []*model.PackageDetail{}
	for i, input := range testData.Input {
		s := strings.Split(input, " ")
		if i == 0 {
			base, _ := strconv.Atoi(s[0])
			baseDeliveryCost = base
			noOfPkg, _ := strconv.Atoi(s[1])
			noOfPackages = noOfPkg

		} else if i == (len(testData.Input) - 1) {
			vehicleNo, _ := strconv.Atoi(s[0])
			noOfVehicle = vehicleNo
			speed, _ := strconv.Atoi(s[1])
			maxSpeed = speed
			maxWeight, _ := strconv.Atoi(s[2])
			maxCarryWeight = maxWeight

		} else {
			weight, _ := strconv.Atoi(s[1])
			distance, _ := strconv.Atoi(s[2])
			pkgDetailList = append(pkgDetailList, &model.PackageDetail{PkgId: s[0], PkgWeight: weight, Distance: distance, OfferCode: s[3]})
		}
	}
}

func getTestData() []*model.TestData {
	input := []string{
		"100 5",
		"PKG1 50 30 OFR001",
		"PKG2 75 125 OFR008",
		"PKG3 175 100 OFR003",
		"PKG4 110 60 OFR002",
		"PKG5 155 95 NA",
		"2 70 200",
	}

	expectedOutput := []string{
		"PKG1 0 750 3.98",
		"PKG2 0 1475 1.78",
		"PKG3 0 2350 1.42",
		"PKG4 105 1395 0.85",
		"PKG5 0 2125 4.19",
	}

	testDataList := []*model.TestData{}
	testDataList = append(testDataList, &model.TestData{Input: input, ExpectedOutput: expectedOutput})
	return testDataList
}
