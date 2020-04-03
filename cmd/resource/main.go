package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/cloudboss/ofcourse/ofcourse"

	"github.com/pivotal/concourse-kpack-resource/k8s"
	"github.com/pivotal/concourse-kpack-resource/resource"
)

func main() {
	switch filepath.Base(os.Args[0]) {
	case "check":
		ofcourse.Check(&concourseResource{})
	case "in":
		ofcourse.In(&concourseResource{})
	case "out":
		ofcourse.Out(&concourseResource{})
	default:
		log.Fatalf("invalid args %s", os.Args)
	}
}

type concourseResource struct{}

func (concourseResource) Check(ofcourseSource ofcourse.Source, version ofcourse.Version, env ofcourse.Environment, logger *ofcourse.Logger) ([]ofcourse.Version, error) {
	clientSet, err := k8s.Authenticate(ofcourseSource)
	if err != nil {
		return nil, err
	}

	source, err := resource.NewSource(ofcourseSource)
	if err != nil {
		return nil, err
	}

	return resource.Check(clientSet, source, version, env, logger)
}

func (concourseResource) In(outputDirectory string, source ofcourse.Source, params ofcourse.Params, version ofcourse.Version, env ofcourse.Environment, logger *ofcourse.Logger) (ofcourse.Version, ofcourse.Metadata, error) {

	return version, nil, nil
}

func (concourseResource) Out(inputDirectory string, source ofcourse.Source, params ofcourse.Params, env ofcourse.Environment, logger *ofcourse.Logger) (ofcourse.Version, ofcourse.Metadata, error) {

	return nil, nil, nil
}
