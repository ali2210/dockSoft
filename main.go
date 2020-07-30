package main


 	/*Docker service .... */
import (
		"fmt"
		"net/http"
		"html/template"
)



func main(){

	fmt.Println("Listening.... :5000 ")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		temp := template.Must(template.ParseFiles("index.html"))
		if r.Method == "GET" {
			fmt.Println("Method:" + r.Method)																										
			temp.Execute(w, "Home")
		}

	})

	http.ListenAndServe(":5000", nil)
}
