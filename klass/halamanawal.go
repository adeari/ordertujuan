package klass

import (
    "net/http"
	"github.com/labstack/echo"
)

func Awal (c echo.Context) error {
  return c.Render(http.StatusOK, "index.html", map[string]interface{}{
      "a": "a",
  })
}
