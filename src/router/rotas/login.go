package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasLogin = Rota{
	URI:                "/login",
	Method:             http.MethodPost,
	Funcao:             controllers.Login,
	RequerAutenticacao: false,
}
