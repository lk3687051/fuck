package pool
import ()

type Stat struct {
    mean  float64
  }
}
type Pool struct {
    Type  string
    Name  string
    NUms  int
  	Quotes    []*Quote
  }
}

func NewPool(type string, name string)  {
  p := Pool{
    Type: type,
    Name: name,
  }
}

func (p *Pool)Statistics()  {
  sum := 0
  num := 0
  for _, q := range p.Quotes {
    sum += q.PctChg
    num ++
  }
  p.Stat.mean = sum / num
}
