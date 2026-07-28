package client

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// EncodeLMv1JSONBody serializes a value for LMv1 request signing (matches go-swagger client behavior).
func EncodeLMv1JSONBody(v interface{}) ([]byte, error) {
	if v == nil {
		return nil, nil
	}
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// lmComputeSignature builds the base64 LMv1 signature from body parts and request path.
func lmComputeSignature(accessKey, method, epoch string, bodyParts [][]byte, path string) string {
	h := hmac.New(sha256.New, []byte(accessKey))
	h.Write([]byte(method + epoch))
	for _, part := range bodyParts {
		if len(part) > 0 {
			h.Write(part)
		}
	}
	h.Write([]byte(path))
	hexDigest := hex.EncodeToString(h.Sum(nil))
	return base64.StdEncoding.EncodeToString([]byte(hexDigest))
}

// SignLMv1 returns a complete LMv1 Authorization header value.
func SignLMv1(accessID, accessKey, method string, body []byte, path string) string {
	epoch := strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
	var bodyParts [][]byte
	if len(body) > 0 {
		bodyParts = [][]byte{body}
	}
	signature := lmComputeSignature(accessKey, method, epoch, bodyParts, path)
	return fmt.Sprintf("LMv1 %s:%s:%s", accessID, signature, epoch)
}

// SignLMv1WithBodyParts returns a complete LMv1 Authorization header for multi-part bodies (e.g. file uploads).
func SignLMv1WithBodyParts(accessID, accessKey, method, path string, bodyParts [][]byte) string {
	epoch := strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
	signature := lmComputeSignature(accessKey, method, epoch, bodyParts, path)
	return fmt.Sprintf("LMv1 %s:%s:%s", accessID, signature, epoch)
}
