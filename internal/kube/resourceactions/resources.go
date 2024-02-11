package resourceactions

var SupportedResources = []Resource{
	&Pod{kind: "Pod"},
	&Service{kind: "Service"},
	&Ingress{kind: "Ingress"},
}

type Resource interface{ Kind() string }

type Pod struct{ kind string }

func (p *Pod) Kind() string { return p.kind }

type Service struct{ kind string }

func (p *Service) Kind() string { return p.kind }

type Ingress struct{ kind string }

func (p *Ingress) Kind() string { return p.kind }
