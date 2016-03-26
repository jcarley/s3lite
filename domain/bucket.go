package domain

type Bucket struct {
	Name string `json:"name"`
}

func NewBucket() *Bucket {
	return &Bucket{}
}
