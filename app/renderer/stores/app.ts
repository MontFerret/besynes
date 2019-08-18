import { Instance, types } from 'mobx-state-tree';
import { QueryListStore } from './query-list';
import { TabListStore } from './tab-list';

export const AppStore = types.model({
    queries: types.optional(QueryListStore, {
        items: [],
    }),
    tabs: types.optional(TabListStore, {
        items: [],
    }),
});

export interface IAppStore extends Instance<typeof AppStore> {}
