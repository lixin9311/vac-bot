package vyoyaku

import (
	"embed"
	"encoding/json"
	"log"
)

// please be careful about the file names
//go:embed institutions.*.json
var institutionsData embed.FS

var institutions = map[string]string{}

func GetInstitutions() map[string]string {
	return institutions
}

func init() {
	files, err := institutionsData.ReadDir(".")
	if err != nil {
		log.Fatalln("failed to load institution data", err)
	}
	for _, v := range files {
		data, err := institutionsData.ReadFile(v.Name())
		if err != nil {
			log.Fatalln("failed to load institution data", err)
		}
		institutionsJson := &InstitutionData{}
		if err := json.Unmarshal(data, institutionsJson); err != nil {
			log.Fatalln("failed to load institution data", err)
		}
		for _, v := range institutionsJson.Data {
			if v.ReservationReception == "〇" {
				institutions[v.MedicalCenterCd] = v.MedicalCentername
			}
		}
	}
	log.Printf("%d institutions loaded\n", len(institutions))
}
