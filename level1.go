package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Ok       bool    `json:"ok"`
	Message  string `json:"message"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/checkout", checkout)
	log.Fatal(http.ListenAndServe(":8888", mux))
}

func checkout(w http.ResponseWriter, r *http.Request) {
	// ヘッダーの装着
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// リクエストメソッド確認
	if r.Method != http.MethodPost{
		log.Printf("Error: Bad Requset Method. This API is POST only. This Requset is %s", r.Method)
		response := &ErrorResponse{false,fmt.Sprintf("Bad Requset Method. This API is POST only. This Requset is %s", r.Method)}
		responseJson, err := json.Marshal(response)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w,`{"ok": false,"message": "Internal Server Error"}`)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w,string(responseJson))
		return
	}


	// リクエストjsonの読み込み
	d := json.NewDecoder(r.Body)
	var body interface{}
	err := d.Decode(&body)
	if err != nil {
		log.Printf("%+v\n", err)

		fmt.Fprintln(w, "%d","hoge")
	}
	fmt.Printf("Request = %+v\n", body)
	fmt.Fprintln(w, "{}")
}