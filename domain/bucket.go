package domain

type Bucket struct {
  Id      int64  `db:"id"`
  Name    string `db:"name"`
  Created int64  `db:"created_on"`
  Updated int64  `db:"updated_on"`
}

func NewBucket(name string) *Bucket {
  return &Bucket{Name: name}
}
