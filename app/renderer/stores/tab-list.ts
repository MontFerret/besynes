import { Instance, types } from 'mobx-state-tree';
import { IQuery } from '../../common/models/query';
import { ITabStore, TabStore } from './tab';

export const TabListStore = types
    .model('TabListStore', {
        isLoading: types.optional(types.boolean, false),
        error: types.maybe(types.string),
        items: types.array(TabStore),
    })
    .actions(self => {
        return {
            add(values: IQuery): void {
                self.items.push(values);
            },
            remove(idx: number): void {
                self.items.splice(idx, 1);
            },
            findById(id: string): ITabStore | undefined {
                return self.items.find(i => i.id === id);
            },
        };
    });

export interface ITabListStore extends Instance<typeof TabListStore> {}
