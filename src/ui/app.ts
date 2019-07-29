import { History } from 'history';
import React from 'react';
import { create } from './components/index';

export default class Application {
    private readonly __history: History;

    constructor(history: History) {
        this.__history = history;
    }

    public createElement(): React.ReactElement<any> {
        return create(this.__history);
    }
}
