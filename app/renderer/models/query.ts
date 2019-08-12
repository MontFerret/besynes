export interface Query {
    id: string;
    name: string;
    text: string;
    params: {
        [key: string]: any;
    };
}
