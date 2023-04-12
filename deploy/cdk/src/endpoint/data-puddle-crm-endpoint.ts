import { CfnOutput } from 'aws-cdk-lib';
import { LambdaIntegration } from 'aws-cdk-lib/aws-apigateway';
import { IBucket } from 'aws-cdk-lib/aws-s3';
import { ITopic } from 'aws-cdk-lib/aws-sns';
import { Construct } from 'constructs';
import { DataPuddleCrnApi } from './data-puddle-crn-api';
import { DataPuddleCrmHandler } from '../data-puddle-crm-handler';

export interface CrmEndpointProps {
  readonly alarmNotification: ITopic;
  readonly dataBucket: IBucket;
}

export class DataPuddleCrmEndpoint extends Construct {
  public readonly urlOutput: CfnOutput;

  constructor(scope: Construct, id: string, props: CrmEndpointProps) {
    super(scope, id);

    // RestApi
    const crmApi = new DataPuddleCrnApi(this, 'DataPuddleCrnApi', {
      alarmNotification: props.alarmNotification,
    });

    const createAgreementFunc = new DataPuddleCrmHandler(this, 'CreateAgreementFunc', {
      serviceName: 'create-agreement',
      environment: {
        DATA_BUCKET_NAME: props.dataBucket.bucketName,
      },
    });
    props.dataBucket.grantWrite(createAgreementFunc);

    const agreementResource = crmApi.root.addResource('agreement');

    agreementResource.addMethod('POST', new LambdaIntegration(createAgreementFunc), {
      apiKeyRequired: true,
    });

    this.urlOutput = new CfnOutput(this, 'Url', {
      value: crmApi.url,
    });
  }
}