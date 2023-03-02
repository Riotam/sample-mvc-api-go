package controller

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/Riotam/sample-mvc-api-go/test"
)

// URLとハンドラ関数を関連付けるマルチプレクサと呼ばれる構造体
// 複数のテストで共通して使うのでパッケージ変数として定義
var mux *http.ServeMux

// TestMain は前後処理のようなテストのフロー制御を行うための関数
// 前後処理を行う必要がない場合は不要
func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

// setUp は前処理用の関数 (名前は何でも良い)
func setUp() {
	target := NewRouter(&test.MockTodoController{})
	mux = http.NewServeMux()
	mux.HandleFunc("/todos/", target.HandleTodosRequest)

}

func TestGetTodos(t *testing.T) {
	// リクエスト生成
	r, _ := http.NewRequest("GET", "/todos/", nil)

	// レスポンスを取得するための処理
	w := httptest.NewRecorder()

	// テスト対象のハンドラ関数にリクエストを送信
	mux.ServeHTTP(w, r)

	// MockTodoControllerで設定しているステータスコード200が返却されていることを確認
	if w.Code != 200 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestPostTodo(t *testing.T) {
	json := strings.NewReader(`{"title":"test-title","content":"test-content"}`)
	r, _ := http.NewRequest("POST", "/todos/", json)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	if w.Code != 201 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestPutTodo(t *testing.T) {
	json := strings.NewReader(`{"title":"test-title","contents":"test-content"}`)
	r, _ := http.NewRequest("PUT", "/todos/2", json)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	if w.Code != 204 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestDeleteTodo(t *testing.T) {
	r, _ := http.NewRequest("DELETE", "/todos/2", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	if w.Code != 204 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestInvalidMethod(t *testing.T) {
	r, _ := http.NewRequest("PATCH", "/todos/", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	if w.Code != 405 {
		t.Errorf("Response cod is %v", w.Code)
	}
}
