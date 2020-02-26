# Kubernetes Workshop

Welcome to the Kubernetes workshop! Please proceed with the steps below.

For the best workshop experience you should register your docker ID, install docker desktop, build and push your own image to the
docker registry. If for some reason you don't wish to install new software and create new accounts - it's ok too, there is an image
prepared for you to use. If you opt for not building your own container image, please skip the optional steps. 

### Prerequisites
1. Open your favorite terminal (Powershell, bash or gitbash)

### Install Docker Desktop (optional, but recommended)
For the best workshop experience, instal 
1. Create DockerID and install Docker Desktop: [https://www.docker.com/products/docker-desktop](https://www.docker.com/products/docker-desktop)
2. Start docker desktop

If there's a problem with docker installation, skip this step.
### Build and push docker image (optional, but recommended)
1. Clone this repository locally
2. Login to your DockerId via Docker Desktop (from your machine)
2. In the repo directory, run command: `docker build -t <yourdockerID name>/k8s-101:latest .`, this will build the container image
3. Push the container image with command: `docker push <yourdockerID name>/k8s-101:latest`
If there's a problem with docker installation or image building, skip this step.

### Deploy demo application to your Kubernetes namespace

1. In the terminal, go to directory in this repo: `infrastructure/k8s`
2. Open `deployment.yaml` file with text editor and replace `__YOUR_IMAGE_NAME__` with `<yourdockerID name>/k8s-101`. In case you skipped docker image building and pushing,
use this image instead: `tcentric/k8s-101`
3. Open `certificate.yaml` and replace `__FIRST_PART_OF_UID__` with first 8 characters of your UID key, which is displayed in https://kubernetes101.centric.engineering
4. Do the same in `ingress.yaml` file
5. Save updated file and create manifests in kubernetes with the command (you need to be in `infrastructure/k8s` directory): `kubectl create -f .`
7. Open your browser and go to https://`first 8 chars of your UID`.kubernetes101.centric.engineering, in a few minutes you should see your application live

### Scale application horizontally

1. While keeping the browser open, open another shell window
2. Scale application horizontally with the command: `kubectl scale deploy k8s101 --replicas 3`
3. Keep an eye on the browser window now
4. The application sends async HTTP requests to the backend retrieving the server(running in container) hostname
5. You should be able to see POD hostname changing now, this means that all HTTP requests load balanced across 3 replicas of your application
6. Scale your application down to one replica with command:   `kubectl scale deploy k8s101 --replicas 1`
7. Keep an eye on the browser window, hostnames should stop changing now


__You have successfully deployed and scaled an application in Kubernetes - wasn't it easy?__

## Thanks for attending the workshop!

Created by
[Tomas Adomavicius](mailto:tomas.adomavicius@centric.eu), 2020
