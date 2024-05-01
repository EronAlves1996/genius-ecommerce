package product_test

import (
	"encoding/json"
	"testing"

	"github.com/EronAlves1996/services/products/internal/product"
)

func TestSerializedFormat(t *testing.T) {
	ps := product.FetchAllProducts()[0]
	mps, err := json.Marshal(ps)

	if err != nil {
		t.Fatal(err)
	}

	found := string(mps)
	want := `{"id":1,"name":"Test Product","price":20.9,"image_url":"./test.jpg"}`
	if want != found {
		t.Fatalf("The json marshalling result is different from the expected! Desired is %q, found %q", want, found)
	}
}
