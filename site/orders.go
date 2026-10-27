package site

import (
	"cmp"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/utils"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var ErrNoOrder error = &utils.FlagError{Err: errors.New("no order given: pass the order ID")}

var orderListPaths = map[string]string{
	"hosting":        "/api/hosting/v1/orders",
	"agency-hosting": "/api/agency-hosting/v1/orders",
	"mail":           "/api/mail/v1/orders",
}

// Order is any product's order, reduced to what a prompt needs.
type Order struct {
	ID     string
	Label  string
	Status string
}

// ResolveOrder: arg, else --order-id, HOSTINGER_ORDER_ID, the link (hosting
// only), then a prompt.
func ResolveOrder(cmd *cobra.Command, args []string, extras int, group string) (orderID string, rest []string, err error) {
	if len(args)-extras >= 1 {
		return args[0], args[1:], nil
	}
	rest = args
	if cmd != nil {
		if f := contextFlag(cmd, "order-id"); f != nil && f.Changed {
			return f.Value.String(), rest, nil
		}
	}
	if v := os.Getenv("HOSTINGER_ORDER_ID"); v != "" {
		return v, rest, nil
	}
	if group == "hosting" {
		p, _, err := Find(".")
		if err != nil {
			return "", nil, linkError(err)
		}
		if p != nil && p.OrderID != 0 && (p.Kind == "" || p.Kind == "website") {
			return strconv.Itoa(p.OrderID), rest, nil
		}
	}
	if !Interactive() {
		return "", nil, ErrNoOrder
	}
	ctx := context.Background()
	if cmd != nil && cmd.Context() != nil {
		ctx = cmd.Context()
	}
	o, err := ChooseOrder(ctx, group)
	if err != nil {
		return "", nil, err
	}
	return o.ID, rest, nil
}

func ListOrders(ctx context.Context, c *api.RawClient, group string) ([]Order, error) {
	path, ok := orderListPaths[group]
	if !ok {
		return nil, fmt.Errorf("orders of %s cannot be listed; pass the order id", group)
	}
	var all []Order
	for page := 1; ; page++ {
		var resp struct {
			Data []struct {
				ID     json.RawMessage `json:"id"`
				Status string          `json:"status"`
				Domain json.RawMessage `json:"domain"` // a name, or {id, name} (mail)
				Plan   *struct {
					Name  string `json:"name"`
					Title string `json:"title"` // mail: "Free Email", where name is a key
				} `json:"plan"`
			} `json:"data"`
			Meta struct {
				Total   int `json:"total"`
				PerPage int `json:"per_page"`
			} `json:"meta"`
		}
		q := url.Values{"page": {strconv.Itoa(page)}, "per_page": {"100"}}
		if err := c.JSON(ctx, http.MethodGet, path, q, nil, &resp); err != nil {
			return nil, err
		}
		for _, row := range resp.Data {
			id := strings.Trim(string(row.ID), `"`)
			var parts []string
			if row.Plan != nil {
				if plan := cmp.Or(row.Plan.Title, row.Plan.Name); plan != "" {
					parts = append(parts, plan)
				}
			}
			if d := orderDomain(row.Domain); d != "" {
				parts = append(parts, d)
			}
			label := id
			if len(parts) > 0 {
				label = strings.Join(parts, " · ") + " (order " + id + ")"
			}
			if row.Status != "" && row.Status != "active" {
				label += " [" + row.Status + "]"
			}
			all = append(all, Order{ID: id, Label: label, Status: row.Status})
		}
		if len(resp.Data) == 0 || len(all) >= resp.Meta.Total || resp.Meta.PerPage == 0 {
			return all, nil
		}
	}
}

// orderDomain reads an order's domain, given as a name or as {id, name}.
func orderDomain(raw json.RawMessage) string {
	var name string
	if json.Unmarshal(raw, &name) == nil {
		return name
	}
	var obj struct {
		Name string `json:"name"`
	}
	_ = json.Unmarshal(raw, &obj)
	return obj.Name
}

// ChooseOrder returns the only usable order without asking.
func ChooseOrder(ctx context.Context, group string) (*Order, error) {
	orders, err := ListOrders(ctx, Client(), group)
	if err != nil {
		return nil, err
	}
	usable := orders[:0]
	for _, o := range orders {
		switch strings.ToLower(o.Status) {
		case "deleted", "deleting", "cancelled":
			continue
		}
		usable = append(usable, o)
	}
	if len(usable) == 0 {
		return nil, fmt.Errorf("this account has no %s order", group)
	}
	if len(usable) == 1 {
		return &usable[0], nil
	}
	options := make([]string, len(usable))
	for i, o := range usable {
		options[i] = o.Label
	}
	picked, err := pick(pterm.DefaultInteractiveSelect.WithOptions(options).WithFilter(len(options) > 8).WithMaxHeight(SelectHeight(len(options))), "Which "+group+" order?")
	if err != nil {
		return nil, err
	}
	for i, o := range options {
		if o == picked {
			return &usable[i], nil
		}
	}
	return nil, errors.New("no order selected")
}
