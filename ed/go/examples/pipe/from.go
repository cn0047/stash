package downloader

import (
  "archive/zip"
  "bytes"
  "fmt"
  "io"
)

func fromZip(src io.ReadCloser, next func(r io.ReadCloser, file string) error) error {
  b := new(bytes.Buffer)
  _, err := b.ReadFrom(src)
  if err != nil {
    return fmt.Errorf("failed to read from reader, error: %s", err)
  }
  br := bytes.NewReader(b.Bytes())

  zr, err := zip.NewReader(br, br.Size())
  if err != nil {
    return fmt.Errorf("filed to create new zip reader, error: %s", err)
  }

  for _, zf := range zr.File {
    r, err := zf.Open()
    if err != nil {
      return fmt.Errorf("failed to ope file in zip archive, error: %s", err)
    }

    er := next(r, zf.Name)
    if er != nil {
      return fmt.Errorf("failed to handle zip archive, error: %s", er)
    }
  }

  return nil
}
