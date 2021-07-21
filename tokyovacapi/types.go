package tokyovacapi

import (
	"encoding/json"
	"time"
)

type VacItem int

const (
	Pfizer      VacItem = 1
	AstraZeneca VacItem = 2
	Moderna     VacItem = 3
)

type Partition string

const (
	ShibuyaKu       Partition = "131130"
	KitaKu          Partition = "131172"
	SetagayaKu      Partition = "131121"
	Shinjuku        Partition = "131041"
	Sanjo           Partition = "152048"
	Musashino       Partition = "132039"
	Kumamoto        Partition = "431001"
	Fukuoak         Partition = "401307"
	Toshima         Partition = "131164"
	Minamishimabara Partition = "422142"
)

func (v VacItem) String() string {
	switch v {
	case Pfizer:
		return "Pfizer"
	case AstraZeneca:
		return "AstraZeneca"
	case Moderna:
		return "Moderna"
	default:
		return "Unknown Vaccine"
	}
}

type (
	ReservationList struct {
		Reservations []*Reservation `json:"reservations"`
	}

	LoginRequest struct {
		PartitionKey Partition `json:"partition_key"`
		RangeKey     string    `json:"range_key"`
		Password     string    `json:"password"`
	}

	Reservation struct {
		ID          int `json:"id"`
		Information struct {
			LotNumber         string `json:"lot_number"`
			Memo              string `json:"memo"`
			VaccinationDocker string `json:"vaccination_doctor"`
			VaccinationNumber string `json:"vaccination_number"`
		} `json:"information"`
		IsCancelled      bool `json:"is_cancelled"`
		IsCompleted      bool `json:"is_completed"`
		ReservationFrame struct {
			ID         int    `json:"id"`
			Name       string `json:"name"`
			Department struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Number string `json:"number"`
			} `json:"department"`
			StartAt     time.Time       `json:"start_at"`
			EndAt       time.Time       `json:"end_at"`
			Information json.RawMessage `json:"information"`
			Item        struct {
				ID       int    `json:"id"`
				Interval int    `json:"interval"`
				Name     string `json:"name"`
			} `json:"item"`
		} `json:"reservation_frame"`
		Source string `json:"source"`
		Status string `json:"status"` // completed, init
	}

	LoginResponse struct {
		AccessToken string `json:"access"`
		Code        string `json:"code"` // 100 == success, 200 & others == failure
		Person      struct {
			EmailEnabled bool            `json:"email_enabled"`
			ID           int             `json:"id"`
			Information  json.RawMessage `json:"information"`
			Organization int             `json:"organization"`
			PartitionKey Partition       `json:"partition_key"`
			RangeKey     string          `json:"range_key"`
			Reservations []*Reservation  `json:"reservation"`
		} `json:"person"`
		RefreshToken string `json:"refresh"`
		ErrorMessage string `json:"message"` // in case API error
	}

	Department struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Number      string `json:"number"`
		Information struct {
			Area              string   `json:"area"`
			Text              string   `json:"text"`
			Access            string   `json:"access"`
			Address1          string   `json:"address1"`
			Address2          string   `json:"address2"`
			Address3          string   `json:"address3"`
			Homepage          string   `json:"homepage"`
			Postcode          string   `json:"postcode"`
			PhoneNumber       string   `json:"phone_number"`
			DisplayedName     string   `json:"displayed_name"`
			NearestStations   []string `json:"nearest_station"`
			DisplayedNameKana string   `json:"displayed_name_kana"`
		} `json:"information"`
		Items []VacItem `json:"item"`
	}

	GetDepartmentsResponse struct {
		Departments []*Department `json:"department"`
	}

	ReservationFrame struct {
		ID                  int             `json:"id"`
		Name                string          `json:"name"`
		StartAt             time.Time       `json:"start_at"`
		EndAt               time.Time       `json:"end_at"`
		IsPublished         bool            `json:"is_published"`
		ReservationCnt      int             `json:"reservation_cnt"`
		ReservationCntLimit int             `json:"reservation_cnt_limit"`
		Department          int             `json:"department"`
		Item                int             `json:"item"`
		Next                json.RawMessage `json:"next"`
	}

	Error struct {
		Detail         string          `json:"detail"`
		Code           string          `json:"code"`
		Messages       json.RawMessage `json:"messages"`
		NonFieldErrors string          `json:"non_field_errors"`
	}

	AvailableFrameResponse struct {
		ReservationFrames []*ReservationFrame `json:"reservation_frame"`
	}

	AvailableDepartmentResponse struct {
		DepartmentList []int `json:"department_list"`
	}

	ReserveRequest struct {
		ReservationFrameID int `json:"reservation_frame_id"`
	}

	ReserveResponse struct {
		Reservation struct {
			ID int `json:"id"`
		} `json:"reservation"`
	}
	Token struct {
		TokenType string `json:"token_type"`
		Exp       int    `json:"exp"`
		JTI       string `json:"jti"`
		UserID    int    `json:"user_id"`
	}
	Article struct {
		Body        string          `json:"body"`
		Category    json.RawMessage `json:"category"`
		CreateAt    time.Time       `json:"created_at"`
		Description string          `json:"description"`
		Header      string          `json:"header"`
		ID          int             `json:"id"`
		Information struct {
			Important bool `json:"important"`
		} `json:"information"`
		UpdatedAt time.Time `json:"updated_at"`
		URL       string    `json:"url"`
	}
	GetArticlesResponse struct {
		Articles []*Article `json:"articles"`
	}
	GetPersonResponse struct {
		Person struct {
			EmailEnabled bool            `json:"email_enabled"`
			ID           int             `json:"id"`
			Information  json.RawMessage `json:"information"`
			Organization int             `json:"organization"`
			PartitionKey Partition       `json:"partition_key"`
			RangeKey     string          `json:"range_key"`
			Reservations []*Reservation  `json:"reservation"`
		} `json:"person"`
		RefreshToken string `json:"refresh"`
		ErrorMessage string `json:"message"` // in case API error
	}
)
