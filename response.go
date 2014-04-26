package dowser

import (
  "encoding/json"
  "log"
  )

type DowserResponse struct {
  Result string;
}

func (this *DowserResponse) ToString () string{
  res, err := json.Marshal(this);
  if err != nil {
    log.Fatal(err);
  }
  return string(res);
}
