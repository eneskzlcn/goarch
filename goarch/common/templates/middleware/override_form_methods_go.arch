package middleware

import "net/http"

/*OverrideFormMethods is a middleware for the problem that
we can not send a method except GET and POST from html forms.
So this method will look at a hidden input named _method which
should keep the value of the actual method DELETE, PUT , PATCH or etc.
If you placed an input like that to your form, it will find it and replace
the operation with correct method.
Another method is putting a header value generally with the key
"X-HTTP-Method-Override" which keeps the value of target method.
If input not placed then that middleware looks at the header and
get the target from there.

I saw the method and the idea from Alex Edwards
from https://www.alexedwards.net/blog/http-method-spoofing
*/
func OverrideFormMethods(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			next.ServeHTTP(w, r)
			return
		}
		method := r.PostFormValue("_method")
		if method == "" {
			method = r.Header.Get("X-HTTP-Method-Override")
		}
		if method == http.MethodDelete || method == http.MethodPatch || method == http.MethodPut {
			r.Method = method
		}

		next.ServeHTTP(w, r)
	})
}
