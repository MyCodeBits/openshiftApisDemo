package main

import "io/ioutil"

func ReadFile(path string) ([]byte, error) {
	// TODO : Check for path existence
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		// TODO : Better error message.
		panic(err)
	}
	return dat, err
}
