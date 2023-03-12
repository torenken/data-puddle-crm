## Create Infrastructure Project
[Projen](https://github.com/projen/projen) is used for the AWS CDK project generation.

### CDK Project Generation
In the /deploy/cdk directory, the AWS Cloudformation templates are created using the cdk tool. To generate the project, 
the Progen configuration ([.projenrc.js](.projenrc.js)) must first be copied to the /deploy/cdk directory.

Then the following shell commands must be executed

```shell
$ npx projen@latest --no-post
$ yarn (sometimes this has to be done several times)
$ yarn build 
```