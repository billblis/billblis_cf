package getAllPemasukan

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"

	module "github.com/billblis/billblis_be/module"
)

func init() {
	functions.HTTP("getPemasukanFromUser", GetPemasukanFromUser)
}

func GetPemasukanFromUser(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "https://billblis.my.id")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "https://billblis.my.id")
	fmt.Fprintf(w, module.GCFHandlerGetPemasukanFromUser("PASETOPUBLICKEY", "MONGOSTRING", "billblis", "pemasukan", r))

}
