package main

import (
	"fmt"
	"math"
	"sort"

	"example.com/courier-service/model"
)

var baseDeliveryCost, noOfPackages, noOfVehicle, maxSpeed, maxCarryWeight int
var pkgDetailList = []*model.PackageDetail{}
var vehicleList = []*model.Vehicle{}
var couponMap map[string]*model.Coupon = getInitialCouponMap()

func main() {
	getInitialInput()       //Get the base delivery cost and number of package
	getPkgInputList()       //Get the list of package detail
	getFinalInput()         //Get the noOfVehicle, maxSpeed, and maxCarryWeight
	calculateDeliveryTime() //For each package detail, build the possible combination and the delivery time
	calculateDeliveryCost() //For each package detail, calculate the deliveryCost and discount

	var anyKey string
	fmt.Println("Type any character and enter to close the program...")
	fmt.Scan(&anyKey)
}

func calculateDeliveryCost() {
	fmt.Print("\nCalculating Delivery Cost...\n\n")
	for _, pkg := range pkgDetailList {
		var discount, deliveryCost float64
		deliveryCost = float64(baseDeliveryCost) + (float64(pkg.PkgWeight) * 10) + (float64(pkg.Distance) * 5)
		discount = calculateDiscount(deliveryCost, pkg)
		deliveryCost -= discount
		pkg.DeliveryCost = deliveryCost
		pkg.Discount = discount
	}

	fmt.Print("\n*****Final Result*****\n")
	for _, pkg := range pkgDetailList {
		fmt.Println(pkg.PkgId, pkg.Discount, pkg.DeliveryCost, pkg.DeliveryTime)
	}
}

func calculateDiscount(deliveryCost float64, pkgDetail *model.PackageDetail) (result float64) {
	coupon, couponExist := couponMap[pkgDetail.OfferCode]
	if couponExist {
		if pkgDetail.PkgWeight <= coupon.MaxWeight && pkgDetail.PkgWeight >= coupon.MinWeight && pkgDetail.Distance <= coupon.MaxDistance && pkgDetail.Distance >= coupon.MinDistance {
			result = deliveryCost * float64(coupon.DiscountPerc) / 100
			fmt.Printf("%v: Calculated discount is %v.\n", pkgDetail.PkgId, result)
		} else {
			fmt.Println(pkgDetail.PkgId, ": Weight or distance does not meet coupon", pkgDetail.OfferCode, "criteria.")
		}
	} else {
		fmt.Println(pkgDetail.PkgId, ": Coupon with offer code", pkgDetail.OfferCode, "does not exist.")
	}
	return result
}

func calculateDeliveryTime() {
	fmt.Print("\nCalculating Delivery Time...\n")

	//Build all the possible package combination
	allPkgCombination := []*model.PackageCombination{}
	allPkgCombination = loopPkgCombination(0, 0, allPkgCombination, 0, 0, []string{})

	//Sort combination by highest totalWeight, if same weight, sort by lowest totalDistance
	sort.Slice(allPkgCombination, func(a, b int) bool {
		if allPkgCombination[a].TotalWeight == allPkgCombination[b].TotalWeight {
			return allPkgCombination[a].TotalDistance < allPkgCombination[b].TotalDistance
		} else {
			return allPkgCombination[a].TotalWeight > allPkgCombination[b].TotalWeight
		}
	})

	//Initialize the vehicleList
	for i := 0; i < noOfVehicle; i++ {
		vehicleList = append(vehicleList, &model.Vehicle{IsAvailable: false, DeliveryStartTime: 0})
	}

	//Loop through all the possible combination
	var calculatedPackages = []string{}
	for _, pkgComb := range allPkgCombination {
		fmt.Println(pkgComb.TotalWeight, pkgComb.TotalDistance, pkgComb.PackageIDs)

		//Skip the combination if the package in the combination already calculated
		if len(calculatedPackages) > 0 && isPackageCalculated(pkgComb.PackageIDs, calculatedPackages) {
			// fmt.Println("Skip package ", pkgComb.PackageIDs)
			continue
		}

		vehicleIdx := getNextAvailableVehicle() //Find out the first available vehicle and later get it's start time
		deliveryStartTime := vehicleList[vehicleIdx].DeliveryStartTime
		var longestDeliveryTime float64
		for _, pkgId := range pkgComb.PackageIDs {
			deliveryTime := setDeliveryTime(pkgId, deliveryStartTime)
			if deliveryTime > longestDeliveryTime {
				longestDeliveryTime = deliveryTime
			}
		}
		vehicleList[vehicleIdx].DeliveryStartTime = longestDeliveryTime * 2 //Set the vehicle next available time
		calculatedPackages = append(calculatedPackages, pkgComb.PackageIDs...)
		if len(calculatedPackages) == noOfPackages {
			break
		}
	}
}

func isPackageCalculated(currentPackages []string, calculatedPackages []string) bool {
	for _, calPkgId := range calculatedPackages {
		for _, curPkgId := range currentPackages {
			if curPkgId == calPkgId {
				return true
			}
		}
	}
	return false
}

func getNextAvailableVehicle() int { //Return the index of the available vehicle
	var nearestDeliveryTime float64 = 0
	var nearestIdx int
	for idx, vehicle := range vehicleList {
		// fmt.Println(idx, "check time: ", vehicle.DeliveryStartTime)
		if vehicle.DeliveryStartTime == 0 { //If 0 means vehicle is available straight away
			return idx
		}
		if nearestDeliveryTime == 0 || vehicle.DeliveryStartTime < nearestDeliveryTime {
			nearestDeliveryTime = vehicle.DeliveryStartTime
			nearestIdx = idx
		}
	}
	return nearestIdx
}

func setDeliveryTime(pkgId string, deliveryStartTime float64) float64 {
	for _, pkgDetail := range pkgDetailList {
		if pkgDetail.PkgId == pkgId {
			var deliveryTime float64 = deliveryStartTime + (float64(pkgDetail.Distance) / float64(maxSpeed))
			deliveryTime = math.Floor(deliveryTime*100) / 100
			pkgDetail.DeliveryTime = deliveryTime
			return deliveryTime
		}
	}
	return 0
}

func loopPkgCombination(loopIdx int, previousPkgIdx int, allPkgCombination []*model.PackageCombination, sumOfWeight int, sumOfDistance int, pkgArray []string) []*model.PackageCombination {
	pkgList := pkgDetailList[loopIdx:] //Only loop
	for pkgIdx, pkg := range pkgList {
		if loopIdx > 0 && (pkgIdx+loopIdx) <= previousPkgIdx { //So that we wont add the same PackageId or package combination in different order
			continue
		}
		var newSumOfWeight int = sumOfWeight + pkg.PkgWeight
		var newSumOfDistance int = sumOfDistance + pkg.Distance
		var newPkgArray []string = append(pkgArray, pkg.PkgId)
		if newSumOfWeight <= maxCarryWeight {
			pkgCombination := &model.PackageCombination{TotalWeight: newSumOfWeight, TotalDistance: newSumOfDistance, PackageIDs: newPkgArray}
			allPkgCombination = append(allPkgCombination, pkgCombination)
		}

		if loopIdx < (noOfPackages / noOfVehicle) {
			allPkgCombination = loopPkgCombination((loopIdx + 1), (pkgIdx + loopIdx), allPkgCombination, newSumOfWeight, newSumOfDistance, newPkgArray)
		}
	}
	return allPkgCombination
}

func getInitialInput() {
	fmt.Println("Please enter Base Delivery Cost and Number of Package separated by space: ")
	_, err := fmt.Scan(&baseDeliveryCost, &noOfPackages)

	if err != nil {
		fmt.Print("Base Delivery Cost and Number of Package must be a number...\n\n")
		err = nil
		getInitialInput()
	}
}

func getPkgInputList() {
	var pkgId, offerCode string
	var pkgWeight, distance int
	fmt.Println("Please enter", noOfPackages, "lines of package details:")

	for i := 0; i < noOfPackages; i++ {
		_, err := fmt.Scan(&pkgId, &pkgWeight, &distance, &offerCode)

		if err != nil {
			fmt.Print("Package detail not in \"string int int string\" format, please enter the list again...\n\n")
			pkgDetailList = []*model.PackageDetail{}
			i = -1
			err = nil
		} else {
			pkgDetailList = append(pkgDetailList, &model.PackageDetail{PkgId: pkgId, PkgWeight: pkgWeight, Distance: distance, OfferCode: offerCode})
			// fmt.Printf("Saved %v...\n", pkgId)
		}
	}
}

func getFinalInput() {
	fmt.Println("Please enter Number of Vehicle, Max Speed, and Max Carriable Weight separated by space: ")
	_, err := fmt.Scan(&noOfVehicle, &maxSpeed, &maxCarryWeight)

	if err != nil {
		fmt.Print("Number of Vehicle, Max Speed, and Max Carriable Weight must be a number...\n\n")
		err = nil
		getFinalInput()
	}
}

func getInitialCouponMap() map[string]*model.Coupon {
	var couponMap = make(map[string]*model.Coupon)

	couponMap["OFR001"] = &model.Coupon{OfferCode: "OFR001", DiscountPerc: 10, MinDistance: 0, MaxDistance: 199, MinWeight: 70, MaxWeight: 200}
	couponMap["OFR002"] = &model.Coupon{OfferCode: "OFR002", DiscountPerc: 7, MinDistance: 50, MaxDistance: 150, MinWeight: 100, MaxWeight: 250}
	couponMap["OFR003"] = &model.Coupon{OfferCode: "OFR003", DiscountPerc: 5, MinDistance: 50, MaxDistance: 250, MinWeight: 10, MaxWeight: 150}

	return couponMap
}
