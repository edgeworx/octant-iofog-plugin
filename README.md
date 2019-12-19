![iofog-logo](images/iofog-logo.png?raw=true "iofog logo")

# Octant ioFog Plugin

>This is an [Octant](https://octant.dev/) plugin for [Eclipse ioFog](http://iofog.org) that provides visibility into an 
>Edge deployment of ioFog from a Kubernetes cluster.

_**Note**: The plugin currently only supports Octant [release-0.9.0](https://github.com/vmware-tanzu/octant/tree/release-0.9)_

[Eclipse ioFog](http://iofog.org) is an Edge Native Cloud platform, that supports the deployment, orchestration, 
management and monitoring of containerized microservice applications to the Edge. ioFog supports deployment onto bare 
metal, VMs, all major Cloud platforms and all major Kubernetes distributions.

![iofog-ecn-viewer](images/ecn-viewer.png?raw=true "iofog Viewer")

When [deployed onto a Kubernetes cluster](https://iofog.org/docs/1.3.0/remote-deployment/prepare-your-kubernetes-cluster.html), 
ioFog provides a seamless extension of Kubernetes to the Edge, enabling microservices to be scheduled from, for example, 
a cloud based Kubernetes cluster to an Nvidia Nano edge device running ioFog. You are able to use your existing tooling 
such as `kubectl` or `helm` to manage and deploy microservices from the Cloud to the Edge. 

## Installation

- Follow instructions for how to use [iofogctl](https://iofog.org/docs/1.3.0/iofogctl/usage.html#iofog-unified-command-line-interface)
to [deploy iofog](https://iofog.org/docs/1.3.0/remote-deployment/prepare-your-kubernetes-cluster.html) to your Kubernetes cluster. 
If you prefer, you can also deploy ioFog using [Helm](https://iofog.org/docs/1.3.0/tools/how-to-helm.html)
- Follow the instructions for [installing Octant](https://github.com/vmware-tanzu/octant#installation)

### Building the plugin 

Ensure you have all the needed dependencies locally by running `make dep`. 

Run `make build install` to build the plugin and copy the binary to Octant's plugins folder.
                                                             
## Screenshots

Once you have installed the plugin, you should (re)start Octant. Assuming you have deployed ioFog to your Kubernetes
cluster, you will now be able to see the ioFog components of Controller, Connector and Agent(s) in a new tree structure 
in Octant's nav menu:

![octant-iofog-viewer](images/octant-iofog-viewer.png?raw=true "Octant ioFog Viewer")

**Note**: Currently, Octant has a limitation of which namespaces can be used. You must select the 'k8s' namespace in
the drop down menu before you will be able to see the ioFog components.

## Uninstall

Run the following command to remove the plugin:

```
rm -f ~/.config/octant/plugins/octant-iofog
```