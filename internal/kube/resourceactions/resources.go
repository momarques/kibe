package resourceactions

import "github.com/momarques/kibe/internal/kube/pod"

var SupportedResources = []Resource{
	pod.New(),
	&Service{kind: "Service"},
	&Ingress{kind: "Ingress"},
}

type Resource interface{ Kind() string }

type Service struct{ kind string }

func (p *Service) Kind() string { return p.kind }

type Ingress struct{ kind string }

func (p *Ingress) Kind() string { return p.kind }
