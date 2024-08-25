import { DynamoDB } from "aws-sdk";

export function newDynamoDB(region: string): DynamoDB.DocumentClient {
    const config = {
        region: region,
        endpoint: "http://dynamodb-sales:8000",
    };

    return new DynamoDB.DocumentClient(config);
}