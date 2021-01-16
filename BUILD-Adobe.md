# Building cert-manager with BoringSSL for FIPS 140-2 compliance

## Intro

This repository is a fork of the official [cert-manager.io](https://github.com/jetstack/cert-manager.git) repository.  The changes introduced by Adobe are strictly for building cert-manager with FIPS 140-2 compliant libraries.  They loosely follow the guidance from the BoringSSL fork of golang when compiling with Bazel [here](https://github.com/golang/go/tree/dev.boringcrypto/misc/boring#building-from-bazel).


Unlike the RedHat FIPS-compliant `go-toolset`, compiling cert-manager with BoringSSL means there are no kernel dependencies on FIPS libraries in order to build it.  Thus, it can be compiled using the official Bazel image in a container from any generic build machine.

A quick list of the modifications made are:

- addition of a `go_download_sdk` rule in `WORKSPACE` to download the BoringSSL golang fork
- explicitly set `pure = "off"` on all `go_binary` targets so that `cgo` is enabled

## Execution

Based on the Basel [documentation](https://docs.bazel.build/versions/3.7.0/bazel-container.html), the following can be incorporated into a build job to use a container for building cert-manager:

```bash
docker run -e DOCKER_REGISTRY=${DOCKER_REGISTRY} -e APP_VERSION=${APP_VERSION} -v $(pwd):/src/workspace -v /tmp/build_output:/tmp/build_output -v /var/run/docker.sock:/var/run/docker.sock -w /src/workspace l.gcr.io/google/bazel:latest --output_user_root=/tmp/build_output run --stamp --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //build:server-images
```

This will load the resulting images in the build host's Docker daemon.  There's no build target for pushing the resulting containers from Bazel, so a separate step in the build process needs to be used to push the resulting images (and remove them from the local Docker daemon).

## Validating the binaries are FIPS-enabled

Either copy a binary from the compiled images or locate the binary in the output of /tmp/build_output.  The output from `go tool nm` will show the BoringSSL libraries included.

```bash
$ go tool nm webhook | grep boring
 241ecc0 D crypto/internal/boring..inittask
 24d40e0 D crypto/internal/boring/fipstls.required
  5f37a0 T crypto/internal/boring/sig.StandardCrypto
 1a3d8d8 r crypto/internal/boring/sig.StandardCrypto.args_stackmap
```
