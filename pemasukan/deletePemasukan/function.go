package deletePemasukan

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"

	module "github.com/billblis/billblis_be/module"
)

func init() {
	functions.HTTP("DeletePemasukan", DeletePemasukan)
}

func DeletePemasukan(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "https://billblis.my.id")
		w.Header().Set("Access-Control-Allow-Methods", "DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "https://billblis.my.id")
	fmt.Fprintf(w, module.GCFHandlerDeletePemasukan("PASETOPUBLICKEY", "MONGOSTRING", "billblis", "pemasukan", r))

}
