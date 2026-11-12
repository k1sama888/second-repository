package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",") //делим строку на слайс строк
	if len(parts) != 2 {
		return fmt.Errorf("2 elements were expected, received:%d", len(parts))
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("Steps conversion error: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("The number of steps must be greater than 0, received: %d", steps)
	}
	ds.Steps = steps
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("Time conversion error: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("The duration must be greater than 0, received: %v", duration)
	}
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distanceKm := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distanceKm, calories), nil
}
