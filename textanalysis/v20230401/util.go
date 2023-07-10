package v20230401

import (
	"net/url"
	"path"
)

func ParseJobID(jobLocation string) (string, error) {
	u, err := url.Parse(jobLocation)
	if err != nil {
		return "", err
	}
	_, jobID := path.Split(u.Path)
	return jobID, nil
}
