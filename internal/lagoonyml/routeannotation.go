package lagoonyml

import (
	"fmt"
	"regexp"
)

// ServerSnippet is the annotation for server snippets with ingress-nginx.
const ServerSnippet = "nginx.ingress.kubernetes.io/server-snippet"

// validSnippets is the allow-list of snippets that Lagoon will accept.
var validSnippets []*regexp.Regexp = []*regexp.Regexp{
	regexp.MustCompile(`^(rewrite +[^; ]+ +[^; ]+( (last|break|redirect|permanent))?;\n?)+$`),
	regexp.MustCompile(`^(add_header +[^; ]+ +[^;]+;\n?)+$`),
	regexp.MustCompile(`^(set_real_ip_from +[^; ]+;\n?)+$`),
}

// RouteAnnotation checks for valid annotations on defined routes.
func RouteAnnotation() Linter {
	return func(l *Lagoon) error {
		for eName, e := range l.Environments {
			for _, routeMap := range e.Routes {
				for rName, lagoonRoutes := range routeMap {
					for _, lagoonRoute := range lagoonRoutes {
						for iName, route := range lagoonRoute.Ingresses {
							if ss, ok := route.Annotations[ServerSnippet]; ok {
								valid := false
								for _, v := range validSnippets {
									if v.MatchString(ss) {
										valid = true
										break
									}
								}
								if !valid {
									return fmt.Errorf(
										"invalid %s annotation on environment %s, route %s, ingress %s: %s",
										ServerSnippet, eName, rName, iName, ss)
								}
							}
						}
					}
				}
			}
		}
		return nil
	}
}
