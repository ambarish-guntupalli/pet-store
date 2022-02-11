package routes

import (
	"context"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

type PetStore struct {
	IntegrationUrl string
}

func (ps PetStore) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", ps.List)    // GET /pets - Read a list of pets.
	r.Post("/", ps.Create) // POST /pets - Create a new pet.

	r.Route("/{id}", func(r chi.Router) {
		r.Use(ps.PetCtx)
		r.Get("/", ps.Get) // GET /pets/{id} - Read a single pet by :id.

		/*Currently petstore integration api has no delete

		Http method is DELETE

		200 (OK). 404 (Not Found), if ID not found or invalid

		DELETE operation is considered idempotent, since the
		server state remains the same how many times the endpoint is called
		irrespective of the status code return

		DELETE: curl -H "Content-Type: application/json" -X DELETE 'http://localhost:8080/pets/{:id}'
		*/
		//r.Delete("/", Delete) // DELETE /pets/{id} - Delete a single pet by :id.
	})
	return r
}

// List Request Handler - GET /pets - Read a list of pets.
func (ps PetStore) List(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(ps.IntegrationUrl)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Create Request Handler - POST /pets - Create a new pet.
func (ps PetStore) Create(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Post(ps.IntegrationUrl, "application/json", r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (ps PetStore) PetCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "id", chi.URLParam(r, "id"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Get Request Handler - GET /pets/{id} - Read a single pet by :id.
func (ps PetStore) Get(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)

	resp, err := http.Get(ps.IntegrationUrl + "/" + id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Delete Request Handler - DELETE /pets/{id} - Delete a single pet by :id.
/*func (ps PetStore) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)
	client := &http.Client{}

	req, err := http.NewRequest("DELETE", ps.IntegrationUrl + "/" + id, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := client.Do(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}*/
