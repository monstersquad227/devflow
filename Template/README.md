### Java17Template
The content in the template file is the pipeline content of Jenkins, and the interface request parameters are as follows.

#### body

```json
{
    "gitlab_name": "xxx-gateway",
    "deployment_name": "gateway",
    "build_template_id": "101",
    "branch": "master",
    "gitlab_repo": "git@gitlab.xxxxxxxxxxx.com:cdd_java/xxx-gateway.git",
    "environment_unique": "dev",
    "harbor_url": "harbor.xxxxxxxxxxx.com",
    "short_id": "7aeeeb1f",
    "command": "mvn clean package"
}
```