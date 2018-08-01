package cookieservice

import (
	"net/http"
	"time"
)

func NewCookie(w http.ResponseWriter,cookiename ,val string,hour int){
	expiration := time.Now().Add(time.Duration(hour)*time.Hour)

	cookie := http.Cookie{
		Name:cookiename,
		Value:val,
		Expires:expiration,
		Path:"/",
	}

	http.SetCookie(w,&cookie)
}

func ReadCookie(r *http.Request,cookieName string)(cookie *http.Cookie){
	var err error
	cookie ,err = r.Cookie(cookieName)
	if err != nil {
		return nil
	}
	return cookie
}

func DeleteCookie(w http.ResponseWriter,name string){
	expiration := time.Now().Add(-1)
	cookie := http.Cookie{
		Name:name,
		Value:"",
		Expires:expiration,
		Path:"",
	}
	http.SetCookie(w,&cookie)
}

func GetCookieValue(r *http.Request,cookieName string)string{
	var err error
	cookie,err := r.Cookie(cookieName)
	if err!=nil{
		return ""
	}
	if cookie == nil{
		return ""
	}
	return cookie.Value
}