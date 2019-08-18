export abstract class Disposable {
    private __isDisposed: boolean;

    constructor() {
        this.__isDisposed = false;
    }

    public get isDisposed(): boolean {
        return this.__isDisposed;
    }

    public dispose(): void {
        if (!this.__isDisposed) {
            this.__isDisposed = true;
        }
    }
}
