import React from 'react';
import { IndexScreen } from './screens/index';

export interface Props {}

export class AppComponent extends React.PureComponent<Props> {
    public render(): any {
        return <IndexScreen />;
    }
}
