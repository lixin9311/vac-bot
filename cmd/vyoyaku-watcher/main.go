package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/lixin9311/vac-bot/vyoyaku"
)

var (
	entrypoint = flag.String("entry", "131083-koto", "the url that the ward gives you to login, we only need the absolute path. ex: 131083-koto")
	interval   = flag.Duration("interval", time.Minute, "interval between interations")
	fromNow    = flag.Duration("from-now", 24*time.Hour, "watch availabilities which are after the specified period from now")
	after      = flag.String("after", "", "watch availabilities which are after the specified data. ex: 2021-07-23")
	before     = flag.String("before", "", "watch availabilities which are before the specified data. ex: 2021-07-23")

	locAsiaTokyo = time.FixedZone("Asia/Tokyo", 9*3600)
)

func main() {
	ctx := context.Background()
	client, err := vyoyaku.NewClient(ctx, *entrypoint)
	if err != nil {
		log.Fatalf("failed to init v-yotaku client: %v", err)
	}
	cutoff := *fromNow
	var afterDate, beforeDate time.Time
	if *after != "" {
		afterDate, err = time.ParseInLocation("2006-01-02", *after, locAsiaTokyo)
		if err != nil {
			log.Fatalln("cannot parse from date:", err)
		}
	}
	if *before != "" {
		afterDate, err = time.ParseInLocation("2006-01-02", *before, locAsiaTokyo)
		if err != nil {
			log.Fatalln("cannot parse from date:", err)
		}
	}
	for {
		hit := false
		total := 0
		totalSlot := 0
		for k, insti := range vyoyaku.GetInstitutions() {
			// it seems like the given time range will be ignored by the api
			calendar, err := client.GetCalendar(ctx, time.Now(), time.Now().AddDate(0, 1, 0), k, 1)
			if err != nil {
				log.Fatalf("failed to get calendar: %v", err)
			}
			total += len(calendar)
			for i, v := range calendar {
				totalSlot += v.LimitCount
				if v.Empty > 0 {
					starTime, err := time.Parse("2006-01-02 15:04:05Z07:00", v.Start)
					if err != nil {
						log.Fatalf("failed to parse time: %v", err)
					}
					if !starTime.After(time.Now().Add(cutoff)) {
						continue
					} else if !afterDate.IsZero() && !starTime.After(afterDate) {
						continue
					} else if !beforeDate.IsZero() && !starTime.Before(beforeDate) {
						continue
					}
					hit = true
					log.Printf("[%d/%d]: %d availabilities on %s at %s\n", i+1, len(calendar), v.Empty, starTime.Format("2006-01-02"), insti)
				}
			}
		}
		if !hit {
			log.Printf("no luck out of %d time slots & %d vaccine shots\n", total, totalSlot)
		}
		time.Sleep(*interval)
	}
}
