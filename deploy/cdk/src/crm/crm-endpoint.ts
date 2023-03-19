import { CfnOutput } from 'aws-cdk-lib';
import { MockIntegration } from 'aws-cdk-lib/aws-apigateway';
import { ITopic } from 'aws-cdk-lib/aws-sns';
import { Construct } from 'constructs';
import { Api } from './api';

export interface CrmEndpointProps {
  readonly alarmNotification: ITopic;
}

export class CrmEndpoint extends Construct {
  public readonly urlOutput: CfnOutput;

  constructor(scope: Construct, id: string, props: CrmEndpointProps) {
    super(scope, id);

    // RestApi
    const crmApi = new Api(this, 'CrmApi', {
      alarmNotification: props.alarmNotification,
    });

    this.urlOutput = new CfnOutput(this, 'Url', {
      value: crmApi.url,
    });

    crmApi.root.addMethod('GET', new MockIntegration());

  }

}