# whatsmyip
Simple webserver written in go that returns ip address. Main goal is to get used to docker and travis.

## Getting started
Run 
```bash
docker run --rm -p 80:8080 leonnicolas/whatsmyip
```
or to run as a daemon that starts at reboot run 
```
docker run --restart unless-stopped -d -p 80:8080 leonnicolas/whatsmyip
```
## nginx as reverse Proxy
Server works behind __one__ nginx as reverse proxy. Use _X-Real-IP_ in nginx config:
```bash
location / {
	proxy_set_header Host $host;
	proxy_set_header X-Real-IP $remote_addr;
	proxy_pass http://localhost:8080;
}
