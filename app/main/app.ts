import { Disposable, free } from 'disposable-class';
import { InboundRequest, Socket, Transport } from 'electron-ipc-socket';
import { ExecutionService } from './modules/execution/service';

export interface Settings {
    worker: string;
    transport: Transport;
}

export class Application extends Disposable {
    @free()
    private __socket: Socket;

    @free()
    private __execution: ExecutionService;

    constructor(settings: Settings) {
        super();

        this.__socket = new Socket(settings.transport, {
            timeout: 999999,
        });

        this.__execution = new ExecutionService({
            path: settings.worker,
            pub: 5051,
            sub: 5052,
        });

        this.__socket.open('@besynes');

        this.__socket.onRequest('/execute', (req: InboundRequest) => {
            return {};
        });
    }
}
