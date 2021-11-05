package httpd

import (
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func FileDownload(w http.ResponseWriter, r *http.Request, filePath string, fileName ...string) error {
	// check get file error, file not found or other error.
	if _, err := os.Stat(filePath); err != nil {
		http.ServeFile(w, r, filePath)
		return err
	}

	var fName string
	if len(fileName) > 0 && fileName[0] != "" {
		fName = fileName[0]
	} else {
		fName = filepath.Base(filePath)
	}
	//https://tools.ietf.org/html/rfc6266#section-4.3
	fn := url.PathEscape(fName)
	if fName == fn {
		fn = "filename=" + fn
	} else {
		/**
		  The parameters "filename" and "filename*" differ only in that
		  "filename*" uses the encoding defined in [RFC5987], allowing the use
		  of characters not present in the ISO-8859-1 character set
		  ([ISO-8859-1]).
		*/
		fn = "filename=" + fName + "; filename*=utf-8''" + fn
	}
	wHeader := w.Header()
	wHeader.Set("Content-Disposition", "attachment; "+fn)
	wHeader.Set("Content-Description", "File Transfer")
	wHeader.Set("Content-Type", "application/octet-stream")
	wHeader.Set("Content-Transfer-Encoding", "binary")
	wHeader.Set("Expires", "0")
	wHeader.Set("Cache-Control", "must-revalidate")
	wHeader.Set("Pragma", "public")
	http.ServeFile(w, r, filePath)
	return nil
}
