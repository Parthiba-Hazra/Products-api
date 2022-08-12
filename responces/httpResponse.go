package responces

import "net/http"

type HTTPResponce struct {
	Status  int         `json:"status"`
	Sucess  bool        `json:"sucess"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewHTTPResponce(status int, data interface{}) HTTPResponce {
	switch status {
	case http.StatusBadRequest,
		http.StatusInternalServerError,
		http.StatusUnauthorized,
		http.StatusForbidden,
		http.StatusNotFound,
		http.StatusRequestTimeout:

		if e, ok := data.(error); ok {
			return HTTPResponce{
				Status:  status,
				Sucess:  false,
				Message: e.Error(),
			}
		}

		return HTTPResponce{
			Status:  status,
			Sucess:  false,
			Message: data.(string),
		}
	default:
		return HTTPResponce{
			Status: status,
			Sucess: true,
			Data:   data,
		}
	}
}
