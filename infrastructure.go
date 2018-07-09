package loginradius

import (
	"os"
	"time"
)

// Configurations contains the dashboard configuration information retrieved from the Get Configurations API
type Configurations struct {
	SocialSchema struct {
		Providers []struct {
			Name     string `json:"Name"`
			Endpoint string `json:"Endpoint"`
		} `json:"Providers"`
	} `json:"SocialSchema"`
	RegistrationFormSchema []struct {
		Checked          bool        `json:"Checked"`
		Type             string      `json:"type"`
		Name             string      `json:"name"`
		Display          string      `json:"display"`
		Rules            string      `json:"rules"`
		Options          interface{} `json:"options"`
		Permission       string      `json:"permission"`
		DataSource       interface{} `json:"DataSource"`
		Parent           string      `json:"Parent"`
		ParentDataSource interface{} `json:"ParentDataSource"`
	} `json:"RegistrationFormSchema"`
	SecurityQuestions struct {
		Questions             []interface{} `json:"Questions"`
		SecurityQuestionCount int           `json:"SecurityQuestionCount"`
	} `json:"SecurityQuestions"`
	IsHTTPS                bool   `json:"IsHttps"`
	AppName                string `json:"AppName"`
	IsCustomerRegistration bool   `json:"IsCustomerRegistration"`
	APIVersion             struct {
		V1 bool `json:"v1"`
		V2 bool `json:"v2"`
	} `json:"ApiVersion"`
	EmailVerificationFlow                    string `json:"EmailVerificationFlow"`
	IsPhoneLogin                             bool   `json:"IsPhoneLogin"`
	IsDisabledSocialRegistration             bool   `json:"IsDisabledSocialRegistration"`
	IsDisabledAccountLinking                 bool   `json:"IsDisabledAccountLinking"`
	IsAgeRestriction                         bool   `json:"IsAgeRestriction"`
	IsSecurityQuestion                       bool   `json:"IsSecurityQuestion"`
	AskRequiredFieldsOnTraditionalLogin      bool   `json:"AskRequiredFieldsOnTraditionalLogin"`
	IsLogoutOnEmailVerification              bool   `json:"IsLogoutOnEmailVerification"`
	IsNoCallbackForSocialLogin               bool   `json:"IsNoCallbackForSocialLogin"`
	IsUserNameLogin                          bool   `json:"IsUserNameLogin"`
	IsMobileCallbackForSocialLogin           bool   `json:"IsMobileCallbackForSocialLogin"`
	IsInvisibleRecaptcha                     bool   `json:"IsInvisibleRecaptcha"`
	AskPasswordOnSocialLogin                 bool   `json:"AskPasswordOnSocialLogin"`
	AskEmailIDForUnverifiedUserLogin         bool   `json:"AskEmailIdForUnverifiedUserLogin"`
	AskOptionalFieldsOnSocialSignup          bool   `json:"AskOptionalFieldsOnSocialSignup"`
	IsRiskBasedAuthentication                bool   `json:"IsRiskBasedAuthentication"`
	IsV2Recaptcha                            bool   `json:"IsV2Recaptcha"`
	CheckPhoneNoAvailabilityOnRegistration   bool   `json:"CheckPhoneNoAvailabilityOnRegistration"`
	DuplicateEmailWithUniqueUsername         bool   `json:"DuplicateEmailWithUniqueUsername"`
	StoreOnlyRegistrationFormFieldsForSocial bool   `json:"StoreOnlyRegistrationFormFieldsForSocial"`
	OTPEmailVerification                     bool   `json:"OTPEmailVerification"`
	LoginLockedConfiguration                 struct {
		LoginLockedType            string      `json:"LoginLockedType"`
		MaximumFailedLoginAttempts interface{} `json:"MaximumFailedLoginAttempts"`
		SuspendConfiguration       struct {
		} `json:"SuspendConfiguration"`
	} `json:"LoginLockedConfiguration"`
	IsInstantSignin struct {
		EmailLink bool `json:"EmailLink"`
		SmsOtp    bool `json:"SmsOtp"`
	} `json:"IsInstantSignin"`
	IsLoginOnEmailVerification bool `json:"IsLoginOnEmailVerification"`
	TwoFactorAuthentication    struct {
		IsEnabled             bool `json:"IsEnabled"`
		IsRequired            bool `json:"IsRequired"`
		IsGoogleAuthenticator bool `json:"IsGoogleAuthenticator"`
	} `json:"TwoFactorAuthentication"`
	IsRememberMe               bool        `json:"IsRememberMe"`
	V2RecaptchaSiteKey         string      `json:"V2RecaptchaSiteKey"`
	QQTencentCaptchaKey        string      `json:"QQTencentCaptchaKey"`
	NoRegistration             bool        `json:"NoRegistration"`
	CustomDomain               interface{} `json:"CustomDomain"`
	PrivacyPolicyConfiguration struct {
	} `json:"PrivacyPolicyConfiguration"`
}

// ServerTime contains timing and location information that is useful for SOTT generation
type ServerTime struct {
	ServerLocation string `json:"ServerLocation"`
	ServerName     string `json:"ServerName"`
	CurrentTime    string `json:"CurrentTime"`
	Sott           struct {
		StartTime      string `json:"StartTime"`
		EndTime        string `json:"EndTime"`
		TimeDifference string `json:"TimeDifference"`
	} `json:"Sott"`
}

// SOTT is a struct that contains information on SOTTs generated by the API
type SOTT struct {
	Sott       string    `json:"Sott"`
	ExpiryTime time.Time `json:"ExpiryTime"`
}

// ActiveSession contains information on the current session received by the API
type ActiveSession struct {
	Data []struct {
		AccessToken string    `json:"AccessToken"`
		Browser     string    `json:"Browser"`
		Device      string    `json:"Device"`
		Os          string    `json:"Os"`
		DeviceType  string    `json:"DeviceType"`
		City        string    `json:"City"`
		Country     string    `json:"Country"`
		IP          string    `json:"Ip"`
		LoginDate   time.Time `json:"LoginDate"`
	} `json:"data"`
	Nextcursor int `json:"nextcursor"`
}

// GetConfiguration is used to get the configurations which are set in the
// LoginRadius Dashboard for a particular LoginRadius site/environment.
func GetConfiguration() (Configurations, error) {
	data := new(Configurations)
	req, reqErr := CreateRequest("GET", "https://config.lrcontent.com/ciam/appinfo", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetServerTime allows you to query your LoginRadius account for basic server information
// and server time information which is useful when generating an SOTT token.
func GetServerTime(timeDifference string) (ServerTime, error) {
	data := new(ServerTime)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN") + "/identity/v2/serverinfo", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("apikey", os.Getenv("APIKEY"))
	q.Add("timedifference", timeDifference)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetGenerateSottAPI allows you to generate SOTT with a given expiration time.
func GetGenerateSottAPI(timeDifference string) (SOTT, error) {
	data := new(SOTT)
	req, reqErr := CreateRequest("GET", os.Getenv("DOMAIN") + "/identity/v2/manage/account/sott", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("timedifference", timeDifference)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("X-LoginRadius-ApiKey", os.Getenv("APIKEY"))
	req.Header.Add("X-LoginRadius-ApiSecret", os.Getenv("APISECRET"))
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}

// GetActiveSessionDetails is use to get all active seesions by Access Token.
func GetActiveSessionDetails(accessToken string) (ActiveSession, error) {
	data := new(ActiveSession)
	req, reqErr := CreateRequest("GET", "http://api.loginradius.com/api/v2/access_token/activesession", "")
	if reqErr != nil {
		return *data, reqErr
	}

	q := req.URL.Query()
	q.Add("token", accessToken)
	q.Add("key", os.Getenv("APIKEY"))
	q.Add("secret", os.Getenv("APISECRET"))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")

	err := RunRequest(req, data)
	return *data, err
}