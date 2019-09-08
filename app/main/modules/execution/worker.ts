import { Disposable, free } from 'disposable-class';
import execa from 'execa';

export interface Settings {
    path: string;
    pub: number;
    sub: number;
}

export class WorkerProcess extends Disposable {
    @free()
    private readonly __settings: Settings;
    private readonly __process: execa.ExecaSyncReturnValue<string>;

    constructor(settings: Settings) {
        super();

        this.__settings = settings;
        this.__process = execa.sync(
            this.__settings.path,
            [`--sub=${this.__settings.sub}`, `--pub=${this.__settings.pub}`],
            {
                cleanup: true,
            },
        );

        console.log(this.__process.stdout);
    }
}
