rmdir /s /q "%cd%/gen"

docker run --rm -v "%cd%:/work" uber/prototool:1.9.0 prototool format -f -w ./proto
docker run --rm -v "%cd%:/work" uber/prototool:1.9.0 prototool lint ./proto
docker run --rm -v "%cd%:/work" uber/prototool:1.9.0 prototool generate ./proto
