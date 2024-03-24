package kube

type ServiceDescription struct{}

func (s Service) Describe(c *ClientReady) ResourceDescription {
	return ServiceDescription{}
}

func (sd ServiceDescription) TabContent() []string {
	return []string{}
}

func (sd ServiceDescription) TabNames() []string {
	return []string{}
}

func (sd ServiceDescription) SubContent(subContentIndex int) []string {
	return []string{}
}
