package vyoyaku

import (
	"bytes"
	"context"
	"crypto/sha1"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
)

var (
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"
)

type Option func(c *Client)

func WithCSRFToken(token string) Option {
	return func(c *Client) {
		c.csrfToken = token
	}
}

func WithUserToken(token string) Option {
	return func(c *Client) {
		c.userToken = token
	}
}

type Client struct {
	client    *resty.Client
	baseURL   string
	csrfToken string
	userToken string
}

// /131083-koto
func NewClient(ctx context.Context, entrypoint string, opts ...Option) (*Client, error) {
	certpool, err := x509.SystemCertPool()
	if err != nil {
		return nil, fmt.Errorf("failed to read system cert pool: %w", err)
	}
	entrypoint = "/" + strings.TrimPrefix(entrypoint, "/")
	baseURL := strings.TrimSuffix("https://v-yoyaku.jp/", "/")
	client := resty.New().
		SetRedirectPolicy(resty.FlexibleRedirectPolicy(5)).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("User-Agent", userAgent).
		SetTLSClientConfig(&tls.Config{
			RootCAs:    certpool,
			MinVersion: tls.VersionTLS12,
		})
	c := &Client{client: client, baseURL: baseURL}
	for _, fn := range opts {
		fn(c)
	}
	if err := c.init(ctx, entrypoint); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) Login(ctx context.Context, username, password string) error {
	api1 := "/certifyapi/charangecd"
	req1 := map[string]string{
		"seq":      "1",
		"login_id": username,
	}
	result1 := &CharangecdResponse{}
	_, err := c.do(ctx, "POST", api1, req1, result1)
	if err != nil {
		log.Println("login step 1 failed:", err)
		return err
	}
	data := result1.Data + password
	hashed := sha1.Sum([]byte(data))
	api2 := "/certifyapi/checkauth"
	req2 := map[string]string{
		"login_id":    username,
		"seq":         "2",
		"charange_cd": result1.Data,
		"hash_cd":     hex.EncodeToString(hashed[:]),
	}
	result2 := &CheckAuthResponse{}
	_, err = c.do(ctx, "POST", api2, req2, result2)
	if err != nil {
		log.Println("login step 2 failed:", err)
		return err
	}
	c.userToken = result2.LoginToken
	return nil
}

func (c *Client) GetCalendar(ctx context.Context, start time.Time, end time.Time, institutionCode string, reservationNum int) (CalendarResponse, error) {
	api := "/dataapi/monthly_calendar"
	req := map[string]string{
		"medical_center_cd":   institutionCode,
		"reservations_number": strconv.Itoa(reservationNum),
		"disp_type":           "2",
		"start":               start.Format(time.RFC3339),
		"end":                 end.Format(time.RFC3339),
	}
	result := CalendarResponse{}
	_, err := c.do(ctx, "POST", api, req, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) SearchInstitutions(ctx context.Context, page int) (*InstitutionData, error) {
	api := "/searchmedicalinstitutions/search_medical_institutions"
	pageStr := strconv.Itoa(page)
	var iDisplayStart string
	if page == 0 {
		iDisplayStart = "true"
	} else {
		iDisplayStart = strconv.Itoa(page * 10)
	}
	req := map[string]string{
		"sEcho":                                pageStr,
		"iColumns":                             "9",
		"sColumns":                             ",,,,,,,,",
		"iDisplayStart":                        iDisplayStart,
		"iDisplayLength":                       "10",
		"mDataProp_0":                          "toggle",
		"mDataProp_1":                          "select",
		"mDataProp_2":                          "medical_center_name",
		"mDataProp_3":                          "street_address",
		"mDataProp_4":                          "vaccine_name",
		"mDataProp_5":                          "reservation_reception",
		"mDataProp_6":                          "select_explanation",
		"mDataProp_7":                          "medical_center_cd",
		"mDataProp_8":                          "reservation_site_cd",
		"position":                             "1",
		"search":                               "0",
		"search_medical_institution_name":      "",
		"search_medical_institution_name_kana": "",
		"search_ward_name":                     "",
		"search_street_address":                "",
		"search_date_limit":                    "",
		"start_count":                          "0",
		"first_load":                           "1",
		"medical_saerch_flg":                   "1",
		"free_saerch_flg":                      "1",
		"reservations_number":                  "1",
	}
	result := &InstitutionData{}
	_, err := c.do(ctx, "POST", api, req, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) init(ctx context.Context, entrypoint string) error {
	r := c.client.R().
		SetHeader("Accept", "").
		SetHeader("Content-Type", "").
		SetContext(ctx)
	_, err := r.Get(c.baseURL + entrypoint)
	// _, err := c.do(ctx, "GET", entrypoint, nil, nil)
	if err != nil {
		return err
	}
	for i, v := range c.client.Cookies {
		log.Printf("[%d/%d]: %s\n", i+1, len(c.client.Cookies), v)
	}
	if c.csrfToken == "" {
		return c.fetchCSRF(ctx)
	}
	return nil
}

func (c *Client) fetchCSRF(ctx context.Context) error {
	r := c.client.R().
		SetHeader("Accept", "").
		SetHeader("Content-Type", "").
		SetContext(ctx)
	hres, err := r.Get(c.baseURL + "/login")
	if err != nil {
		return err
	}
	for i, v := range c.client.Cookies {
		log.Printf("[%d/%d]: %s\n", i+1, len(c.client.Cookies), v)
	}
	buf := bytes.NewBuffer(hres.Body())
	doc, err := goquery.NewDocumentFromReader(buf)
	if err != nil {
		return err
	}
	val, ok := doc.Find(`meta[name="csrf-token"]`).Attr("content")
	if !ok {
		return fmt.Errorf("unable to find the csrf token")
	}
	log.Println("set client csrf-token to ", val)
	c.csrfToken = val
	return nil
}

func (c *Client) do(ctx context.Context, method, endpoint string, request map[string]string, response interface{}) (*resty.Response, error) {
	r := c.client.R().
		SetContext(ctx)
	if response != nil {
		r.SetResult(response)
	}
	if request != nil {
		request["_token"] = c.csrfToken
		r.SetFormData(request)
	}
	do := r.Get
	switch method {
	case "GET":
	case "PUT":
		do = r.Put
	case "POST":
		do = r.Post
	case "DELETE":
		do = r.Delete
	case "PATCH":
		do = r.Patch
	default:
		panic("invalid method: " + method)
	}
	hres, err := do(c.baseURL + endpoint)
	if err != nil {
		return nil, err
	} else if hres.IsError() {
		return nil, fmt.Errorf("error response: %d: %s", hres.StatusCode(), hres.Body())
	}
	if strings.Contains(hres.Header().Get("Content-Type"), "text/html") {
		log.Println("csrf-token expired, refreshing...")
		if err := c.fetchCSRF(ctx); err != nil {
			return nil, err
		}
		hres, err = do(c.baseURL + endpoint)
		if err != nil {
			return nil, err
		} else if hres.IsError() {
			return nil, fmt.Errorf("error response: %d: %s", hres.StatusCode(), hres.Body())
		}
	}
	if response != nil {
		if v, ok := response.(errorable); ok {
			if v.GetIsError() {
				return nil, fmt.Errorf("request failed: %s", v.GetErrorMsg())
			}
		}
	}
	return hres, nil
}
