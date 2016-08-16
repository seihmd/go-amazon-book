package amazonbook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

func mapResponse(res []byte) (*Response, error) {
	r := Response{}
	err := xml.Unmarshal(res, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (api AmazonProductAPI) genSignAndFetch(Operation string, Parameters map[string]string) (*Response, error) {
	genURL, err := generateAmazonURL(api, Operation, Parameters)
	if err != nil {
		return nil, err
	}
	setTimestamp(genURL)

	signedurl, err := signAmazonURL(genURL, api)
	if err != nil {
		return nil, err
	}

	if api.Client == nil {
		api.Client = http.DefaultClient
	}

	resp, err := api.Client.Get(signedurl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return mapResponse(body)
}

func generateAmazonURL(api AmazonProductAPI, Operation string, Parameters map[string]string) (finalURL *url.URL, err error) {

	result, err := url.Parse(api.Host)
	if err != nil {
		return nil, err
	}

	result.Host = api.Host
	result.Scheme = "http"
	result.Path = "/onca/xml"

	values := url.Values{}
	values.Add("Operation", Operation)
	values.Add("Service", "AWSECommerceService")
	values.Add("AWSAccessKeyId", api.ID)
	values.Add("Version", "2013-08-01")
	values.Add("AssociateTag", api.AssociateTag)

	for k, v := range Parameters {
		values.Set(k, v)
	}

	params := values.Encode()
	result.RawQuery = params

	return result, nil
}

func setTimestamp(origURL *url.URL) (err error) {
	values, err := url.ParseQuery(origURL.RawQuery)
	if err != nil {
		return err
	}
	values.Set("Timestamp", time.Now().UTC().Format(time.RFC3339))
	origURL.RawQuery = values.Encode()
	return nil
}

func signAmazonURL(origURL *url.URL, api AmazonProductAPI) (signedURL string, err error) {

	escapeURL := strings.Replace(origURL.RawQuery, ",", "%2C", -1)
	escapeURL = strings.Replace(escapeURL, ":", "%3A", -1)
	params := strings.Split(escapeURL, "&")
	sort.Strings(params)
	sortedParams := strings.Join(params, "&")
	toSign := fmt.Sprintf("GET\n%s\n%s\n%s", origURL.Host, origURL.Path, sortedParams)

	hasher := hmac.New(sha256.New, []byte(api.Secret))
	_, err = hasher.Write([]byte(toSign))
	if err != nil {
		return "", err
	}

	hash := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	hash = url.QueryEscape(hash)
	newParams := fmt.Sprintf("%s&Signature=%s", sortedParams, hash)
	origURL.RawQuery = newParams

	return origURL.String(), nil
}
