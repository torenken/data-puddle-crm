import { Template } from 'aws-cdk-lib/assertions';
import { TestApp, TestCrmStack } from './cdk-test-helper';

test('DataPuddleCrmStackSnapshotTest', () => {
  const app = new TestApp();
  const stack = new TestCrmStack(app, 'TestDataPuddleCrmStack');

  const template = Template.fromStack(stack);

  template.hasResource('AWS::S3::Bucket', {});
  template.hasResource('AWS::SNS::Topic', {});

  template.hasResource('AWS::ApiGateway::RestApi', {});

  template.hasResource('AWS::Lambda::Function', {});

  expect(template.toJSON()).toMatchSnapshot();
});