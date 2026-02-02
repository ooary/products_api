package handlers

import (
	"log"
	"net/http"
	"products_api/internal/models"
	"products_api/internal/services"
	"products_api/internal/utils"
	"strconv"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAllService()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, products)

}

func (h *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	idConv, err := strconv.ParseInt(idString, 10, 64)
	log.Println(idConv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	products, err := h.service.GetProductByDetailService(int(idConv))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, http.StatusOK, products)

}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var products models.Products
	utils.ReadJson(r, &products)
	err := h.service.CreateService(&products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusCreated, products)

}
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var products models.Products
	utils.ReadJson(r, &products)
	err := h.service.UpdateProduct(&products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, http.StatusOK, products)

}
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	idConv, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	is_deleted := h.service.DeleteProductService(idConv)
	if !is_deleted {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	message := map[string]string{
		"message": "success",
	}

	utils.WriteJson(w, http.StatusOK, message)
}
