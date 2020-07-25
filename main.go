package main

import (
	"fmt"

	devfile "github.com/devfile/parser/pkg/devfile"
)

// Testing end-point
func main() {
	devfile, err := ParseDevfile("devfile.yaml")
	if err != nil {
		fmt.Println(err)
	} else {
		for _, component := range devfile.Data.GetAliasedComponents() {
			if component.Dockerfile != nil {
				fmt.Println(component.Dockerfile.DockerfileLocation)
				break
			}
		}
	}

}

// End point defined to return parsed devfile object
func ParseDevfile(devfileLocation string) (devfileobj devfile.DevfileObj, err error) {

	var devfile devfile.DevfileObj
	devfile, err = ParseAndValidate(devfileLocation)
	return devfile, err
}
