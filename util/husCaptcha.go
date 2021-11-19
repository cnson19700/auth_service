//nolint:wrapcheck
package util

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// Google recaptcha response.
type recaptchaResponse struct {
	Success     bool      `json:"success"`
	Score       float32   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"` //nolint:tagliatelle
}

// Recaptcha api endpoint.
const (
	requestTimeout  = time.Second * 10
	recaptchaServer = "https://www.google.com/recaptcha/api/siteverify"
)

// Main object.
type HusCaptcha struct {
	PrivateKey string
}

// Check : initiate a recaptcha verify request.
func (r *HusCaptcha) requestVerify(remoteAddr string, captchaResponse string) (recaptchaResponse, error) {
	// Fire off request with a timeout of 10 seconds
	httpClient := http.Client{Timeout: requestTimeout}
	data := url.Values{
		"secret":   {r.PrivateKey},
		"remoteip": {remoteAddr},
		"response": {captchaResponse},
	}
	req, _ := http.NewRequestWithContext(context.Background(),
		"POST",
		recaptchaServer,
		strings.NewReader(data.Encode()))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := httpClient.Do(req)
	// Request failed
	if err != nil {
		return recaptchaResponse{Success: false}, err
	}

	// Close response when function exits
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return recaptchaResponse{Success: false}, errors.Wrap(err, "fail to read the response")
	}

	// Parse json to our response object
	var response recaptchaResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return recaptchaResponse{Success: false}, err
	}

	// Return our object response
	return response, nil
}

// Check : check user IP, captcha subject (= page) and captcha response but return treshold.
func (r *HusCaptcha) Check(remoteip string, action string, response string) (success bool, score float32, err error) {
	resp, err := r.requestVerify(remoteip, response)
	// fetch/parsing failed
	if err != nil {
		return false, 0, err
	}

	// Captcha subject did not match
	if !strings.EqualFold(resp.Action, action) {
		err := errors.New("reCaptcha actions do not match")

		return false, 0, err
	}

	// recaptcha token was not valid
	if !resp.Success {
		return false, 0, nil
	}

	// user treshold was not enough
	return true, resp.Score, nil
}

// Verify : check user IP, captcha subject (= page) and captcha response.
func (r *HusCaptcha) Verify(
	remoteIP string,
	action string,
	response string,
	minScore float32,
) (success bool, err error) {
	success, score, err := r.Check(remoteIP, action, response)

	// return false if response failed
	if !success || err != nil {
		return false, err
	}

	// user score was not enough
	return score >= minScore, nil
}
