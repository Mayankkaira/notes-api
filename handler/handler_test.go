package handler
import (
	"strings"
	"github.com/Mayankkaira/notes-api/storage"
	"github.com/Mayankkaira/notes-api/application"
	"net/http/httptest"
)

	s:=storage.NewInMemoryStorage()
	app:=application.NewApplication(s)
	handler:=NewHandler(app)
	req := httptest.NewRequest(
    http.MethodPost,
    "/notes",
    strings.NewReader(`{"title":"Learn Go","content":"Practice testing"}`),
)
	recoder:=httptest.NewRecoder()
	
