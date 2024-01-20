package utils

import (
   "crypto/rand"
   "encoding/base64"
)

func Substring(s string, start int, end int) string {
   r := []rune(s)
   return string(r[start:end])
}

func GenerateRandomString(length int) string {
   b := make([]byte, length)
   _, err := rand.Read(b)
   if err != nil {
      panic(err)
   }

   rndstring := base64.URLEncoding.EncodeToString(b)
   return Substring(rndstring, 0, length)
}

