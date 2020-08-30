package stock
import (

)
type Statistics struct {
  
  Daily  time.Time  `json:"date"`
  UpPencent  float  `json:"upPencent"`
  DownPencent  float  `json:"downPencent"`
}
