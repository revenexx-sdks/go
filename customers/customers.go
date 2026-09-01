package customers

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// Customers service
type Customers struct {
	client client.Client
}

func New(clt client.Client) *Customers {
	return &Customers{
		client: clt,
	}
}

			
// CustomersAuthLogin an email and a password go in; a session and the CONTACT
// behind it come back, so a storefront knows in one call both that the buyer
// is signed in and who they are. The session is minted server-side rather
// than handed back from the credential check, because the account route hides
// the session secret from non-privileged responses and a trusted BFF needs
// it. `permissions` carries the buyer's effective grants, so a BFF does not
// need a second call to decide what to render.
func (srv *Customers) CustomersAuthLogin(Email string, Password string)(*models.Error, error) {
	path := "/v1/customers/auth/login"
	params := map[string]interface{}{}
	params["email"] = Email
	params["password"] = Password
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// CustomersAuthLogout ends ONE session — the buyer signs out on this device
// and stays signed in on the others, because the session id is what is
// revoked and not the account. The contact row is untouched: signing out is
// not blocking, and a caller wanting the second thing wants `status:
// "blocked"` on the contact instead. Both ids come from what
// `/customers/auth/login` answered, and a BFF should drop its own cookie
// whatever this answers — the session is unusable afterwards either way.
func (srv *Customers) CustomersAuthLogout(SessionId string, UserId string)(*models.Error, error) {
	path := "/v1/customers/auth/logout"
	params := map[string]interface{}{}
	params["session_id"] = SessionId
	params["user_id"] = UserId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// CustomersAuthMagicLink sign in without a password: a link goes to the
// address, and `PUT /customers/auth/magic-link` turns it into a session.
// Creates the account when the address is new, which makes this a
// registration path as much as a sign-in one — and why an address nobody
// holds is not distinguished in the answer. The mail is this shop's own
// template through the messaging service; the secret is not in this response,
// only in the link.
func (srv *Customers) CustomersAuthMagicLink(Email string, Url string)(*models.Error, error) {
	path := "/v1/customers/auth/magic-link"
	params := map[string]interface{}{}
	params["email"] = Email
	params["url"] = Url
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// CustomersAuthMagicLinkConfirm the buyer clicked the link and the storefront
// read `userId` and `secret` out of it. Answers exactly what a password login
// answers — session, contact and effective grants — because a shop must
// not have to branch on how somebody signed in.
func (srv *Customers) CustomersAuthMagicLinkConfirm(Secret string, UserId string)(*models.Error, error) {
	path := "/v1/customers/auth/magic-link"
	params := map[string]interface{}{}
	params["secret"] = Secret
	params["user_id"] = UserId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CustomersAuthMeOptions struct {
	SessionId string
	enabledSetters map[string]bool
}
func (options CustomersAuthMeOptions) New() *CustomersAuthMeOptions {
	options.enabledSetters = map[string]bool{
		"SessionId": false,
	}
	return &options
}
type CustomersAuthMeOption func(*CustomersAuthMeOptions)
func (srv *Customers) WithCustomersAuthMeSessionId(v string) CustomersAuthMeOption {
	return func(o *CustomersAuthMeOptions) {
		o.SessionId = v
		o.enabledSetters["SessionId"] = true
	}
}
			
// CustomersAuthMe the platform user, the customer record mirrored against it
// and the effective grants, in one call. The expected caller is a trusted
// storefront BFF holding the session on the buyer's behalf, which is why the
// ids travel in the body rather than in a browser-facing header. The grants
// are derived here on every call rather than returned from anywhere they
// could be cached, so a role changed a second ago is already reflected.
func (srv *Customers) CustomersAuthMe(UserId string, optionalSetters ...CustomersAuthMeOption)(*models.Error, error) {
	path := "/v1/customers/auth/me"
	options := CustomersAuthMeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["user_id"] = UserId
	if options.enabledSetters["SessionId"] {
		params["session_id"] = options.SessionId
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CustomersAuthMfaChallengeOptions struct {
	Factor string
	enabledSetters map[string]bool
}
func (options CustomersAuthMfaChallengeOptions) New() *CustomersAuthMfaChallengeOptions {
	options.enabledSetters = map[string]bool{
		"Factor": false,
	}
	return &options
}
type CustomersAuthMfaChallengeOption func(*CustomersAuthMfaChallengeOptions)
func (srv *Customers) WithCustomersAuthMfaChallengeFactor(v string) CustomersAuthMfaChallengeOption {
	return func(o *CustomersAuthMfaChallengeOptions) {
		o.Factor = v
		o.enabledSetters["Factor"] = true
	}
}
			
// CustomersAuthMfaChallenge between the password and the finished session:
// the buyer has proved one thing and is asked for another. Created by user
// id, because the account route that creates challenges hides the code from
// whoever may call it — and answered with the half-finished session the
// sign-in is in the middle of, through `PUT /customers/auth/mfa/challenge`.
// Needs a platform build that returns the challenge code; without one there
// is no way to read what to send, and the call answers 502 rather than
// mailing an empty challenge.
func (srv *Customers) CustomersAuthMfaChallenge(UserId string, optionalSetters ...CustomersAuthMfaChallengeOption)(*models.Error, error) {
	path := "/v1/customers/auth/mfa/challenge"
	options := CustomersAuthMfaChallengeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["user_id"] = UserId
	if options.enabledSetters["Factor"] {
		params["factor"] = options.Factor
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CustomersAuthMfaChallengeConfirmOptions struct {
	UserId string
	enabledSetters map[string]bool
}
func (options CustomersAuthMfaChallengeConfirmOptions) New() *CustomersAuthMfaChallengeConfirmOptions {
	options.enabledSetters = map[string]bool{
		"UserId": false,
	}
	return &options
}
type CustomersAuthMfaChallengeConfirmOption func(*CustomersAuthMfaChallengeConfirmOptions)
func (srv *Customers) WithCustomersAuthMfaChallengeConfirmUserId(v string) CustomersAuthMfaChallengeConfirmOption {
	return func(o *CustomersAuthMfaChallengeConfirmOptions) {
		o.UserId = v
		o.enabledSetters["UserId"] = true
	}
}
							
// CustomersAuthMfaChallengeConfirm the code the buyer typed, against the
// challenge it was sent for. The session becomes fully authenticated when
// this answers.
func (srv *Customers) CustomersAuthMfaChallengeConfirm(ChallengeId string, Code string, SessionSecret string, optionalSetters ...CustomersAuthMfaChallengeConfirmOption)(*models.Error, error) {
	path := "/v1/customers/auth/mfa/challenge"
	options := CustomersAuthMfaChallengeConfirmOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["challenge_id"] = ChallengeId
	params["code"] = Code
	params["session_secret"] = SessionSecret
	if options.enabledSetters["UserId"] {
		params["user_id"] = options.UserId
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// CustomersAuthOtp the same token as the sign-in link, delivered as a short
// code instead — for a buyer on a phone, where leaving for a mail client
// and coming back loses the checkout they were in the middle of. Redeemed
// with `PUT /customers/auth/otp`.
func (srv *Customers) CustomersAuthOtp(Email string)(*models.Error, error) {
	path := "/v1/customers/auth/otp"
	params := map[string]interface{}{}
	params["email"] = Email
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// CustomersAuthOtpConfirm the code the buyer typed, plus the `userId` the
// send answered with. Answers exactly what a password login answers —
// session, contact and effective grants — so a storefront never has to
// branch on how somebody signed in. The code is spent on first use and
// expires, so a second attempt with the same one is a 401 rather than a
// second session.
func (srv *Customers) CustomersAuthOtpConfirm(Secret string, UserId string)(*models.Error, error) {
	path := "/v1/customers/auth/otp"
	params := map[string]interface{}{}
	params["secret"] = Secret
	params["user_id"] = UserId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// CustomersAuthRecovery step one of two: a link goes to the address given,
// and `PUT /customers/auth/recovery` is what the buyer's browser comes back
// to. The identity service mints the token; the MAIL is this shop's own —
// the tenant's template, layout, language and sending domain, through the
// messaging service. The secret is NOT in this answer: it exists only inside
// the mailed link, which is the whole point of the two-step shape, and
// echoing it here would make the mail decorative. Nothing about the contact
// changes; the password only moves in step two.
func (srv *Customers) CustomersAuthRecovery(Email string, Url string)(*models.Error, error) {
	path := "/v1/customers/auth/recovery"
	params := map[string]interface{}{}
	params["email"] = Email
	params["url"] = Url
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
					
// CustomersAuthRecoveryConfirm step two: the `userId` and `secret` the mailed
// link carried, plus the password the buyer just typed. The secret is spent
// on first use and expires, so a link cannot be replayed and a second attempt
// with the same one is a 401 rather than a second password change. The new
// password is in effect the moment this answers; what happens to sessions
// opened with the old one is the identity service's policy, not this app's.
func (srv *Customers) CustomersAuthRecoveryConfirm(Password string, Secret string, UserId string)(*models.Error, error) {
	path := "/v1/customers/auth/recovery"
	params := map[string]interface{}{}
	params["password"] = Password
	params["secret"] = Secret
	params["user_id"] = UserId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type CustomersAuthRegisterOptions struct {
	FirstName string
	LastName string
	Locale string
	OrganizationId string
	OrganizationName string
	Url string
	VatId string
	VerificationUrl string
	enabledSetters map[string]bool
}
func (options CustomersAuthRegisterOptions) New() *CustomersAuthRegisterOptions {
	options.enabledSetters = map[string]bool{
		"FirstName": false,
		"LastName": false,
		"Locale": false,
		"OrganizationId": false,
		"OrganizationName": false,
		"Url": false,
		"VatId": false,
		"VerificationUrl": false,
	}
	return &options
}
type CustomersAuthRegisterOption func(*CustomersAuthRegisterOptions)
func (srv *Customers) WithCustomersAuthRegisterFirstName(v string) CustomersAuthRegisterOption {
	return func(o *CustomersAuthRegisterOptions) {
		o.FirstName = v
		o.enabledSetters["FirstName"] = true
	}
}
func (srv *Customers) WithCustomersAuthRegisterLastName(v string) CustomersAuthRegisterOption {
	return func(o *CustomersAuthRegisterOptions) {
		o.LastName = v
		o.enabledSetters["LastName"] = true
	}
}
func (srv *Customers) WithCustomersAuthRegisterLocale(v string) CustomersAuthRegisterOption {
	return func(o *CustomersAuthRegisterOptions) {
		o.Locale = v
		o.enabledSetters["Locale"] = true
	}
}
func (srv *Customers) WithCustomersAuthRegisterOrganizationId(v string) CustomersAuthRegisterOption {
	return func(o *CustomersAuthRegisterOptions) {
		o.OrganizationId = v
		o.enabledSetters["OrganizationId"] = true
	}
}
func (srv *Customers) WithCustomersAuthRegisterOrganizationName(v string) CustomersAuthRegisterOption {
	return func(o *CustomersAuthRegisterOptions) {
		o.OrganizationName = v
		o.enabledSetters["OrganizationName"] = true
	}
}
func (srv *Customers) WithCustomersAuthRegisterUrl(v string) CustomersAuthRegisterOption {
	return func(o *CustomersAuthRegisterOptions) {
		o.Url = v
		o.enabledSetters["Url"] = true
	}
}
func (srv *Customers) WithCustomersAuthRegisterVatId(v string) CustomersAuthRegisterOption {
	return func(o *CustomersAuthRegisterOptions) {
		o.VatId = v
		o.enabledSetters["VatId"] = true
	}
}
func (srv *Customers) WithCustomersAuthRegisterVerificationUrl(v string) CustomersAuthRegisterOption {
	return func(o *CustomersAuthRegisterOptions) {
		o.VerificationUrl = v
		o.enabledSetters["VerificationUrl"] = true
	}
}
					
// CustomersAuthRegister one call writes the whole buyer: the contact this app
// is the system of record for, and the platform user behind its login. When
// the body names a company it also FOUNDS one — an organization, mirrored
// into platform auth as a team, with this contact as its admin. The tenant
// setting registration_mode decides what a registration IS. 'open' (the
// default, unchanged behaviour) creates a finished account:
// registration_status='approved', status='active', login works.
// 'approval_required' creates an APPLICATION: registration_status='pending',
// status='invited', the platform user exists with the applicant's own
// password but is DISABLED, and a newly founded organization is parked as
// 'blocked' — check `approval_required` in the response and show a 'we will
// get back to you' screen instead of logging the buyer in. The registration
// gates below are all evaluated BEFORE anything is written, and a failure
// after that point rolls the organization and the contact back together.
func (srv *Customers) CustomersAuthRegister(Email string, Password string, optionalSetters ...CustomersAuthRegisterOption)(*models.Error, error) {
	path := "/v1/customers/auth/register"
	options := CustomersAuthRegisterOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["email"] = Email
	params["password"] = Password
	if options.enabledSetters["FirstName"] {
		params["first_name"] = options.FirstName
	}
	if options.enabledSetters["LastName"] {
		params["last_name"] = options.LastName
	}
	if options.enabledSetters["Locale"] {
		params["locale"] = options.Locale
	}
	if options.enabledSetters["OrganizationId"] {
		params["organization_id"] = options.OrganizationId
	}
	if options.enabledSetters["OrganizationName"] {
		params["organization_name"] = options.OrganizationName
	}
	if options.enabledSetters["Url"] {
		params["url"] = options.Url
	}
	if options.enabledSetters["VatId"] {
		params["vat_id"] = options.VatId
	}
	if options.enabledSetters["VerificationUrl"] {
		params["verification_url"] = options.VerificationUrl
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// CustomersAuthVerification confirm that the address belongs to the buyer.
// Needs no session: the verification is created through the identity
// service's users surface, because its account counterpart reads the
// authenticated user and a caller authenticating AS the user cannot see the
// secret it just created. The buyer still confirms with their own session,
// through `PUT /customers/auth/verification` — only the creation moved.
// Send it right after a registration, or from an account page.
func (srv *Customers) CustomersAuthVerification(Url string, UserId string)(*models.Error, error) {
	path := "/v1/customers/auth/verification"
	params := map[string]interface{}{}
	params["url"] = Url
	params["user_id"] = UserId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// CustomersAuthVerificationConfirm the `userId` and `secret` the mailed link
// carried. The address counts as confirmed the moment this answers; the
// secret is spent, so the link cannot be replayed.
func (srv *Customers) CustomersAuthVerificationConfirm(Secret string, UserId string)(*models.Error, error) {
	path := "/v1/customers/auth/verification"
	params := map[string]interface{}{}
	params["secret"] = Secret
	params["user_id"] = UserId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// CustomersPrincipalResolve the capability the API gateway calls to turn a
// caller's X-Revenexx-Principal assertion into the permission set it forwards
// to every other app as X-Revenexx-Permissions. This app is the platform's
// role provider (manifest#provides_roles), and this is the hot path of every
// attributed storefront request — one contact read plus the tenant's role
// map. A blocked or pending contact always resolves with active=false; what
// its `permissions` then say is the tenant's blocked_contact_behavior setting
// — 'keep' (the default, the role's grants), 'catalog_only' or 'deny_all'.
func (srv *Customers) CustomersPrincipalResolve(ContactId string)(*models.Error, error) {
	path := "/v1/customers/principal/resolve"
	params := map[string]interface{}{}
	params["contact_id"] = ContactId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
