#!/usr/bin/env sh
set -eu

LINTER=./dist/lagoon-linter_linux_amd64_v1/lagoon-linter
chmod +x "$LINTER"

# profile: required
for lagoonyml in ./internal/lagoonyml/required/testdata/valid.*.yml; do
	echo "$lagoonyml"
	rm -f .lagoon.yml
	$LINTER validate --lagoon-yaml="$lagoonyml" # explicitly validate $lagoonyml
	ln -fs "$lagoonyml" .lagoon.yml
	$LINTER # implicitly validate .lagoon.yml
done
for lagoonyml in ./internal/lagoonyml/required/testdata/invalid.*.yml; do
	echo "$lagoonyml"
	rm -f .lagoon.yml
	$LINTER validate --lagoon-yaml="$lagoonyml" && exit 1
	ln -fs "$lagoonyml" .lagoon.yml
	$LINTER && exit 1
done

# profile: deprecated
for lagoonyml in ./internal/lagoonyml/deprecated/testdata/valid.*.yml; do
	echo "$lagoonyml"
	rm -f .lagoon.yml
	$LINTER validate --profile=deprecated --lagoon-yaml="$lagoonyml"
	ln -fs "$lagoonyml" .lagoon.yml
	$LINTER validate --profile=deprecated
done
for lagoonyml in ./internal/lagoonyml/deprecated/testdata/invalid.*.yml; do
	echo "$lagoonyml"
	rm -f .lagoon.yml
	$LINTER validate --profile=deprecated --lagoon-yaml="$lagoonyml" && exit 1
	ln -fs "$lagoonyml" .lagoon.yml
	$LINTER validate --profile=deprecated && exit 1
done

exit 0
