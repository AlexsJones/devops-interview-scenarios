## Scenario 0

This project contains components required to complete a set of exercises.


1. Familiarise with the `blacklist-example` code.
   Assume it has been pushed to `tibbar/blacklist-example` and `docker pull tibbar/blacklist-example` will retrieve it.

2. Pull the image and run the program - Use curl to test the endpoint to get a feel for the response.

3. Create the minimal amount of Kuberentes manifests to deploy locally to `minikube` or `docker-for-desktop` and port-forward to test it works (Recommendation here is a Deployment/Service).
   - Port-forward and test the locally forwarded service.

4. Change the code within `blacklist-example/main.go` to allow for a server port set by an environmental variable.
   - Change the necessary parameters to set the port to `9093`
   - Push the docker image (If necessary to a registry user you own - and make the necessary local adjustments).
   - Test changes as per step 3

5. Create a Helm project and convert the Kubernetes resources to a helm compatible setup.
   - Set the port, image and default command within the Helm values.

6. Extend the blacklist API to support GET verb and to return a boolean as a result of checking a parameterised hostname.
   e.g. `piratebay.io` will respond with a 200 and the payload parameter of `{ value: "true" }` or a similar format.
   - Push the docker image (If necessary to a registry user you own - and make the necessary local adjustments).
   - Test changes as per step 3
