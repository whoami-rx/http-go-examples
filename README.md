## Hello World


### Adapt function as handler

```
http.handleFunc(path, handler)
```

- path: route path
- handler: function that implements `Handler` interface.


### Start server 

```
http.ListenAndServe(port, handler)
```

- port: port to bind server(string)
- handler: should implement `Handler`