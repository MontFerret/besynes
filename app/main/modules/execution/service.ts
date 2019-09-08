import { Disposable, free } from 'disposable-class';
import { Query } from './model';
import { Publisher } from './publisher';
import { Subscriber } from './subscriber';
import { WorkerProcess } from './worker';

export interface Settings {
    path: string;
    pub: number;
    sub: number;
}

export class ExecutionService extends Disposable {
    @free()
    private __publisher: Publisher;

    @free()
    private __subscriber: Subscriber;

    @free()
    private __worker: WorkerProcess;

    constructor(settings: Settings) {
        super();

        this.__publisher = new Publisher(`tcp://localhost:${settings.pub}`);
        this.__subscriber = new Subscriber(`tcp://localhost:${settings.sub}`);
        this.__worker = new WorkerProcess({
            path: settings.path,
            // Reverse logic here
            // Electron's pub -> Worker sub
            // Electron's sub -> Worker pub
            pub: settings.sub,
            sub: settings.pub,
        });
    }

    public async execute(query: Query): Promise<any> {
        this.__publisher.publish(query);
    }
}
