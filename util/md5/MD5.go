package MD5

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Encode(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}
