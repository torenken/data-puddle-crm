import { CfnOutput } from 'aws-cdk-lib';
import { LambdaIntegration } from 'aws-cdk-lib/aws-apigateway';
import { IBucket } from 'aws-cdk-lib/aws-s3';
import { ITopic } from 'aws-cdk-lib/aws-sns';
import { Construct } from 'constructs';
import { Api } from './api';
import { CrmHandler } from './crm-handler';

export interface CrmEndpointProps {
  readonly alarmNotification: ITopic;
  readonly dataBucket: IBucket;
}

export class CrmEndpoint extends Construct {
  public readonly urlOutput: CfnOutput;

  constructor(scope: Construct, id: string, props: CrmEndpointProps) {
    super(scope, id);

    // RestApi
    const crmApi = new Api(this, 'CrmApi', {
      alarmNotification: props.alarmNotification,
    });

    const provideDataUrlFunc = new CrmHandler(this, 'ProvideDataUrlFunc', {
      serviceName: 'provideDataUrl',
      environment: {
        DATA_BUCKET_NAME: props.dataBucket.bucketName,
      },
    });
    props.dataBucket.grantRead(provideDataUrlFunc);

    const storeResource = crmApi.root.addResource('store');
    storeResource.addMethod('GET', new LambdaIntegration(provideDataUrlFunc));

    this.urlOutput = new CfnOutput(this, 'Url', {
      value: crmApi.url,
    });
  }
}