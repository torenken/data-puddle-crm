import { App } from 'aws-cdk-lib';
import { Construct } from 'constructs';
import { CrmStack } from '../src';

export class TestApp extends App {
  constructor() {
    super({
      context: {
        'aws:cdk:bundling-stacks': ['NoStack'], //disable bundling lambda (asset), by using dummy stack-name (=> reduce the unit-test-time. jest-booster)
        '@aws-cdk/core:newStyleStackSynthesis': 'true',
      },
    });
  }
}

export class TestDataPuddleMockCrmStack extends CrmStack {
  constructor(scope: Construct, id: string) {
    super(scope, id, { });
  }
}