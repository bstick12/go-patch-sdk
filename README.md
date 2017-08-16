## Update Kubo Release

Modify the `glide.yaml` to update the version number for kubo-deployment

## Embed the Kubo resources

```bash
go-bindata -pkg kubo -prefix vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ -o kubo/kubo.go vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/...
```


