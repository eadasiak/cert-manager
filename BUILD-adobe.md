# Building cert-manager with RedHat's `go-toolset` for FIPS 140-2 compliance

## Intro

This repository is a fork of the official [cert-manager.io](https://github.com/jetstack/cert-manager.git) repository.  The changes introduced by Adobe are strictly for building cert-manager with FIPS 140-2 certified libraries.  This particular branch was created to utilize a host-based solution to the lack of a compliant crypto library in the standard version of golang.  RedHat provides a certified implementation with [`go-toolset`](https://developers.redhat.com/blog/2019/06/24/go-and-fips-140-2-on-red-hat-enterprise-linux/), however there are some important caveats:

- CentOS 7/RHEL 7 have an ancient version of golang in their `go-toolset`.  Version 1.10 is the latest available.
- CentOS 8/RHEL 8 has a more recent version (1.14 at the time of writing), but CentOS 8 has been deprecated.  A new program for free use of RHEL 8 is planned.  `cert-manager` release 1.1 is currently pinned to golang version 1.15, but compiles with 1.14.
- For `go-toolset` to be FIPS certified, the entire operating system has to be modified.  See [this](https://access.redhat.com/discussions/3487481)
- Once the host modifications have been made, remote access via SSH is impossible with CyberArk (you need a jump host)
- `cert-manager` can't be compiled in a container with this method, so there are several build dependencies outlined below.

## Host dependencies

These would apply just about anywhere the build host is provisioned, but worked in both an Azure VM and an IT Cloud VM:

- Bazel:  It's probably easiest to use [bazelisk](https://github.com/bazelbuild/bazelisk) instead of maintaining a version on the host.  Version 3.7.1 works (via `export USE_BAZEL_VERSION=3.7.1`)
- Python: `alternatives --install /usr/bin/python python /usr/bin/python3 99`
- Docker: CentOS/RHEL is pushing `podman` in version 8 instead of Docker. Use the official [Docker instructions](https://docs.docker.com/engine/install/centos/)
- Bash: The `/tmp` directory may be mounted as `noexec`.  Set your `TMPDIR` somewhere else.

The only source modifications to `cert-manager` modifications were:

- changing of the `go_version` in the `go_register_toolchains` rule in `/WORKSPACE` to use the host's installed version of golang instead of downloading one

## Execution

The following seems to work:

```bash
DOCKER_REGISTRY=${DOCKER_REGISTRY} APP_VERSION=${APP_VERSION} bazelisk run --stamp --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //build:server-images
```

This will load the resulting images in the build host's Docker daemon.  There's no build target for pushing the resulting containers from Bazel, so a separate step in the build process needs to be used to push the resulting images to the destination registry (and subsequently removing them from the local Docker daemon).

## Validating the binaries are FIPS-enabled

Either copy a binary from the compiled images or locate the binary in the bazel output directory (`bazel-out`).  The output from `go tool nm` will show the FIPS libraries included.

```bash
$ go tool nm webhook | grep fips
 26842c0 D crypto/internal/boring/fipstls.required
  690200 T crypto/tls.fipsCipherSuites
  690080 T crypto/tls.fipsCurvePreferences
 26376c0 D crypto/tls.fipsSupportedSignatureAlgorithms
```
