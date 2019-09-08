import { Disposable, free, protect } from 'disposable-class';
import { Socket, socket } from 'zeromq';
import { Query } from './model';

export class Publisher extends Disposable {
    @free({ call: 'close' })
    private __conn: Socket;

    constructor(addr: string) {
        super();

        this.__conn = socket('pub');
        this.__conn.bindSync(addr);
    }

    @protect()
    public publish(q: Query): void {
        this.__conn.send(JSON.stringify(q));
    }
}
