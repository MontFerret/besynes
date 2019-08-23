import { Socket } from 'electron-ipc-socket';
import { IQuery } from '../../common/models/query';

export class Api {
    private __socket: Socket;

    constructor(socket: Socket) {
        this.__socket = socket;
        this.__socket.open('@besynes');
    }

    public async executeQuery(q: IQuery): Promise<any> {
        return this.__socket.request('query/execute', q);
    }
}
