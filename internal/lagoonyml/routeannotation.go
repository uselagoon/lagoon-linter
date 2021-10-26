package lagoonyml

import (
	"fmt"
	"regexp"
	"strings"
)

// ServerSnippet is the annotation for server snippets with ingress-nginx.
const ServerSnippet = "nginx.ingress.kubernetes.io/server-snippet"

// validSnippets is the allow-list of snippets that Lagoon will accept.
var validServerSnippets = regexp.MustCompile(
	`^(rewrite +[^; ]+ +[^; ]+( (last|break|redirect|permanent))?;|` +
		`add_header +[^; ]+ +[^;]+;|` +
		`set_real_ip_from +[^; ]+;|` +
		` )+$`)

// validate returns true if the annotations are valid, and false otherwise.
func validate(annotations map[string]string, r *regexp.Regexp, annotation string) (string, bool) {
	if ss, ok := annotations[annotation]; ok {
		for _, line := range strings.Split(ss, "\n") {
			line = strings.TrimSpace(line)
			if len(line) > 0 && !r.MatchString(line) {
				return line, false
			}
		}
	}
	return "", true
}

// RouteAnnotation checks for valid annotations on defined routes.
func RouteAnnotation() Linter {
	return func(l *Lagoon) error {
		for eName, e := range l.Environments {
			for _, routeMap := range e.Routes {
				for rName, lagoonRoutes := range routeMap {
					for _, lagoonRoute := range lagoonRoutes {
						for iName, ingress := range lagoonRoute.Ingresses {
							if annotation, ok := validate(ingress.Annotations, validServerSnippets, ServerSnippet); !ok {
								return fmt.Errorf(
									"invalid %s annotation on environment %s, route %s, ingress %s: %s",
									ServerSnippet, eName, rName, iName, annotation)
							}
						}
					}
				}
			}
		}
		return nil
	}
}
