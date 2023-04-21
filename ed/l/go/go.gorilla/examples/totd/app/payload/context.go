package payload

type ContextKey string

const (
	// ContextKeyToken represents context key for token header.
	ContextKeyToken ContextKey = "X-Token"

	// ContextKeyRetailer represents context key for retailer (client) header.
	ContextKeyRetailer ContextKey = "X-Retailer-Id"

	// ContextKeyEnv represents context key for env header.
	ContextKeyEnv ContextKey = "X-Env-Type"

	// ContextKeyMfc represents context key for mfc header.
	ContextKeyMfc ContextKey = "X-Mfc"
)
