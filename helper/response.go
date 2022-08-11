package helper

import (
	"encoding/json"
	"net/http"
)

func AuthOK(id int, token string, w http.ResponseWriter) {
	msg, _ := json.Marshal(map[string]interface{}{
		"message": "success",
		"code":    200,
		"id":      id,
		"token":   token,
	})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}

func SuccessCreateNoData(w http.ResponseWriter) {
	msg, _ := json.Marshal(map[string]interface{}{
		"message": "success",
		"code":    201,
	})
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(msg))
}
func SuccessGetData(data interface{}, w http.ResponseWriter) {
	msg, _ := json.Marshal(map[string]interface{}{
		"message": "success",
		"data":    data,
		"code":    200,
	})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))

}

func FailedBadRequest(w http.ResponseWriter) {
	msg, _ := json.Marshal(map[string]interface{}{
		"message": "bad request",
		"code":    400,
	})
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(msg))

}

func FailedBadRequestWithMSG(content string, w http.ResponseWriter) {
	msg, _ := json.Marshal(map[string]interface{}{
		"message": "bad request => " + content,
		"code":    400,
	})
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(msg))
}

func FailedNotFound(content string, w http.ResponseWriter) {
	msg, _ := json.Marshal(map[string]interface{}{
		"message": content,
		"code":    404,
	})
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(msg))
}

func FailedForbidden(w http.ResponseWriter) {
	msg, _ := json.Marshal(map[string]interface{}{
		"message": "unauthorized",
		"code":    403,
	})
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte(msg))
}

func FailedMethodNotAllowed(w http.ResponseWriter) {
	msg, _ := json.Marshal(map[string]interface{}{
		"message": "method not allowed",
		"code":    405,
	})
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte(msg))
}
