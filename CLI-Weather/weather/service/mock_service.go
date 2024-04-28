package service

type mytype struct {
    size string
    hash uint32
}

// interface for exposed methods
type myinterface interface {
    do() string
}

// constructor (optional)
func newMytype(size string, hash uint32) myinterface {
    return &mytype{size, hash}
}


func (r *mytype) do() string {
    return r.size
}