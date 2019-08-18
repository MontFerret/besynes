import React from 'react';
import { IAppStore } from '../../stores/app';

const context = React.createContext({});

export interface StoreProviderProps extends React.Props<any> {
    store: IAppStore;
}

export function StoreProvider(props: StoreProviderProps): any {
    return (
        <context.Provider value={props.store}>
            {props.children}
        </context.Provider>
    );
}

export interface StoreConsumerProps {
    children: (store: any) => any;
}

export function StoreConsumer(props: StoreConsumerProps): any {
    return <context.Consumer>{props.children}</context.Consumer>;
}
