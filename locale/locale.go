package locale

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// Locale service
type Locale struct {
	client client.Client
}

func New(clt client.Client) *Locale {
	return &Locale{
		client: clt,
	}
}


// LocaleGet get the current user location based on IP. Returns an object with
// user country code, country name, continent name, continent code, ip address
// and suggested currency. You can use the locale header to get the data in a
// supported language.
// 
// ([IP Geolocation by DB-IP](https://db-ip.com))
func (srv *Locale) LocaleGet()(*models.Locale, error) {
	path := "/v1/locale"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Locale{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Locale
	parsed, ok := resp.Result.(models.Locale)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// LocaleListCodes list of all locale codes in [ISO
// 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes).
func (srv *Locale) LocaleListCodes()(*models.LocaleCodeList, error) {
	path := "/v1/locale/codes"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.LocaleCodeList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.LocaleCodeList
	parsed, ok := resp.Result.(models.LocaleCodeList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// LocaleListContinents list of all continents. You can use the locale header
// to get the data in a supported language.
func (srv *Locale) LocaleListContinents()(*models.ContinentList, error) {
	path := "/v1/locale/continents"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.ContinentList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ContinentList
	parsed, ok := resp.Result.(models.ContinentList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// LocaleListCountries list of all countries. You can use the locale header to
// get the data in a supported language.
func (srv *Locale) LocaleListCountries()(*models.CountryList, error) {
	path := "/v1/locale/countries"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.CountryList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.CountryList
	parsed, ok := resp.Result.(models.CountryList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// LocaleListCountriesEU list of all countries that are currently members of
// the EU. You can use the locale header to get the data in a supported
// language.
func (srv *Locale) LocaleListCountriesEU()(*models.CountryList, error) {
	path := "/v1/locale/countries/eu"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.CountryList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.CountryList
	parsed, ok := resp.Result.(models.CountryList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// LocaleListCountriesPhones list of all countries phone codes. You can use
// the locale header to get the data in a supported language.
func (srv *Locale) LocaleListCountriesPhones()(*models.PhoneList, error) {
	path := "/v1/locale/countries/phones"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.PhoneList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PhoneList
	parsed, ok := resp.Result.(models.PhoneList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// LocaleListCurrencies list of all currencies, including currency symbol,
// name, plural, and decimal digits for all major and minor currencies. You
// can use the locale header to get the data in a supported language.
func (srv *Locale) LocaleListCurrencies()(*models.CurrencyList, error) {
	path := "/v1/locale/currencies"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.CurrencyList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.CurrencyList
	parsed, ok := resp.Result.(models.CurrencyList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// LocaleListLanguages list of all languages classified by ISO 639-1 including
// 2-letter code, name in English, and name in the respective language.
func (srv *Locale) LocaleListLanguages()(*models.LanguageList, error) {
	path := "/v1/locale/languages"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.LanguageList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.LanguageList
	parsed, ok := resp.Result.(models.LanguageList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
