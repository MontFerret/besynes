import { Disposable, free } from 'disposable-class';
import { Socket, socket } from 'zeromq';

export interface Settings {
    addr: string;
}

export class ExecutionService extends Disposable {
    @free({ call: 'close' })
    private __conn: Socket;

    constructor(settings: Settings) {
        super();

        this.__conn = socket('req');
        this.__conn.connect(settings.addr);
    }
}
