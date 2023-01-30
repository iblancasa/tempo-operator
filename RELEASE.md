# Releasing the Tempo Operator for Kubernetes

## Releasing

Steps to release a new version of the Tempo Operator:


1. Change the `versions.txt `so that it lists the target version of the Tempo (if it is required). **Don't touch the operator version**: it will be changed automatically in the next step.

2. Run `OPERATOR_VERSION=2.0.0 make prepare-release`, using the operator version that will be released.

3. Prepare a changelog since last release.

4. Commit the changes and create a pull request:

   ```sh
   git commit -sm "Preparing release v2.0.0"
   ```
