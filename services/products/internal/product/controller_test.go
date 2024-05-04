package product

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGettingAllProducts(t *testing.T) {
	old := getAllProducts
	defer func() { getAllProducts = old }()

	getAllProducts = func() ([]Product, error) {
		return []Product{
			{
				Id:       1,
				Name:     "Test Product",
				Price:    20.9,
				ImageUrl: "./test.jpg",
			},
		}, nil
	}

	w := httptest.NewRecorder()
	e := gin.Default()
	RegisterController(e)

	r, _ := http.NewRequest("GET", "/products", nil)
	e.ServeHTTP(w, r)

	if w.Code != 200 {
		t.Fatalf("Status code is different from expected. Want %q, get %q", w.Code, 200)
	}

	found := w.Body.String()
	want := `[{"id":1,"name":"Test Product","price":20.9,"image_url":"./test.jpg"}]`
	if want != found {
		t.Fatalf("The json marshalling result is different from the expected! Desired is %q, found %q", want, found)
	}
}
