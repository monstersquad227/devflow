### Java17Template

The content in the template file is the pipeline content of Jenkins, and the interface request parameters are as follows.

#### body

```json
{
    "gitlab_name": "xxx-gateway",
    "deployment_name": "gateway",
    "task_id": "101",
    "branch": "master",
    "gitlab_repo": "git@gitlab.xxxxxxxxxxx.com:cdd_java/xxx-gateway.git",
    "environment_unique": "dev",
    "harbor_url": "harbor.xxxxxxxxxxx.com",
    "short_id": "7aeeeb1f",
    "command": "mvn clean package"
}
```

### VueTemplate

The content in the template file is the pipeline content of Jenkins, and the interface request parameters are as follows.

#### body

```json
{
    "gitlab_name": "xxx-gateway",
    "deployment_name": "gateway",
    "build_template_id": "102",
    "branch": "master",
    "gitlab_repo": "git@gitlab.xxxxxxxxxxx.com:cdd_java/xxx-gateway.git",
    "environment_unique": "dev",
    "harbor_url": "harbor.xxxxxxxxxxx.com",
    "short_id": "7aeeeb1f",
    "command": "npm install && npm run build:prod"
}
```

### Dotnet5Template

The content in the template file is the pipeline content of Jenkins, and the interface request parameters are as follows.

#### body

```json
{
    "gitlab_name": "restful",
    "deployment_name": "restfuladmin",
    "build_template_id": "104",
    "branch": "release_prod",
    "depend_branch": "release_20220418",
    "gitlab_repo": "git@gitlab.xxxxxxx.com:cdd/restful.git",
    "depend_gitlab_repo": "git@gitlab.xxxxxxx.com:mojory/commonlibs.git",
    "project_build_path": "XXX.AdminRestful/",
    "project_package_name": "XXX.AdminRestful",
    "environment_unique": "dev",
    "harbor_url": "harbor.xxxxxxxxxxx.com",
    "short_id": "4c2afdc1",
    "command": "dotnet restore && dotnet build && dotnet publish -c Debug -o out"
}
```

### Dotnet2Template

The content in the template file is the pipeline content of Jenkins, and the interface request parameters are as follows.

#### body

```json
{
  "gitlab_name": "qmessage",
  "deployment_name": "qmessage",
  "build_template_id": "105",
  "branch": "release_prod",
  "depend_branch": "release",
  "gitlab_repo": "git@gitlab.xxxxxxx.com:mojory/qmessage.git",
  "depend_gitlab_repo": "git@gitlab.xxxxxxx.com:mojory/commonlibs.git",
  "project_build_path": "Mojory.QMessage.Service/",
  "project_package_name": "Mojory.QMessage.Service",
  "environment_unique": "dev",
  "harbor_url": "harbor.xxxxxxxx.com",
  "short_id": "4c2afdc1",
  "command": "dotnet restore && dotnet build && dotnet publish -c Debug -o out"
}
```