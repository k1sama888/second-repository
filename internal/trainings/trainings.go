package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return fmt.Errorf("3 elements were expected, received:%d", len(parts))
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("Steps conversion error: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("The number of steps must be greater than 0, received: %d", steps)
	}
	t.Steps = steps
	t.TrainingType = parts[1]
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("Time conversion error: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("The duration must be greater than 0, received: %v", duration)
	}
	t.Duration = duration
	return nil
}
func (t Training) ActionInfo() (string, error) {
	distanceKm := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var calories float64
	var err error
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("unknown type of training")
	}
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distanceKm, speed, calories), nil

}
