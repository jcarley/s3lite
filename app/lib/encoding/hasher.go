package encoding

import (
  "hash"
  "bytes"
  "crypto/md5"
  "crypto/sha1"
  "hash/adler32"
  "encoding/base64"
)

type Hasher struct {
  Algorithm string
  Buffer *bytes.Buffer
  Hash hash.Hash
}

func NewHasher(algorithm string, buffer *bytes.Buffer) *Hasher {
  hasher := Hasher{Algorithm: algorithm, Buffer: buffer}
  hasher.initAlgorithm()
  return &hasher
}

func (h *Hasher) initAlgorithm() {
  switch h.Algorithm {
  case "SHA1":
    h.Hash = sha1.New()
  case "MD5":
    h.Hash = md5.New()
  case "ADLER32":
    h.Hash = adler32.New()
  }
}

func (h *Hasher) HashString() string {
  h.Hash.Write(h.Buffer.Bytes())
  sum := h.Hash.Sum(nil)
  return base64.URLEncoding.EncodeToString(sum)
}

