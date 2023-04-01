import { App, Tags } from 'aws-cdk-lib';
import { CrmStack } from './crm-stack';

const app = new App();

Tags.of(app).add('domain', 'data-puddle');
Tags.of(app).add('owner', 'torenken');

const technicalStakeholders = app.node.tryGetContext('technicalStakeholders');
const exportDataBucketName = app.node.tryGetContext('exportDataBucketName');

new CrmStack(app, 'DataPuddleMockCrmStack', {
  emailAddresses: technicalStakeholders,
  exportDataBucketName,
});

app.synth();