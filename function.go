package billblis

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	billblis_be "github.com/billblis/billblis_be"
)

func init() {
	functions.HTTP("Billblis", billblisPost)
}

func billblisPost(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "https://billblis.github.io")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "https://billblis.github.io")
	fmt.Fprintf(w, billblis_be.GCFPostHandler("PASETOPRIVATEKEY", "MONGOULBI", "billblis", "user", r))

}

// package billblis

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
// 	"github.com/petapedia/peda"
// )

// func init() {
// 	functions.HTTP("PetaPedia", petaPediaPost)
// }

// func petaPediaPost(w http.ResponseWriter, r *http.Request) {
// 	// Set CORS headers for the preflight request
// 	if r.Method == http.MethodOptions {
// 		w.Header().Set("Access-Control-Allow-Origin", "https://jscroot.github.io")
// 		w.Header().Set("Access-Cpedontrol-Allow-Methods", "POST")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
// 		w.Header().Set("Access-Control-Max-Age", "3600")
// 		w.WriteHeader(http.StatusNoContent)
// 		return
// 	}
// 	// Set CORS headers for the main request.
// 	w.Header().Set("Access-Control-Allow-Origin", "https://jscroot.github.io")
// 	fmt.Fprintf(w, peda.GCFPostHandler("PASETOPRIVATEKEY", "MONGOULBI", "petapedia", "user", r))

// }
