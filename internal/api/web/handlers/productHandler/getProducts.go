package productHandler

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/kourai55k/store/internal/repositories"
	"github.com/kourai55k/store/pkg/render"
)

// GetProducts returns the list of all products.
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	const op = "web.productHandler.GetProducts"

	log := h.log.With(
		slog.String("op", op),
	)

	products, err := h.productService.GetProducts()
	if errors.Is(err, repositories.ErrProductNotFound) {
		log.Error("Products not found")
		render.RenderError(w, http.StatusNotFound, "Products not found")
		return
	}
	if err != nil {
		log.Error("Failed to get products")
		render.RenderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	err = render.RenderTemplate(w, "products.html", products)
	if err != nil {
		log.Error("Failed to render template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
