# Edge deployment (The Pit) - Simulated using DigitalOcean for edge nodes
terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }
}

provider "digitalocean" {}

# Edge Nodes (Simulating K3s hardware at the mine pit)
resource "digitalocean_droplet" "edge_node" {
  count  = 2
  name   = "minelink-edge-node-${count.index + 1}"
  region = "nyc1"
  size   = "s-2vcpu-4gb" # Mimic 2-Pod constraint resource profile
  image  = "ubuntu-22-04-x64"

  ssh_keys = [var.ssh_fingerprint]

  provisioner "remote-exec" {
    inline = [
      "curl -sfL https://get.k3s.io | sh -s - --disable traefik",
      "# K3s setup complete"
    ]
  }
}
