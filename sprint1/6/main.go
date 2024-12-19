package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

type Report struct {
	User
	ReportID int
	Date     string
}

var lastReportID int

func CreateReport(user User, reportDate string) Report {
	lastReportID++
	return Report{user, lastReportID, reportDate}
}

func GenerateUserReports(users []User, reportDate string) []Report {
	reports := make([]Report, 0)

	for _, user := range users {
		reports = append(reports, CreateReport(user, reportDate))
	}

	return reports
}

func PrintReport(report Report) {
	fmt.Println(report)
}
