package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("The number of steps must be greater than zero, received: %d", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("The weight must be greater than zero, received: %.2f", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("The height must be greater than zero, received: %.2f", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("The duration must be greater than zero, received: %v", duration)
	}
	speed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	walkingCalories := ((weight * speed * durationInMinutes) / float64(minInH)) * walkingCaloriesCoefficient
	return walkingCalories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("The number of steps must be greater than zero, received: %d", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("The weight must be greater than zero, received: %.2f", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("The height must be greater than zero, received: %.2f", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("The duration must be greater than zero, received: %v", duration)
	}
	speed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	runningCalories := (weight * speed * durationInMinutes) / float64(minInH)
	return runningCalories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	distanceKm := Distance(steps, height)
	durationInHours := duration.Hours()
	return distanceKm / durationInHours
}

func Distance(steps int, height float64) float64 {
	distanceKm := (height * stepLengthCoefficient) * float64(steps) / float64(mInKm)
	return distanceKm

}
