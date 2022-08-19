package main

import (
	"fmt"
	"time"
)

var (
	weekday, timeFormat                       string
	cMonth                                    time.Month
	cYear, cDate, cHour, cMin, cSec, cNanoSec int
	localTime                                 time.Location
)

func init() {
	weekday = time.Now().Weekday().String()
	timeFormat = time.RFC3339 //time.RFC3339Nano -> (2021-08-15T02:30:45Z)
	// OR "1996-12-31 10:10:30", "1996-12-31 10:10:30 pm"
	cYear = 2022
	cMonth = 3
	cDate = 12
	cHour = 11
	cMin = 43
	cSec = 54
	cNanoSec = 100
	localTime = *time.Local
}

func workingDays() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its a weekEND")
	default:
		fmt.Println("its a not a weekEND")
	}
}

func main() {
	currentTime()
}

func currentTime() {
	currentTime := time.Now()

	fmt.Println("The year is", currentTime.Year())
	fmt.Println("The month is", currentTime.Month()) //August,Dec...
	fmt.Println("The day is", currentTime.Day())     //Date 23,14...
	fmt.Println("The hour is", currentTime.Hour())
	fmt.Println("The minute is", currentTime.Hour())
	fmt.Println("The second is", currentTime.Second())
	fmt.Printf("Today is %s\n", weekday) //Day Monday, Tue...
	fmt.Println("Time is:", currentTime.Format(timeFormat))
	//OR
	fmt.Printf("%d-%d-%d %d:%d:%d\n", currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(),
		currentTime.Hour(), currentTime.Second())
}

func specificNow() {
	//(year int, month time.Month, day int, hour int, min int, sec int, nsec int, loc *time.Location)
	theTime := time.Date(cYear, cMonth, cDate, cHour, cMin, cSec, cNanoSec, &localTime)
	fmt.Println("The time is", theTime)

	//OR
	fmt.Println(theTime.Format(timeFormat)) //prints current time with given format
}

func parseDateTimeANdFormat() {
	timeServer := "2021-08-15 02:30:45"
	theTime, err := time.Parse(timeFormat, timeServer) //returns time.Time then call time.Format
	if err != nil {
		fmt.Println("Could not parse time:", err)
	}
	fmt.Println("The time is", theTime)

	fmt.Println(theTime.Format(time.RFC3339Nano))

	utcTime := theTime.UTC()
	fmt.Println("The UTC time is", utcTime)
	fmt.Println(utcTime.Format(time.RFC3339Nano))
	//converting back
	localTime := utcTime.Local()
	fmt.Println("The Local time is", localTime)
}

func compareDates() {
	firstTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.UTC)
	fmt.Println("The first time is", firstTime)

	secondTime := time.Date(2021, 12, 25, 16, 40, 55, 200, time.UTC)
	fmt.Println("The second time is", secondTime)

	fmt.Println("First time before second?", firstTime.Before(secondTime))
	fmt.Println("First time after second?", firstTime.After(secondTime))
	//also
	fmt.Println("Duration between first and second time is", firstTime.Sub(secondTime))
}

func addSubTime() {
	theTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.UTC)
	toAdd := 1 * time.Hour
	fmt.Println("1:", toAdd)

	toAdd += 1 * time.Minute
	fmt.Println("2:", toAdd)

	toAdd += 1 * time.Second
	fmt.Println("3:", toAdd)

	nextDay := 24 * time.Hour
	theTime.Add(nextDay)

	previousDay := -24 * time.Hour
	theTime.Add(previousDay)

}
