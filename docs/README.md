# deploy-ovh — docs

**OVH deploy.** Provision an OVH Public Cloud instance via the `openstack` CLI.

## Install

```bash
togo install togo-framework/deploy-ovh
```

Registers on the [`deploy`](https://github.com/togo-framework/deploy) base; select it with **deploy.provider in togo.yaml (or DEPLOY_PROVIDER)**, then use **`togo deploy`**.

## Interface

`Deployer` — `Provision`/`Deploy`/`Destroy`/`Status` over a `Spec{App,Dir,BuildCmd,Host,User,Image,Region,Domain}` built from your `togo.yaml`.

## Usage & notes

Requires the `openstack` CLI configured for your OVH project. Creates a server running `spec.Image`.

## Example

```bash
togo deploy --provider ovh --dry-run   # preview the plan
togo deploy --provider ovh
```

## Links

- [OVH Public Cloud](https://help.ovhcloud.com/csm/en-public-cloud-compute)
- [Marketplace](https://to-go.dev/marketplace)
- [Source](https://github.com/togo-framework/deploy-ovh)
