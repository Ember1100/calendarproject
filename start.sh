docker build -t gin-app .

ocker run --publish 8016:8016 --name gin-app  --rm gin-app