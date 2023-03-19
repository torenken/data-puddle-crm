import { App, Tags } from 'aws-cdk-lib';
import { CrmStack } from './crm-stack';

const app = new App();

Tags.of(app).add('domain', 'data-puddle');
Tags.of(app).add('owner', 'torenken');

const technicalStakeholders = app.node.tryGetContext('technicalStakeholders');

new CrmStack(app, 'DataPuddleMockCrmStack', {
  emailAddresses: technicalStakeholders,
});

app.synth();