import { Stack, StackProps } from 'aws-cdk-lib';
import { Construct } from 'constructs';
import { CrmEndpoint, DataBucket, TechnicalNotification } from './crm';

export interface CrmStackProps extends StackProps {
  readonly emailAddresses: string[];
}

export class CrmStack extends Stack {
  constructor(scope: Construct, id: string, props: CrmStackProps) {
    super(scope, id, props);

    const crmTechnicalNotification = new TechnicalNotification(this, 'CrmTechnicalNotification', {
      emailAddresses: props.emailAddresses,
    });

    new DataBucket(this, 'CrmDataBucket', {
      bucketName: 'torenken-data-puddle-crm-data',
    });

    new CrmEndpoint(this, 'CrmEndpoint', {
      alarmNotification: crmTechnicalNotification,
    });
  }
}