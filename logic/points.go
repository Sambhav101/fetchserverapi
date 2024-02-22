package points

import (
	"fetchapi/receipt"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// return one point for every alphanumeric character in retailer name
func pointsFromName(retailer string) int {
	points  := 0;
	for _, char := range retailer {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			points += 1
		}
	}
	return points
}

func pointsFromTotal(total string) int {
	points := 0;
	totalPrice, _ := strconv.ParseFloat(total, 64)
	if totalPrice == 0 {
		return points
	}
	if float64(int(totalPrice)) == totalPrice {
		points += 50
	}
	if math.Mod(totalPrice, 0.25) == 0 {
		points += 25
	}
	return points
}

func pointsFromItems(items []receipt.Item) int {
	points  := len(items)/2 * 5
	for _, item := range items {
		desc := strings.TrimSpace(item.ShortDescription)
		price, _ := strconv.ParseFloat(item.Price, 64)
		if len(desc)%3 == 0 {
			points += int(math.Ceil(0.2 * price))
		}
	}
	return points
}

func pointsFromTime(purchaseDate, purchaseTime, startTime, endTime string) int {
	points := 0
	pDate, _ := time.Parse("2006-01-02", purchaseDate)
	pTime, _ := time.Parse("15:04", purchaseTime)
	sTime, _ := time.Parse("15:04", startTime)
	eTime, _ := time.Parse("15:04", endTime)
	if pDate.Day() % 2 != 0 {
		points += 6
	}
	if pTime.After(sTime) && pTime.Before(eTime) {
		points += 10
	}
	return points
}

func CalculatePoints(receipt receipt.Receipt) int {
	startTime := "14:00"
	endTime := "16:00"
	totalPoints := pointsFromName(receipt.Retailer) + 
	pointsFromTotal(receipt.Total) + 
	pointsFromTime(receipt.PurchaseDate, receipt.PurchaseTime, startTime, endTime) +
	pointsFromItems(receipt.Items)
	return totalPoints
}