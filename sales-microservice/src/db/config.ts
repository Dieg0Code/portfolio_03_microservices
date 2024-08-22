import { DynamoDB, Config } from "aws-sdk";

export function newDynamoDB(region: string): DynamoDB.DocumentClient {
    const sess = new Config({
        region: region
    });

    return new DynamoDB.DocumentClient(sess);
}