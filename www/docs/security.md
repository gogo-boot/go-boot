## Security
`go-boot` support `oauth2` and `oidc`.
`go-boot` is using casbin to control the user access. 
if a user doesn't own token, the http request will be rejected.

### Casbin Access Control Model 
`platform/config/authz_model.conf` file defines access control model.
It uses RBAC(Role-Based Access Control)model as a default.

### Casbin Policy
`platform/config/authz_policy.csv` file defines casbin policy. 
currently, all users can access the all the http endpoints
