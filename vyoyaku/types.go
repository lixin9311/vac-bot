package vyoyaku

type (
	errorable interface {
		GetIsError() bool
		GetErrorMsg() string
	}
	CharangecdResponse struct {
		Result   bool   `json:"result"`
		Data     string `json:"data"`
		IsError  bool   `json:"is_error"`
		ErrorMsg string `json:"error_msg"`
	}
	CheckAuthResponse struct {
		Result     bool   `json:"result"`
		Data       string `json:"data"`
		IsError    bool   `json:"is_error"`
		ErrorMsg   string `json:"error_msg"`
		LoginToken string `json:"login_token"`
	}
	InstitutionData struct {
		Result       bool           `json:"result"`
		Data         []*Institution `json:"data"`
		IsError      bool           `json:"is_error"`
		ErrorMsg     string         `json:"error_msg"`
		TotalCount   int            `json:"total_count"`
		DisplayStart int            `json:"iDisplayStart"`
	}
	Institution struct {
		MedicalCentername    string `json:"medical_center_name"`
		MedicalCenterCd      string `json:"medical_center_cd"`
		ReservationReception string `json:"reservation_reception"`
	}
	CalendarItem struct {
		AllDay          int    `json:"allDay"`
		BackgroundColor string `json:"backgroundColor"`
		Color           string `json:"color"`
		Empty           int    `json:"empty"`
		Start           string `json:"start"`
		End             string `json:"end"`
		LimitCount      int    `json:"limitCount"`
		Rendering       string `json:"rendering"`
		ReservedCount   int    `json:"reservedCount"`
		Title           string `json:"title"`
	}
	CalendarResponse []*CalendarItem
)

func (r *CharangecdResponse) GetIsError() bool {
	return r.IsError
}
func (r *CharangecdResponse) GetErrorMsg() string {
	return r.ErrorMsg
}
func (r *CheckAuthResponse) GetIsError() bool {
	return r.IsError
}
func (r *CheckAuthResponse) GetErrorMsg() string {
	return r.ErrorMsg
}
func (r *InstitutionData) GetIsError() bool {
	return r.IsError
}
func (r *InstitutionData) GetErrorMsg() string {
	return r.ErrorMsg
}
