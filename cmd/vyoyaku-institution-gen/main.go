package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/lixin9311/vac-bot/vyoyaku"
	"golang.org/x/net/context"
)

var (
	entrypoint = flag.String("entry", "131083-koto", "the url that the ward gives you to login, we only need the absolute path. ex: 131083-koto")
	outputDir  = flag.String("output", "./", "output dir")
	username   = flag.String("username", "", "username")
	password   = flag.String("password", "", "password")
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := vyoyaku.NewClient(ctx, *entrypoint)
	if err != nil {
		log.Fatalln("failed to create v-tokyo api client:", err)
	}
	err = client.Login(ctx, *username, *password)
	if err != nil {
		log.Fatalln("failed to login:", err)
	}
	for i := 0; ; i++ {
		resp, err := client.SearchInstitutions(ctx, i)
		if err != nil {
			log.Fatalln("failed to search institution:", err)
		}
		filename := fmt.Sprintf("institutions.%d.json", i)
		file, err := os.Create(path.Join(*outputDir, filename))
		if err != nil {
			log.Fatalf("failed to save %s: %v", filename, err)
		}
		data, _ := json.MarshalIndent(resp, "", "  ")
		if _, err := file.Write(data); err != nil {
			log.Fatalf("failed to write %s: %v", filename, err)
		}
		file.Close()
		if resp.DisplayStart+10 >= resp.TotalCount {
			break
		}
	}
}
