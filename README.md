# DigitalOcean Kubernetes Challenge

## Setup

Kubernetes cluster is provisioned using `terraform`, to get this up and running, use the following commands. This relies on you having an exported `DIGITALOCEAN_TOKEN` within your environment variable for access.

    cd terraform && terraform apply --auto-approve

This will provision the cluster with a single worker node.

Retrieve the `.kubeconfig` file using

    doctl kubernetes cluster kubeconfig save <cluster-id>
    # You can get this full command from the DigitalOcean Kubernetes UI.
