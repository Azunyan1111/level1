package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Request struct {
	Order []string `json:"order"`
}

type Response struct {
	Ok     bool     `json:"ok"`
	Amount int64    `json:"amount"`
	Items  []string `json:"items"`
}

type ErrorResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

// 問題の返却するレスポンスの「items」の値が文字列型なのでそれに合わせます。
type Menu struct {
	Id    string
	Name  string
	Price int64
}

// メニュー情報を格納
var Menus []Menu

func main() {
	setMenus()
	mux := http.NewServeMux()
	mux.HandleFunc("/api/checkout", checkout)
	log.Fatal(http.ListenAndServe(":8888", mux))
}

// gistということなので変数にデータを入れます。
func setMenus() {
	Menus = append(Menus, Menu{"101", "ハンバーガー", 100})
	Menus = append(Menus, Menu{"102", "チーズバーガー", 130})
	Menus = append(Menus, Menu{"103", "ダブルチーズバーガー", 320})
	Menus = append(Menus, Menu{"104", "てりやきバーガー", 320})
	Menus = append(Menus, Menu{"105", "ビッグバーガー", 380})
	Menus = append(Menus, Menu{"201", "ポテトS", 150})
	Menus = append(Menus, Menu{"202", "ポテトM", 270})
	Menus = append(Menus, Menu{"203", "ポテトL", 320})
	Menus = append(Menus, Menu{"204", "サラダ", 280})
	Menus = append(Menus, Menu{"301", "コーラS", 100})
	Menus = append(Menus, Menu{"302", "コーラM", 220})
	Menus = append(Menus, Menu{"303", "コーラL", 250})
	Menus = append(Menus, Menu{"304", "オレンジS", 150})
	Menus = append(Menus, Menu{"305", "オレンジM", 240})
	Menus = append(Menus, Menu{"306", "オレンジL", 270})
	Menus = append(Menus, Menu{"307", "ホットコーヒーS", 100})
	Menus = append(Menus, Menu{"308", "ホットコーヒーM", 150})
}

func checkout(w http.ResponseWriter, r *http.Request) {
	// ヘッダーの装着
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// リクエストメソッド確認
	if r.Method != http.MethodPost {
		log.Printf("Error: Bad Requset Method. This API is POST only. This Requset is %s", r.Method)
		response := &ErrorResponse{false, fmt.Sprintf("Bad Requset Method. This API is POST only. This Requset is %s", r.Method)}
		responseJson, err := json.Marshal(response)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, `{"ok": false,"message": "Internal Server Error"}`)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, string(responseJson))
		return
	}

	// リクエストjsonの読み込み
	var RequestBody Request
	err := json.NewDecoder(r.Body).Decode(&RequestBody)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, `{"ok": false,"message": "Internal Server Error"}`)
		return
	}

	// メニュー内容の読み込みと合計金額の計算
	var amount int64
requestOrder:
	for _, orderOne := range RequestBody.Order {
		for _, menu := range Menus {
			if menu.Id == orderOne {
				amount += menu.Price
				continue requestOrder
			}
		}
		// アイテムが存在しない場合
		response := &ErrorResponse{false, "item_not_found"}
		responseJson, err := json.Marshal(response)
		if err != nil {
			log.Printf("Error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, `{"ok": false,"message": "Internal Server Error"}`)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, string(responseJson))
		return
	}

	// 合計金額を返却する
	response := &Response{true, amount, RequestBody.Order}
	responseJson, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, `{"ok": false,"message": "Internal Server Error"}`)
		return
	}
	fmt.Fprintln(w, string(responseJson))
}
