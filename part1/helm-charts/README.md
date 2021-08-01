## Helm charts

We use this directory to store example of helm chart - home-assessment. 

### Requirements

| Name  | Version |
| ------------- | ------------- |
| helm  | >=3.5  |


We assume, that you have configured aws cli and logged in to correct aws account(using aws cli).

### Usage

After you have created s3 in part1/terraform directory, we can upload test helm chart to this bucket.

1. To install the helm-s3 plugin on your client machine, run the following command: `helm plugin install https://github.com/hypnoglow/helm-s3.git`.

2. To initialize the target folder as a Helm repository, use the following command: `helm s3 init s3://home-helm-repository/stable/home-assessment/`. (Assume that you use default variable from terraform configuration and your bucket name and path has the same value). Make sure that you have set AWS_REGION env variable.

3. To add the target repository alias to the Helm client machine, use the following command: `helm repo add home-assessment s3://home-helm-repository/stable/home-assessment`.

4. To package the chart that you created or cloned, use the following command: `helm package home-assessment`.

5. To upload the local package to the Helm repository in Amazon S3, run the following command: `HELM_S3_MODE=3 helm s3 push home-assessment-0.1.0.tgz home-assessment`.

6. To confirm that the chart appears both locally and in the Amazon S3 Helm repository, run the following command: `helm search repo home-assessment`.

More details you can find on this guide from (AWS)[https://docs.aws.amazon.com/prescriptive-guidance/latest/patterns/set-up-a-helm-v3-chart-repository-in-amazon-s3.html].