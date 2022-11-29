package router

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestPingRoute(t *testing.T){
    router, _ := SetupRouter()

    w := httptest.NewRecorder()
    req, err := http.NewRequest(http.MethodGet, "/ping", nil)
    if err != nil {
        t.Fatalf("Fail to request /ping: %s", err)
    }

    router.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Fatal("Server error: Returned ", w.Code, " instead of ", http.StatusOK)

    }

    if w.Body.String() != "pong" {

        t.Fatal("Server error: Returned ", w.Body, " instead of ping" )
    }

    
}
