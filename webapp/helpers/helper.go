package helpers

import (
	"net/http"

	"github.com/eaciit/toolkit"

	"git.eaciitapp.com/sebar/knot"
)

func WriteResultError(k *knot.WebContext, res *toolkit.Result, errorText string) {
	res.SetErrorTxt(errorText)
	k.WriteJSON(res, http.StatusUnauthorized)
}

func WriteResultErrorOK(k *knot.WebContext, res *toolkit.Result, errorText string) {
	res.SetErrorTxt(errorText)
	k.WriteJSON(res, http.StatusOK)
}

func WriteResultOK(k *knot.WebContext, res *toolkit.Result, data interface{}) {
	res.SetData(data)
	k.WriteJSON(res, http.StatusOK)
}
