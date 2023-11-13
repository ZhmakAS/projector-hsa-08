# projector-hsa-08 Web Servers 


## How to launch

Execute `docker compose up --build` and that's all, after that service can be accessed on http://localhost:8080 (nginx)

Open in browser:

```
http://localhost:8080/assets/image.png
```

You should see the image in the browser and corresponding requests logs from server; 

On the third request you shouldn`t see any request logs, meaning that nginx cached the image;

```
projector-hsa-08-proxy-1      | 2023/11/13 23:03:57 [notice] 1#1: start cache loader process 45
projector-hsa-08-go-server-1  | 2023/11/13 23:04:01 main.go:53: 192.168.64.3:42416 GET /assets/image.png
projector-hsa-08-proxy-1      | 192.168.64.1 - - [13/Nov/2023:23:04:01 +0000] "GET /assets/image.png HTTP/1.1" 304 0 "-" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36" "-"
projector-hsa-08-go-server-1  | 2023/11/13 23:04:03 main.go:53: 192.168.64.3:42432 GET /assets/image.png
projector-hsa-08-proxy-1      | 192.168.64.1 - - [13/Nov/2023:23:04:03 +0000] "GET /assets/image.png HTTP/1.1" 304 0 "-" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36" "-"
projector-hsa-08-proxy-1      | 192.168.64.1 - - [13/Nov/2023:23:04:06 +0000] "GET /assets/image.png HTTP/1.1" 304 0 "-" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36" "-"
projector-hsa-08-proxy-1      | 192.168.64.1 - - [13/Nov/2023:23:04:07 +0000] "GET /assets/image.png HTTP/1.1" 304 0 "-" "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36" "-"
```


To reset the cache you can manually rename the image in the `assets` folder and call:

```
curl http://localhost:8080/assets/image.png --header 'Cache-Purge: true' -I
```

As a result you should see `BYPASS` value in `Cache` header and corresponding request log to the server;


```
HTTP/1.1 200 OK
Server: nginx/1.25.2
Date: Mon, 13 Nov 2023 23:04:54 GMT
Content-Type: image/png
Content-Length: 3468
Connection: keep-alive
Last-Modified: Mon, 13 Nov 2023 13:30:15 GMT
Cache: BYPASS
Accept-Ranges: bytes
```
