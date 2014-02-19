package domain

type Bucket struct {
  Id uint
  Name string
}

func NewBucket(name string) *Bucket {
  return &Bucket{Name: name}
}
