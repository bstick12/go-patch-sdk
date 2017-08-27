## Update Kubo Release

Modify the `glide.yaml` to update the version number for kubo-deployment

## Embed the Kubo resources

```bash
go-bindata -pkg adapter -prefix vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ -o adapter/kubo.go vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/...
```


