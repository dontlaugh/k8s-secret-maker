# k8s-secret-maker

Small template util to create yaml files for k8s secrets. Values are provided as
strings but base64 encoded internally. Prints yaml to stdout.

Use like this:

```
k8s-secret-maker SECRETNAME secretkey1 secretvalue1 secretkey2 secretvalue2 ...
```

## Install

You need Go installed, and you need GOPATH/bin on your PATH.

```
go install -u github.com/dontlaugh/k8s-secret-maker
```

Check your environment with `go env`.
