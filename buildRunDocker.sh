docker build --rm -f Dockerfile -t studiguide/studiguide_appservice . --build-arg ionicproduction=--prod
docker run -it --rm -p 8080:8080 studiguide/studiguide_appservice:latest
