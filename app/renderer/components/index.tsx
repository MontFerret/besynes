import { Provider as MobXProvider } from 'mobx-react';
import { Instance } from 'mobx-state-tree';
import React from 'react';
import { AppStore } from '../stores/app';
import { IndexScreen } from './screens/index';

export interface Props {
    store: Instance<typeof AppStore>;
}

export class AppComponent extends React.PureComponent<Props> {
    public render(): any {
        return (
            <MobXProvider store={this.props.store}>
                <IndexScreen />;
            </MobXProvider>
        );
    }
}

export function create(store: Instance<typeof AppStore>): any {
    return <AppComponent store={store} />;
}
