package lagoonyml

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	authSnippet          = "nginx.ingress.kubernetes.io/auth-snippet"
	configurationSnippet = "nginx.ingress.kubernetes.io/configuration-snippet"
	modsecuritySnippet   = "nginx.ingress.kubernetes.io/modsecurity-snippet"
	serverSnippet        = "nginx.ingress.kubernetes.io/server-snippet"
)

// validSnippets is the allow-list of snippets that Lagoon will accept.
// Currently these are only valid in server-snippet and configuration-snippet
// annotations.
var validSnippets = regexp.MustCompile(
	`^(rewrite +[^; ]+ +[^; ]+( (last|break|redirect|permanent))?;|` +
		`add_header +([^; ]+|"[^"]+"|'[^']+') +([^; ]+|"[^"]+"|'[^']+')( always)?;|` +
		`set_real_ip_from +[^; ]+;|` +
		`more_set_headers +(-s +("[^"]+"|'[^']+')|-t +("[^"]+"|'[^']+')|("[^"]+"|'[^']+'))+;|` +
		` )+$`)

// validate returns true if the annotations are valid, and false otherwise.
func validate(annotations map[string]string, r *regexp.Regexp,
	annotation string) (string, bool) {
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
							// auth-snippet
							if _, ok := ingress.Annotations[authSnippet]; ok {
								return fmt.Errorf(
									"invalid %s annotation on environment %s, route %s, ingress %s: %s",
									authSnippet, eName, rName, iName,
									"this annotation is restricted")
							}
							// configuration-snippet
							if annotation, ok := validate(ingress.Annotations, validSnippets,
								configurationSnippet); !ok {
								return fmt.Errorf(
									"invalid %s annotation on environment %s, route %s, ingress %s: %s",
									configurationSnippet, eName, rName, iName, annotation)
							}
							// modsecurity-snippet
							if _, ok := ingress.Annotations[modsecuritySnippet]; ok {
								return fmt.Errorf(
									"invalid %s annotation on environment %s, route %s, ingress %s: %s",
									modsecuritySnippet, eName, rName, iName,
									"this annotation is restricted")
							}
							// server-snippet
							if annotation, ok := validate(ingress.Annotations, validSnippets,
								serverSnippet); !ok {
								return fmt.Errorf(
									"invalid %s annotation on environment %s, route %s, ingress %s: %s",
									serverSnippet, eName, rName, iName, annotation)
							}
						}
					}
				}
			}
		}
		for _, routeMap := range l.ProductionRoutes.Active.Routes {
			for rName, lagoonRoutes := range routeMap {
				for _, lagoonRoute := range lagoonRoutes {
					for iName, ingress := range lagoonRoute.Ingresses {
						// auth-snippet
						if _, ok := ingress.Annotations[authSnippet]; ok {
							return fmt.Errorf(
								"invalid %s annotation on active environment, route %s, ingress %s: %s",
								authSnippet, rName, iName,
								"this annotation is restricted")
						}
						// configuration-snippet
						if annotation, ok := validate(ingress.Annotations, validSnippets,
							configurationSnippet); !ok {
							return fmt.Errorf(
								"invalid %s annotation on active environment, route %s, ingress %s: %s",
								configurationSnippet, rName, iName, annotation)
						}
						// modsecurity-snippet
						if _, ok := ingress.Annotations[modsecuritySnippet]; ok {
							return fmt.Errorf(
								"invalid %s annotation on active environment, route %s, ingress %s: %s",
								modsecuritySnippet, rName, iName,
								"this annotation is restricted")
						}
						// server-snippet
						if annotation, ok := validate(ingress.Annotations, validSnippets,
							serverSnippet); !ok {
							return fmt.Errorf(
								"invalid %s annotation on active environment, route %s, ingress %s: %s",
								serverSnippet, rName, iName, annotation)
						}
					}
				}
			}
		}
		for _, routeMap := range l.ProductionRoutes.Standby.Routes {
			for rName, lagoonRoutes := range routeMap {
				for _, lagoonRoute := range lagoonRoutes {
					for iName, ingress := range lagoonRoute.Ingresses {
						// auth-snippet
						if _, ok := ingress.Annotations[authSnippet]; ok {
							return fmt.Errorf(
								"invalid %s annotation on standby environment, route %s, ingress %s: %s",
								authSnippet, rName, iName,
								"this annotation is restricted")
						}
						// configuration-snippet
						if annotation, ok := validate(ingress.Annotations, validSnippets,
							configurationSnippet); !ok {
							return fmt.Errorf(
								"invalid %s annotation on standby environment, route %s, ingress %s: %s",
								configurationSnippet, rName, iName, annotation)
						}
						// modsecurity-snippet
						if _, ok := ingress.Annotations[modsecuritySnippet]; ok {
							return fmt.Errorf(
								"invalid %s annotation on standby environment, route %s, ingress %s: %s",
								modsecuritySnippet, rName, iName,
								"this annotation is restricted")
						}
						// server-snippet
						if annotation, ok := validate(ingress.Annotations, validSnippets,
							serverSnippet); !ok {
							return fmt.Errorf(
								"invalid %s annotation on standby environment, route %s, ingress %s: %s",
								serverSnippet, rName, iName, annotation)
						}
					}
				}
			}
		}
		return nil
	}
}
