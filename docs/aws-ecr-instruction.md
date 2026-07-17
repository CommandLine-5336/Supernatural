# Supernatural
## ECR installation & use
* To pull and run the application via your terminal, execute the following commands in order:
1. First to log in your aws account(if needed)
```bash
aws config
```
2. Than Authenticate Docker with AWS ECR Log in to the ECR registry to grant Docker permission to pull the image
```bash
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 704427427594.dkr.ecr.us-east-1.amazonaws.com
```
3. Than you can pull image from ECR
```bash
docker pull 704427427594.dkr.ecr.us-east-1.amazonaws.com/supernatural:<image tag>>
```
Note: if you what to be unsure that image on ur device run a command:
```bash
docker image ls
```
* Finaly can easily run that image
```bash
docker run 704427427594.dkr.ecr.us-east-1.amazonaws.com/supernatural:<image tag>
```
