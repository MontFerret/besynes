export interface Query {
    id: string;
    text: string;
    params?: { [name: string]: any };
    cdp: string;
}

export interface Job {}

export interface QueryResult {
    jobID: string;
    queryID: string;
    timestamp: Date;
    status: number;
    error?: string;
    data?: string;
}
