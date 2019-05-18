package klass

import (
    "net/http"
    "database/sql"
	"github.com/labstack/echo"
    _ "github.com/go-sql-driver/mysql"
    "strconv"
    "html/template"
    "fmt"
)

func Awal (c echo.Context) error {
  db, _ := sql.Open("mysql", DB_USER+":"+DB_PASS+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_NAME)
  defer db.Close()
  
  rows, _ := db.Query("SELECT id,nama FROM tb_kota ORDER BY nama")
  defer rows.Close()
  
  kotanyaoptionstr := ""
  for rows.Next() {
    id:=0
    nama := ""
    rows.Scan(&id, &nama)
    kotanyaoptionstr += "<option value="+strconv.Itoa(id)+">"+nama+"</option>"
  }
  kotadikembalikannyaoptionstr := "<option value=''>Kota tujuan</option>"+kotanyaoptionstr
  
  rowmobils, _ := db.Query("SELECT id,merk,sewaperhari FROM tb_mobil ORDER BY merk")
  defer rowmobils.Close()
  mobiloptions := ""
  for rowmobils.Next() {
    id:=0
    merk:=""
    var sewaperhari float64
    rowmobils.Scan(&id, &merk, &sewaperhari)
    mobiloptions +=  "<option value="+strconv.Itoa(id)+">"+merk+" @ Rp "+FormatThousand(fmt.Sprintf("%.f", sewaperhari))+" /hari</option>"
  }
  
  return c.Render(http.StatusOK, "index.html", map[string]interface{}{
      "kotadikembalikannyaoptionstr": template.HTML(kotadikembalikannyaoptionstr),
      "kotanyaoptionstr": template.HTML(kotanyaoptionstr),
      "mobiloptions": template.HTML(mobiloptions),
  })
}

func FormatThousand(str string) string {
  if (len(str) < 4) {return str}
  strs := ""
  ketiga := 0
  for slen := len(str); slen > 0; slen-- {
    strs = string(str[(slen-1):slen]) + strs
    ketiga++
    if (ketiga == 3) {
      if (slen > 1) {
        strs = "." +strs
      }
      ketiga = 0
    }
  }
  return strs  
}

func Ui (c echo.Context) error {
  return c.Render(http.StatusOK, "ui.html", map[string]interface{}{
      "a": "a",
  })
}
