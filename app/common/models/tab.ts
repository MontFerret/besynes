import { Instance, types } from 'mobx-state-tree';
import { Query } from './query';

export const Tab = types.compose(
    'Tab',
    Query,
    types.model({
        queryId: types.maybe(types.string),
    }),
);

export interface ITab extends Instance<typeof Tab> {}
