package tokyovacapi

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

var (
	ErrTokenExpired           = errors.New("token expired")
	ErrTokenInvalid           = errors.New("token invalid")
	ErrReservationUnavailable = errors.New("reservation unavailable")
)

type Client struct {
	client  *resty.Client
	baseURL string // https://api.vaccines.sciseed.jp
}

func NewClient() (*Client, error) {
	baseURL := "https://api.vaccines.sciseed.jp"
	certpool, err := x509.SystemCertPool()
	if err != nil {
		return nil, fmt.Errorf("failed to read system cert pool: %w", err)
	}
	baseURL = strings.TrimSuffix(baseURL, "/")
	client := resty.New().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetTLSClientConfig(&tls.Config{
			RootCAs:    certpool,
			MinVersion: tls.VersionTLS12,
		})

	return &Client{client: client, baseURL: baseURL}, nil
}

func (c *Client) Login(ctx context.Context, partition Partition, rangeKey string, password string) (*LoginResponse, error) {
	endpoint := "/public/{partition}/login/"
	pathParams := map[string]string{
		"partition": string(partition),
	}
	req := &LoginRequest{
		PartitionKey: partition,
		RangeKey:     rangeKey,
		Password:     password,
	}
	result := &LoginResponse{}
	queryParams := trimMap(map[string]string{})
	_, err := c.do(ctx, "POST", endpoint, pathParams, queryParams, req, result)
	return result, err
}

func (c *Client) GetAvailableDepartments(ctx context.Context, partition Partition) (*AvailableDepartmentResponse, error) {
	endpoint := "/public/{partition}/available_department/"
	pathParams := map[string]string{
		"partition": string(partition),
	}
	result := &AvailableDepartmentResponse{}
	queryParams := trimMap(map[string]string{})
	_, err := c.do(ctx, "GET", endpoint, pathParams, queryParams, nil, result)
	return result, err
}

func (c *Client) GetDepartments(ctx context.Context, partition Partition) (*GetDepartmentsResponse, error) {
	endpoint := "/public/{partition}/department/"
	pathParams := map[string]string{
		"partition": string(partition),
	}
	result := &GetDepartmentsResponse{}
	queryParams := trimMap(map[string]string{})
	_, err := c.do(ctx, "GET", endpoint, pathParams, queryParams, nil, result)
	return result, err
}

func (c *Client) GetAvailableFrames(ctx context.Context, partition Partition, start time.Time, end time.Time, item VacItem) (*AvailableDepartmentResponse, error) {
	endpoint := "/public/{partition}/reservation_frame/"
	pathParams := map[string]string{
		"partition": string(partition),
	}
	result := &AvailableDepartmentResponse{}
	queryParams := trimMap(map[string]string{
		"item_id":           strconv.Itoa(int(item)),
		"start_date_after":  start.Format("2006-01-02"),
		"start_date_before": end.Format("2006-01-02"),
	})
	_, err := c.do(ctx, "GET", endpoint, pathParams, queryParams, nil, result)
	return result, err
}

func (c *Client) GetPerson(ctx context.Context, partition Partition, token string) (*GetPersonResponse, error) {
	if err := verifyJWT(token); err != nil {
		return nil, err
	}
	endpoint := "/public/{partition}/person/"
	pathParams := map[string]string{
		"partition": string(partition),
	}
	result := &GetPersonResponse{}
	_, err := c.do(ctx, "GET", endpoint, pathParams, nil, nil, result, func(r *resty.Request) { r.SetAuthScheme("Bearer").SetAuthToken(token) })
	return result, err
}

func (c *Client) GetArticles(ctx context.Context, partition Partition) (*GetArticlesResponse, error) {
	endpoint := "/public/{partition}/articles/"
	pathParams := map[string]string{
		"partition": string(partition),
	}
	result := &GetArticlesResponse{}
	_, err := c.do(ctx, "GET", endpoint, pathParams, nil, nil, result)
	return result, err
}

func (c *Client) Reserve(ctx context.Context, partition Partition, frameID int, token string) (*ReserveResponse, error) {
	if err := verifyJWT(token); err != nil {
		return nil, err
	}
	endpoint := "/public/{partition}/reservation/"
	pathParams := map[string]string{
		"partition": string(partition),
	}
	req := &ReserveRequest{
		ReservationFrameID: frameID,
	}
	result := &ReserveResponse{}
	_, err := c.do(ctx, "POST", endpoint, pathParams, nil, req, result, func(r *resty.Request) { r.SetAuthScheme("Bearer").SetAuthToken(token) })
	return result, err
}

func (c *Client) do(ctx context.Context, method, endpoint string, pathParams, queryParams map[string]string, request interface{}, response interface{}, opts ...func(*resty.Request)) (*resty.Response, error) {
	errResp := &Error{}
	r := c.client.R().
		SetPathParams(pathParams).
		SetQueryParams(queryParams).
		SetContext(ctx).
		SetError(errResp)
	for _, opt := range opts {
		opt(r)
	}
	if response != nil {
		r.SetResult(response)
	}
	if request != nil {
		r.SetBody(request)
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
		if err, ok := toError(errResp); ok {
			return nil, err
		}
		return nil, fmt.Errorf("error response: %d: %s", hres.StatusCode(), hres.Body())
	}
	return hres, nil
}
