# terraform-provider-circle3

## Build local:

Run following commands to build the project.

```
go mod tidy
go build
mv terr
mv terraform-provider-circle3 ~/.terraform.d/plugins/{dir}/tf/circle3/0.1/{system}/terraform-provider-circle3
```

And test it:

```
cd examlple/...
terraform init

terraform apply
```

## Compile project
```
go mod tidy
go build
export OS_ARCH="$(go env GOHOSTOS)_$(go env GOHOSTARCH)"
mv terraform-provider-circle3 ~/.terraform.d/plugins/bmeik/tf/circle3/0.1/$OS_ARCH/terraform-provider-circle3 # create directory, if needed
```