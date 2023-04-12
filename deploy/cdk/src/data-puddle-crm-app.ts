import { App, Tags } from 'aws-cdk-lib';
import { DataPuddleCrmStack } from './data-puddle-crm-stack';

const app = new App();

Tags.of(app).add('domain', 'data-puddle');
Tags.of(app).add('owner', 'torenken');

const technicalStakeholders = app.node.tryGetContext('technicalStakeholders');
const exportDataBucketName = app.node.tryGetContext('exportDataBucketName');

new DataPuddleCrmStack(app, 'DataPuddleCrmStack', {
  emailAddresses: technicalStakeholders,
  exportDataBucketName: exportDataBucketName,
});

app.synth();