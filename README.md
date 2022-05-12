# Lagoon Linter

Lint `.lagoon.yml` for validity.

## Profiles

`lagoon-linter` has the concept of profiles for validating `.lagoon.yml` files.
Different profiles apply different linting checks.

| Name       | Description                                                                                   |
| ---        | ---                                                                                           |
| required   | Checks for invalid or restricted content in `.lagoon.yml`. These will fail a Lagoon build.    |
| deprecated | Checks for deprecated content in `.lagoon.yml`. These will print a warning in a Lagoon build. |

If no profile is specified, the `required` profile will run by default.

## Linters

Currently implemented linters.

| Name            | Profile    | Description                                                                                                                                                                 |
| ---             | ---        | ---                                                                                                                                                                         |
| RouteAnnotation | required   | Validates Lagoon Route / Kubernetes Ingress annotations. See the documentation [here](https://docs.lagoon.sh/using-lagoon-the-basics/lagoon-yml/#restrictions) for details. |
| MonitoringURLs  | deprecated | Checks for the presence of `monitoring_urls`.                                                                                                                               |

## Usage

Run `lagoon-linter` in the directory containing your `.lagoon.yml`.
See `lagoon-linter --help` for options.
