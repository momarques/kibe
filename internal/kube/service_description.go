package kube

type ServiceDescription struct{}

func (s Service) Describe(c *ClientReady, serviceID string) ResourceDescription {
	return ServiceDescription{}
}

func (sd ServiceDescription) TabContent() []string {
	return []string{}
}

func (sd ServiceDescription) TabNames() []string {
	return []string{}
}
