package controller

import (
	"encoding/json"
	"fmt"
	"go-shortURL/configs"
	"go-shortURL/dao"
	"go-shortURL/models"
	"go-shortURL/utils"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

var urlController *UrlController

type UrlController struct {
	converter dao.UrlConverter
}

func NewUrlController() *UrlController {
	config := &configs.Config{}
	config.ParseConfig()
	return &UrlController{
		converter: dao.NewRedisClient(config.RedisAddr, config.RedisPassword, config.RedisDb),
	}
}

func GetUrlController() *UrlController {
	if urlController == nil {
		urlController = NewUrlController()
	}
	return urlController
}

func (c *UrlController) CreateShortLink(w http.ResponseWriter, r *http.Request) {
	var req models.ShortLinkReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, utils.StatusError{
			Code: http.StatusBadRequest,
			Err:  err,
		})
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		utils.RespondWithError(w, utils.StatusError{
			Code: http.StatusBadRequest,
			Err:  err,
		})
		return
	}

	defer r.Body.Close()

	shortLink, err := c.converter.Shorten(req.Url, req.ExpireWithin)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	res := models.ShortLinkRes{ShortLink: shortLink}
	utils.RespondWithJson(w, http.StatusOK, res)
}

func (c *UrlController) RedirectOriginalUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	originalUrl, err := c.converter.Unshorten(vars["shortlink"])
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}
	fmt.Printf("%s", originalUrl)
	http.Redirect(w, r, originalUrl, http.StatusTemporaryRedirect)
}

func (c *UrlController) GetShortLinkDetail(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()

	detailJson, err := c.converter.GetUrlDetail(vals.Get("shortlink"))
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}
	detailStruct := &models.UrlDetail{}
	err = json.Unmarshal([]byte(detailJson), detailStruct)
	if err != nil {
		utils.RespondWithError(w, err)
		return
	}

	utils.RespondWithJson(w, http.StatusOK, detailStruct)
}
