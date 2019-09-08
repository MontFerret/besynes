import { Disposable, free, protect } from 'disposable-class';
import { Socket, socket } from 'zeromq';
import { QueryResult } from './model';

export class Subscriber extends Disposable {
    @free({ call: 'close' })
    private __conn: Socket;

    constructor(addr: string) {
        super();

        this.__conn = socket('pub');
        this.__conn.connect(addr);
    }

    @protect()
    public subscribe(handler: (result: QueryResult) => void): void {
        this.__conn.on('message', handler);
    }
}
