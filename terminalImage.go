package terminalImage

import (
  "fmt"
  "encoding/base64"
)

func Print(data []byte){
  encodedContents := base64.StdEncoding.EncodeToString(data)
  fmt.Printf("\x1b]1337;File=name=fractal;inline=1:%s\a\n", encodedContents)
}
