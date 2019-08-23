import { ipcRenderer } from 'electron';
import { Socket } from 'electron-ipc-socket';
import { Instance } from 'mobx-state-tree';
import React from 'react';
import { Api } from './api/api';
import { create } from './components/index';
import { AppEnv } from './env/env';
import { AppStore } from './stores/app';

export default class Application {
    private readonly __api: Api;
    private readonly __store: Instance<typeof AppStore>;

    constructor() {
        this.__api = new Api(new Socket(ipcRenderer));
        this.__store = AppStore.create({}, {
            api: this.__api,
        } as AppEnv);
    }

    public createElement(): React.ReactElement<any> {
        return create(this.__store);
    }
}
