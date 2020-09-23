#!/bin/bash

rm -r "$(pwd)/gen"

docker run --rm -v "$(pwd):/work" uber/prototool:1.9.0 prototool format -f -w ./proto
docker run --rm -v "$(pwd):/work" uber/prototool:1.9.0 prototool lint ./proto
docker run --rm -v "$(pwd):/work" uber/prototool:1.9.0 prototool generate ./proto
