# Supernatural

## ECR Installation and Use

To pull and run the application via your terminal, execute the following commands in order:
1. First to log in your aws account(if needed)

```bash
aws config
```
2. Than Authenticate Docker with AWS ECR Log in to the ECR registry to grant Docker permission to pull the image

```bash
aws ecr get-login-password --region <aws-region> | docker login --username AWS --password-stdin <aws-organization-id>.dkr.ecr.<aws-region>.amazonaws.com
```

3. Than you can pull image from ECR
```bash
docker pull <aws-organization-id>.dkr.ecr.<aws-region>.amazonaws.com/<repository-name>:<image tag>>
```

Note: if you what to be unsure that image on ur device run a command:

```bash
docker image ls
```

Finally can easily run that image

```bash
docker run <aws-organization-id>.dkr.ecr.<aws-region>.amazonaws.com/<repository-name>:<image tag>
```
