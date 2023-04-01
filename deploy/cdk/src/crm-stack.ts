import { Stack, StackProps } from 'aws-cdk-lib';
import { Construct } from 'constructs';
import { CrmEndpoint, DataBucket, TechnicalNotification } from './crm';

export interface CrmStackProps extends StackProps {
  readonly emailAddresses: string[];
  readonly exportDataBucketName: string;
}

export class CrmStack extends Stack {
  constructor(scope: Construct, id: string, props: CrmStackProps) {
    super(scope, id, props);

    const crmTechnicalNotification = new TechnicalNotification(this, 'CrmTechnicalNotification', {
      emailAddresses: props.emailAddresses,
    });

    const dataBucket = new DataBucket(this, 'CrmDataBucket', {
      bucketName: props.exportDataBucketName,
    });

    new CrmEndpoint(this, 'CrmEndpoint', {
      alarmNotification: crmTechnicalNotification,
      dataBucket,
    });
  }
}