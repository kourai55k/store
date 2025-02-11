package productHandler

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/kourai55k/store/internal/repositories"
	"github.com/kourai55k/store/pkg/render"
)

// GetProductByID returns a single product by its ID.
func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	const op = "web.productHandler.GetProductByID"

	log := h.log.With(
		slog.String("op", op),
	)

	idString := r.PathValue("id")
	if idString == "" {
		log.Error("Missing product ID")
		render.RenderError(w, http.StatusBadRequest, "Missing product ID")
		return
	}

	var id uint
	_, err := fmt.Sscanf(idString, "%d", &id)
	if err != nil {
		log.Error("Invalid product ID")
		render.RenderError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := h.productService.GetProductByID(id)
	if errors.Is(err, repositories.ErrProductNotFound) {
		log.Error("Product not found", "product_id", id)
		render.RenderError(w, http.StatusNotFound, "Product not found")
		return
	}
	if err != nil {
		log.Error("Failed to get product", "product_id", id, "error", err)
		render.RenderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	err = render.RenderTemplate(w, "product.html", product)
	if err != nil {
		log.Error("Rendering error", "error", err)
		render.RenderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}
