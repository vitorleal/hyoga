package pandorabox

import (
	"crypto/md5"
	"fmt"
	"io"
)

// MD5 transform a strind in MD5 string Hash
func MD5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}
