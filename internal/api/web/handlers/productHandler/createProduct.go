package productHandler

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/kourai55k/store/internal/domain/models"
	"github.com/kourai55k/store/pkg/render"
)

var (
	ErrNameNotProvided    = errors.New("product name not provided")
	ErrPriceNotProvided   = errors.New("product price not provided")
	ErrMeasureNotProvided = errors.New("product measure not provided")
)

type ProductDTO struct {
	Price        float64
	Stock        int
	Measure      string
	CategoryName string
	Params       string
	Name         string
}

func (dto *ProductDTO) Validate() error {
	if dto.Name == "" {
		return ErrNameNotProvided
	}

	if dto.Price <= 0 {
		return ErrPriceNotProvided
	}

	if dto.Measure == "" {
		return ErrMeasureNotProvided
	}

	return nil
}

func (dto *ProductDTO) ParseFormToDTO(r *http.Request) {
	dto.Name = r.FormValue("name")
	priceStr := r.FormValue("price")
	price, _ := strconv.ParseFloat(priceStr, 64)
	dto.Price = price
	stockStr := r.FormValue("stock")
	stock, _ := strconv.Atoi(stockStr)
	dto.Stock = stock
	dto.Measure = r.FormValue("measure")
	dto.CategoryName = r.FormValue("categoryName")
	dto.Params = r.FormValue("params")
}

// CreateProduct is a handler for POST /product/new
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	const op = "web.productHandler.CreateProduct"

	log := h.log.With(slog.String("op", op))

	if err := r.ParseForm(); err != nil {
		log.Error("Failed to parse form", "error", err)
		render.RenderError(w, http.StatusBadRequest, "Invalid form data")
		return
	}

	var dto ProductDTO
	dto.ParseFormToDTO(r)

	if err := dto.Validate(); err != nil {
		log.Error("Validation error", "error", err.Error())
		switch err {
		case ErrNameNotProvided:
			render.RenderError(w, http.StatusBadRequest, "не указано имя товара")
			return
		case ErrPriceNotProvided:
			render.RenderError(w, http.StatusBadRequest, "не указана цена товара")
			return
		case ErrMeasureNotProvided:
			render.RenderError(w, http.StatusBadRequest, "не указана единица измерения товара")
			return
		default:
			render.RenderError(w, http.StatusBadRequest, "неправильно указаны данные товара")
		}
	}

	product := models.Product{
		Name:         dto.Name,
		Price:        dto.Price,
		Stock:        dto.Stock,
		Measure:      dto.Measure,
		CategoryName: dto.CategoryName,
		// ImageURL:     ,
		Params: dto.Params,
	}

	err := h.productService.SaveProduct(product)
	if err != nil {
		log.Error("Failed to save product", "error", err)
		render.RenderError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}

// CreateProduct is a handler for GET /product/new
func (h *ProductHandler) CreateProductPage(w http.ResponseWriter, r *http.Request) {
	const op = "web.productHandler.CreateProductPage"

	log := h.log.With(slog.String("op", op))
	// TODO: получить список категорий и передать его в шаблон
	err := render.RenderTemplate(w, "create_product.html", nil)
	if err != nil {
		log.Error("Failed to render template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
