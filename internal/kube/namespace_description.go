package kube

type NamespaceDescription struct{}

func (n Namespace) Describe(c *ClientReady, namespaceID string) ResourceDescription {
	return NamespaceDescription{}
}

func (nd NamespaceDescription) TabContent() []string {
	return []string{}
}

func (nd NamespaceDescription) TabNames() []string {
	return []string{}
}

func (nd NamespaceDescription) SubContent(subContentIndex int) []string {
	return []string{}
}
