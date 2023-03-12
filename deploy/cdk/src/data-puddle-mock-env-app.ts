import { App } from 'aws-cdk-lib';
import { DataPuddleMockCrmStack } from './data-puddle-mock-crm-stack';

const app = new App();

new DataPuddleMockCrmStack(app, 'DataPuddleMockCrmStack', {});

app.synth();