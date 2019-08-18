import { Socket } from 'electron-ipc-socket';
import { Disposable } from '../common/disposable';

export interface Settings {
    socket: Socket;
}

export class Application extends Disposable {
    private __socket: Socket;

    constructor(settings: Settings) {
        super();

        this.__socket = settings.socket;
        this.__socket.open();
    }
}
