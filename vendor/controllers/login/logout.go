package login

import (
	"net/http"
	"services/cookieservice"
	"fmt"
)

func LogoutIndex(w http.ResponseWriter,r *http.Request) {

	cookieservice.DeleteCookie(w,"authCode")
	fmt.Println("Geldi")
	defer http.Redirect(w,r,"http://localhost:8080/anasayfa",http.StatusSeeOther)

}
