encore build docker --config=./infra-config.json url-shortener:v1.0

docker run -e PORT=8081 -p 8081:8081 url-shortener:v1.0