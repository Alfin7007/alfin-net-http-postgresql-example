package presentation

import (
	"encoding/json"
	"http/example/features/articles"
	"http/example/features/articles/presentation/request"
	"http/example/features/articles/presentation/response"
	"http/example/helper"
	"http/example/middlewares"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type ArticleHandler struct {
	artBussiness articles.Bussiness
}

func NewArticleHandler(artBussiness articles.Bussiness) *ArticleHandler {
	return &ArticleHandler{
		artBussiness: artBussiness,
	}
}

func (ah ArticleHandler) InsertArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	userID, _, errJWT := middlewares.JWTTokenCheck(r.Header.Get("Authorization"))
	if errJWT != nil {
		helper.FailedForbidden(w)
		return
	}

	artRequest := request.Article{}
	byteBody, errBody := ioutil.ReadAll(r.Body)
	if errBody != nil {
		helper.FailedBadRequestWithMSG(errBody.Error(), w)
		return
	}
	defer r.Body.Close()
	unmarshalErr := json.Unmarshal(byteBody, &artRequest)
	if unmarshalErr != nil {
		helper.FailedBadRequestWithMSG(unmarshalErr.Error(), w)
		return
	}
	artRequest.UserID = userID

	err := ah.artBussiness.InsertArticle(request.ToCore(artRequest))
	if err != nil {
		helper.FailedBadRequestWithMSG(err.Error(), w)
		return
	}
	helper.SuccessCreateNoData(w)
}

func (ah ArticleHandler) GetArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	articleID, errParam := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/articles/"))
	if errParam != nil {
		helper.FailedBadRequestWithMSG("invalid param", w)
		return
	}
	result, err := ah.artBussiness.GetArticle(articleID)
	if err != nil {
		helper.FailedBadRequestWithMSG(err.Error(), w)
		return
	}

	helper.SuccessGetData(response.FromSingleCore(result), w)
}

func (ah ArticleHandler) GetAllArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	result, err := ah.artBussiness.GetAll()
	if err != nil {
		helper.FailedBadRequestWithMSG(err.Error(), w)
	}

	helper.SuccessGetData(response.FromCoreList(result), w)
}
