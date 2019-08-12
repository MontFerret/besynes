import React from 'react';
import { AppComponent } from './components/index';

export default class Application {
    public createElement(): React.ReactElement<any> {
        return React.createElement(AppComponent);
    }
}
