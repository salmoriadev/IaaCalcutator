package main

import "fmt"

// CalculateUpdatedIAA calculates the IAA after adding new courses.
func CalculateUpdatedIAA(currentIAA float64, completedCredits int, courses []Course) (float64, error) {
	if completedCredits < 0 {
		return 0, fmt.Errorf("completed credits must be non-negative")
	}

	if completedCredits == 0 && len(courses) == 0 {
		return 0, fmt.Errorf("add credits to calculate")
	}

	// Calculate total points.
	totalPoints := currentIAA * float64(completedCredits)
	totalCredits := completedCredits

	for _, c := range courses {
		totalPoints += c.Grade * float64(c.Credits)
		totalCredits += c.Credits
	}

	if totalCredits == 0 {
		return 0, fmt.Errorf("no credits recorded")
	}

	return totalPoints / float64(totalCredits), nil
}

// CalculateIAATarget calculates the required average to reach a target IAA.
func CalculateIAATarget(currentIAA float64, completedCredits int, semesterCredits int, target float64) (float64, float64, error) {
	if semesterCredits <= 0 {
		return 0, 0, fmt.Errorf("semester credits must be positive")
	}

	totalCredits := completedCredits + semesterCredits
	currentPoints := currentIAA * float64(completedCredits)
	requiredPoints := target * float64(totalCredits)
	pointsNeeded := requiredPoints - currentPoints
	requiredAverage := pointsNeeded / float64(semesterCredits)

	return requiredAverage, pointsNeeded, nil
}
