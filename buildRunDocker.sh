docker build --rm -f Dockerfile -t studiguide/studiguide_appservice . --build-arg ionicproduction=--configuration=docker
docker run -it --rm -p 8080:8080 studiguide/studiguide_appservice:latest
