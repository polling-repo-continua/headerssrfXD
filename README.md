# headerssrfXD
Scan ssrf on headers. Inspired by the tool https://github.com/m4ll0k/Bug-Bounty-Toolz/blob/master/ssrf.py. It adds headers like X-Forwarded-For and Proxy-Host and set it to your server or to your burp collaborator
If you got a pingback from your server, it means the url is vulnerable. It takes input from stdin allowing you to pipe urls easily

# How to install
```go get github.com/noobexploiter/headerssrfXD```

# How to use
You need to specify your server address using the -c tag. You can specify threads to run using the -t tag. <br>
Example:
```cat urls.txt | headerssrfXD -c myserver.com -t 20```

# Additional Info
I might update this soon if i got any new idea to improve this
