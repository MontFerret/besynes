import { Disposable, free } from 'disposable-class';
import { Socket } from 'electron-ipc-socket';

export interface Settings {
    socket: Socket;
}

export class Application extends Disposable {
    @free()
    private __socket: Socket;

    constructor(settings: Settings) {
        super();

        this.__socket = settings.socket;
        this.__socket.open('@besynes');
    }
}
