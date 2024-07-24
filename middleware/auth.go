package middleware

import (
	
	"net/http"


)

func ValidateUser(f http.HandlerFunc) http.HandlerFunc {
return  func (w http.ResponseWriter, r *http.Request)  {
	// username := r.Header.Get("username")
	// password := r.Header.Get("password")
	// if data.Users[username]!= password || username ==""{
	// 	w.Write([]byte("Failed to authentificate"))
	// }
	f(w,r)
}
}

func ValidateOwner( f http.HandlerFunc )http.HandlerFunc  {
	return func (w http.ResponseWriter, r *http.Request)  {
		// username := r.Header.Get("username")
		// password := r.Header.Get("password")
		// IfOwner := r.Header.Get("userType")
		// if IfOwner != "owner" {
		// 	w.Write([]byte("youre not the owner"))
		// }

		// if data.Users[username]!= password || username ==""{
		// 	w.Write([]byte("failed to authenticate"))
		// }

		f(w, r)
	}
}
func TrackNumberOfRequests(f http.Handler) http.Handler {
	return http.HandlerFunc(
		func (w http.ResponseWriter, r *http.Request)  {
			// data.NumberOfRequests= data.NumberOfRequests+1
			// fmt.Println("Request number : ", data.NumberOfRequests)

			f.ServeHTTP(w,r)
		})
}