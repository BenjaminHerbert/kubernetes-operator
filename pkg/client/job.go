package client

import (
	"context"
	"net/url"

	"github.com/bndr/gojenkins"
)

func (jenkins *jenkins) GetBuild(jobName string, number int64) (*gojenkins.Build, error) {
	job, err := jenkins.GetJob(context.TODO(), jobName)
	if err != nil {
		return nil, err
	}

	// https://github.com/bndr/gojenkins/issues/176
	// workaround begin
	jobURL, err := url.Parse(job.Raw.URL)
	if err != nil {
		return nil, err
	}
	job.Raw.URL = jobURL.RequestURI()
	// workaround end

	build, err := job.GetBuild(context.TODO(), number)

	if err != nil {
		return nil, err
	}
	return build, nil
}
