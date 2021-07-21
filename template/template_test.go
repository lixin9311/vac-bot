package template

import (
	"testing"
	"time"
)

func TestSetting(t *testing.T) {
	setting := &Setting{
		Name:          "Lucas Lee",
		ID:            1234567,
		RangeKey:      "1234",
		PartitionName: "1234",
		Reservations: []*Reservation{
			{
				ID:         1234,
				Department: "Somewhere",
				Date:       time.Now(),
			},
			{
				ID:         1234,
				Department: "Somewhere",
				Date:       time.Now(),
			},
		},
	}

	t.Logf("%s", setting.Render())
}

func TestAvailability(t *testing.T) {
	avail := &Availabilities{
		Availabilities: []*Availability{
			{
				ID:                1234,
				Department:        "somewhere",
				Date:              time.Now(),
				NumOfAvailability: 1,
			},
			{
				ID:                1234,
				Department:        "somewhere",
				Date:              time.Now(),
				NumOfAvailability: 1,
			},
		},
		LoggedIn: false,
	}
	t.Logf("%s", avail.Render())
}

func TestReservation(t *testing.T) {
	avail := &Reservation{
		ID:         1234,
		Department: "1234",
		Date:       time.Now(),
	}
	t.Logf("%s", avail.Render())
}
