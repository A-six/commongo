package httpd

import (
	"net/http"
)

func FileDownloadBody(
	w http.ResponseWriter,
	headers map[string]string,
	respBody []byte,
	filename string,
) (int, error) {
	Header := w.Header()
	if headers != nil {
		for k, v := range headers {
			Header.Set(k, v)
		}
	}
	//
	Header.Set("Content-Disposition", "attachment; filename="+filename)
	Header.Set("Content-Description", "File Transfer")
	Header.Set("Content-Type", "application/octet-stream")
	Header.Set("Content-Transfer-Encoding", "binary")
	Header.Set("Expires", "0")
	Header.Set("Cache-Control", "must-revalidate")
	Header.Set("Pragma", "public")
	return w.Write(respBody)
}
