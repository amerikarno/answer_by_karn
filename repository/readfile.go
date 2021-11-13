package repository

import (
	"io/ioutil"
	"strings"
)

// Repository ...
type Repository interface{
	Readfile() (*string, error)
}

type repository struct{args string}

// NewRepository ...
func NewRepository(args string) Repository{
	return repository{args: args}
}


func (r repository) Readfile() (*string, error){
	txt, err := ioutil.ReadFile(r.args)
		if err != nil {
			return nil, err
		}
	
	var a []string
		for i, t := range txt {
		
			if i == (len(txt)-1){
				break
			}
			r := string(int(t)-128)
			a = append(a, r)
		}
	csv := strings.Join(a,"")
	return &csv, nil
}
