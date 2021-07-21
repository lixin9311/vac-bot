package vyoyaku

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

var (
	entry    = "/131083-koto"
	username = ""
	password = ""
)

func TestLogin(t *testing.T) {
	ctx := context.Background()
	client, err := NewClient(ctx, entry)
	if err != nil {
		t.Errorf("failed to create client: %v", err)
	}
	if err := client.Login(ctx, username, password); err != nil {
		t.Errorf("failed to login: %v", err)
	}
}

func TestCalendar(t *testing.T) {
	ctx := context.Background()
	client, err := NewClient(ctx, entry, WithCSRFToken("VvP1d7KE7fCpI91VOWBUmu00P0XqCccHfBwFDIkd"))
	if err != nil {
		t.Errorf("failed to create client: %v", err)
	}

	for k, insti := range institutions {
		// k := "1310870071"
		// fmt.Printf("============== %s ==============\n", v)
		calendar, err := client.GetCalendar(ctx, time.Now(), time.Now().AddDate(0, 1, 0), k, 1)
		if err != nil {
			t.Errorf("failed to get calendar: %v", err)
		}
		for i, v := range calendar {
			if v.Empty > 0 {
				starTime, err := time.Parse("2006-01-02 15:04:05Z07:00", v.Start)
				if err != nil {
					log.Fatalf("failed to parse time: %v", err)
				}
				if starTime.Before(time.Now().AddDate(0, 1, 0)) {
					fmt.Printf("[%d/%d]: %d availabilities on %s at %s\n", i+1, len(calendar), v.Empty, starTime.Format("2006-01-02"), insti)
				}
			}
		}
	}
}
