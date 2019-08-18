import { TransportOutput } from 'electron-ipc-socket';

export class DynamicOutput implements TransportOutput {
    private __renderer: TransportOutput;

    constructor(renderer: TransportOutput) {
        this.__renderer = renderer;
    }

    public setRenderer(output: TransportOutput): void {
        this.__renderer = output;
    }

    public send(channel: string | symbol, ...args: any[]): void {
        this.__renderer.send(channel, ...args);
    }
}
