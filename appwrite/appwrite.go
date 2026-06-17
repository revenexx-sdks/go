package appwrite

import (
	"time"

	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/apps"
	"github.com/revenexx-sdks/go/avatars"
	"github.com/revenexx-sdks/go/carts"
	"github.com/revenexx-sdks/go/channels"
	"github.com/revenexx-sdks/go/customers"
	"github.com/revenexx-sdks/go/greetings"
	"github.com/revenexx-sdks/go/inventories"
	"github.com/revenexx-sdks/go/locale"
	"github.com/revenexx-sdks/go/markets"
	"github.com/revenexx-sdks/go/messaging"
	"github.com/revenexx-sdks/go/orders"
	"github.com/revenexx-sdks/go/pages"
	"github.com/revenexx-sdks/go/payments"
	"github.com/revenexx-sdks/go/prices"
	"github.com/revenexx-sdks/go/products"
	"github.com/revenexx-sdks/go/search"
	"github.com/revenexx-sdks/go/shipping"
	"github.com/revenexx-sdks/go/sites"
	"github.com/revenexx-sdks/go/storage"
	"github.com/revenexx-sdks/go/tokens"
)

func NewApps(clt client.Client) *apps.Apps {
	return apps.New(clt)
}
func NewAvatars(clt client.Client) *avatars.Avatars {
	return avatars.New(clt)
}
func NewCarts(clt client.Client) *carts.Carts {
	return carts.New(clt)
}
func NewChannels(clt client.Client) *channels.Channels {
	return channels.New(clt)
}
func NewCustomers(clt client.Client) *customers.Customers {
	return customers.New(clt)
}
func NewGreetings(clt client.Client) *greetings.Greetings {
	return greetings.New(clt)
}
func NewInventories(clt client.Client) *inventories.Inventories {
	return inventories.New(clt)
}
func NewLocale(clt client.Client) *locale.Locale {
	return locale.New(clt)
}
func NewMarkets(clt client.Client) *markets.Markets {
	return markets.New(clt)
}
func NewMessaging(clt client.Client) *messaging.Messaging {
	return messaging.New(clt)
}
func NewOrders(clt client.Client) *orders.Orders {
	return orders.New(clt)
}
func NewPages(clt client.Client) *pages.Pages {
	return pages.New(clt)
}
func NewPayments(clt client.Client) *payments.Payments {
	return payments.New(clt)
}
func NewPrices(clt client.Client) *prices.Prices {
	return prices.New(clt)
}
func NewProducts(clt client.Client) *products.Products {
	return products.New(clt)
}
func NewSearch(clt client.Client) *search.Search {
	return search.New(clt)
}
func NewShipping(clt client.Client) *shipping.Shipping {
	return shipping.New(clt)
}
func NewSites(clt client.Client) *sites.Sites {
	return sites.New(clt)
}
func NewStorage(clt client.Client) *storage.Storage {
	return storage.New(clt)
}
func NewTokens(clt client.Client) *tokens.Tokens {
	return tokens.New(clt)
}

// NewClient initializes a new RevenexxAPIRevenexx client with a given timeout
func NewClient(optionalSetters ...client.ClientOption) client.Client {
	return client.New(optionalSetters...)
}

// Helper method to construct NewClient()
func WithEndpoint(endpoint string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Endpoint = endpoint
		return nil
	}
}

// Helper method to construct NewClient()
func WithTimeout(timeout time.Duration) client.ClientOption {
	return func(clt *client.Client) error {
		httpClient, err := client.GetDefaultClient(timeout)
		if err != nil {
			return err
		}

		clt.Timeout = timeout
		clt.Client = httpClient

		return nil
	}
}

// Helper method to construct NewClient()
func WithSelfSigned(status bool) client.ClientOption {
	return func(clt *client.Client) error {
		clt.SelfSigned = status
		return nil
	}
}

// Helper method to construct NewClient()
func WithChunkSize(size int64) client.ClientOption {
	return func(clt *client.Client) error {
		clt.ChunkSize = size
		return nil
	}
}

// WithTenant sets the X-Revenexx-Tenant header sent on every request,
// scoping calls to the given tenant.
func WithTenant(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Headers["X-Revenexx-Tenant"] = value
		return nil
	}
}

// Helper method to construct NewClient()
// 
// A gateway-managed scoped API key (rvxk_…).
func WithApiKeyAuth(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Headers["X-Revenexx-Api-Key"] = value
		return nil
	}
}
// Helper method to construct NewClient()
// 
// A Zitadel-issued JWT (Cockpit / interactive callers).
func WithBearerAuth(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Headers["Authorization"] = value
		return nil
	}
}
