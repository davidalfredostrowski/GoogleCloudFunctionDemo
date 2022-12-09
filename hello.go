package hello
import "net/http"
func F(w http.ResponseWriter, r *http.Request){
        msg :=[]byte("hello world")
        w.Write(msg)
}
