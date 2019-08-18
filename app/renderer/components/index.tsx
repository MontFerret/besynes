import React from 'react';
import { IAppStore } from '../stores/app';
import { StoreProvider } from './common/context';
import { IndexScreen } from './screens/index';

export interface Props {
    store: IAppStore;
}

export class AppComponent extends React.PureComponent<Props> {
    public render(): any {
        return (
            <StoreProvider store={this.props.store}>
                <IndexScreen />;
            </StoreProvider>
        );
    }
}

export function create(store: IAppStore): any {
    return <AppComponent store={store} />;
}
