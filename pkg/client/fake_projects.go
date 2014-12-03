package client

import "github.com/GoogleCloudPlatform/kubernetes/pkg/labels"
import projectapi "github.com/openshift/origin/pkg/project/api"

type FakeProjects struct {
	Fake      *Fake
	Namespace string
}

func (c *FakeProjects) List(label, field labels.Selector) (*projectapi.ProjectList, error) {
	c.Fake.Actions = append(c.Fake.Actions, FakeAction{Action: "list-projects"})
	return &projectapi.ProjectList{}, nil
}

func (c *FakeProjects) Get(name string) (*projectapi.Project, error) {
	c.Fake.Actions = append(c.Fake.Actions, FakeAction{Action: "get-project"})
	return &projectapi.Project{}, nil
}
