package customers_roles

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// CustomersRoles service
type CustomersRoles struct {
	client client.Client
}

func New(clt client.Client) *CustomersRoles {
	return &CustomersRoles{
		client: clt,
	}
}


// CustomersRolesList the whole catalogue in one read: every role a contact of
// this tenant can hold, the permissions each one grants, and the built-in
// permission vocabulary those grants are drawn from. Roles are held by a
// CONTACT and apply inside that contact's organization; there is no global
// customer role. Permissions are derived from the role at read time and never
// stored per contact, so a role change takes effect immediately and cannot
// leave a stale grant. The role to permission MAPPING is per tenant and
// configurable (PUT /customers/roles/{key}/permissions); a tenant that has
// not configured anything gets the built-ins and 'source' says which of the
// two answered. Built-in roles, least to most privileged: viewer (Viewer),
// requester (Requester), buyer (Buyer), approver (Approver), admin
// (Administrator). The permission KEYS themselves come from the cross-app
// ledger — every installed app declares what it enforces — so a tenant
// may grant a key this list does not mention.
func (srv *CustomersRoles) CustomersRolesList()(*models.RoleCatalogResponse, error) {
	path := "/v1/customers/roles"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.RoleCatalogResponse{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.RoleCatalogResponse
	parsed, ok := resp.Result.(models.RoleCatalogResponse)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// CustomersRolesDefaults idempotent: a role that already exists is left
// completely alone, its permissions included, so re-seeding never undoes a
// merchant's edits. Creates viewer, requester, buyer, approver, admin with
// the built-in mapping. A tenant that never calls this still behaves
// correctly — the catalogue and every permission read fall back to the same
// built-ins.
func (srv *CustomersRoles) CustomersRolesDefaults(Data interface{})(*models.Error, error) {
	path := "/v1/customers/roles/defaults"
	params := map[string]interface{}{}
	params["data"] = Data
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
			
// CustomersRolesPermissionsReplace the whole new set in one call — the
// shape a role editor actually produces, and the one that cannot leave a
// half-applied grant behind if a second call fails. Seeds the built-in roles
// first when the tenant has none, so editing works without calling /defaults.
// Permission keys are free text on purpose: they belong to whichever app
// declared them, and a grant for an app that is not installed simply has
// nothing to act on.
func (srv *CustomersRoles) CustomersRolesPermissionsReplace(Key string, Permissions []string)(*models.Error, error) {
	r := strings.NewReplacer("{key}", Key)
	path := r.Replace("/v1/customers/roles/{key}/permissions")
	params := map[string]interface{}{}
	params["key"] = Key
	params["permissions"] = Permissions
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
