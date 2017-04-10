package woocommerce

import (
	"encoding/json"
	"io"
	"testing"
)

var (
	store      = "http://"
	ck         = ""
	cs         = ""
	orderCount = 0
)

/*
func TestClient(t *testing.T) {
	if store == "" || ck == "" || cs == "" {
		t.Fatal("Information needed")
	}
	client, err := NewClient(store, ck, cs, nil)
	if err != nil {
		t.Fatal(err)
	}
	body, err := client.Get("orders", nil)
	if err != nil {
		t.Fatal(err)
	}
	jsonDecoder := json.NewDecoder(body)
	var data map[string]interface{}
	if err := jsonDecoder.Decode(&data); err != nil && err != io.EOF {
		t.Fatal(err)
	}
	defer body.Close()
	orders, ok := data["orders"]
	if !ok {
		t.Fatal("Wrong return format")
	}
	if orderCount != len(orders.([]interface{})) {
		t.Fatal("Wrong count of orders")
	}
}
*/
