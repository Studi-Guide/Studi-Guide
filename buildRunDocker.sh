docker build --rm -f Dockerfile -t studiguide/studiguide_appservice .
docker run -it --rm -p 8080:8080 studiguide/studiguide_appservice:latest