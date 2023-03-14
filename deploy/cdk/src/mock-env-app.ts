import { App } from 'aws-cdk-lib';
import { CrmStack } from './crm-stack';

const app = new App();

new CrmStack(app, 'DataPuddleMockCrmStack', {});

app.synth();