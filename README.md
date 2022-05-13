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

### validate-config-map-json

`lagoon-linter validate-config-map-json` allows you to validate a dump of configmaps from an existing kubernetes cluster containing Lagoon environments.
This is helpful when making changes to the linter where you want to check if it will cause build failures on existing Lagoon environments.

First get a dump of configmaps:

```
for cluster in abc1 xyz2; do
  # assuming kconfig switches kubectl contexts
  kconfig myname-$cluster && kubectl get configmap -Ao json > ~/download/lagoon-yml-audit/amazeeio-$cluster.cm.json;
done
```

Then run the linter over them.
It will automatically detect and validate `.lagoon.yml` configmaps only.

```
for file in ~/download/lagoon-yml-audit/*.json; do
  echo $file; ./lagoon-linter validate-config-map-json --config-map-json $file;
done
```
