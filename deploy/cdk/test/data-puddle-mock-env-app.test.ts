import { Template } from 'aws-cdk-lib/assertions';
import { TestApp, TestDataPuddleMockCrmStack } from './cdk-test-helper';

test('DataPuddleMockCrmStackSnapshot', () => {
  const app = new TestApp();
  const stack = new TestDataPuddleMockCrmStack(app, 'DataPuddleMockCrmStack');

  const template = Template.fromStack(stack);
  expect(template.toJSON()).toMatchSnapshot();
});