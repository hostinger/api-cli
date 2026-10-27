package site

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"

	"github.com/hostinger/api-cli/api"
)

type HostingOrder struct {
	ID             int    `json:"id"`
	Status         string `json:"status"`
	SubscriptionID string `json:"subscription_id"`
	Plan           struct {
		Name string `json:"name"`
	} `json:"plan"`
	// DisplayName comes from the billing subscription, when known.
	DisplayName string `json:"-"`
}

func (o HostingOrder) name() string {
	if o.DisplayName != "" {
		return o.DisplayName
	}
	return o.Plan.Name
}

func (o HostingOrder) Label() string {
	name := o.name()
	if name == "" {
		name = "Hosting plan"
	}
	return fmt.Sprintf("%s (order %d)", name, o.ID)
}

// HostingOrders names each order after its billing subscription, since the
// orders API only has the plan key.
func HostingOrders(ctx context.Context, c *api.RawClient) ([]HostingOrder, error) {
	var orders []HostingOrder
	for page := 1; ; page++ {
		var resp struct {
			Data []HostingOrder `json:"data"`
			Meta struct {
				Total   int `json:"total"`
				PerPage int `json:"per_page"`
			} `json:"meta"`
		}
		q := url.Values{"page": {strconv.Itoa(page)}, "per_page": {"100"}}
		if err := c.JSON(ctx, http.MethodGet, "/api/hosting/v1/orders", q, nil, &resp); err != nil {
			return nil, err
		}
		orders = append(orders, resp.Data...)
		if len(resp.Data) == 0 || len(orders) >= resp.Meta.Total || resp.Meta.PerPage == 0 {
			break
		}
	}
	var subs []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	if err := c.JSON(ctx, http.MethodGet, "/api/billing/v1/subscriptions", nil, nil, &subs); err == nil {
		names := map[string]string{}
		for _, s := range subs {
			names[s.ID] = s.Name
		}
		for i := range orders {
			orders[i].DisplayName = names[orders[i].SubscriptionID]
		}
	}
	return orders, nil
}

func PlanNames(orders []HostingOrder) map[int]string {
	names := map[int]string{}
	for _, o := range orders {
		names[o.ID] = o.name()
	}
	return names
}

func HostingPlanNames(ctx context.Context, c *api.RawClient) (map[int]string, error) {
	orders, err := HostingOrders(ctx, c)
	if err != nil {
		return nil, err
	}
	return PlanNames(orders), nil
}

// PlanLabels are a grouped picker's labels: unknown orders read "Order N",
// same-named plans get their order id, and one plan gives nil (a flat picker).
func PlanLabels(sites []Website, names map[int]string) map[int]string {
	ids := map[int]bool{}
	for _, s := range sites {
		if s.OrderID != 0 {
			ids[s.OrderID] = true
		}
	}
	if len(ids) < 2 {
		return nil
	}
	seen := map[string]int{}
	for id := range ids {
		if name := names[id]; name != "" {
			seen[name]++
		}
	}
	labels := map[int]string{}
	for id := range ids {
		name := names[id]
		switch {
		case name == "":
			name = fmt.Sprintf("Order %d", id)
		case seen[name] > 1:
			name = fmt.Sprintf("%s (order %d)", name, id)
		}
		labels[id] = name
	}
	return labels
}

func SortByPlan(sites []Website, labels map[int]string) {
	if len(labels) == 0 {
		return
	}
	sort.SliceStable(sites, func(i, j int) bool { return labels[sites[i].OrderID] < labels[sites[j].OrderID] })
}

func PlanOption(w Website, labels map[int]string) string {
	label := Label(w)
	if plan, ok := labels[w.OrderID]; ok {
		label = plan + ": " + label
	}
	return label
}
