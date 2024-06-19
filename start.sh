docker build -t gin-app .

ocker run --privileged=true 8016:8016 --name gin-app  --rm gin-app