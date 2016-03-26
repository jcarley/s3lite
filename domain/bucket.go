package domain

type Bucket struct {
	Id   string `json:"bucket_id, omitempty"`
	Name string `json:"name"`
}

func NewBucket(name string) *Bucket {
	return &Bucket{
		Name: name,
	}
}
