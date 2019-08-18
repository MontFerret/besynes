import { Instance, types } from 'mobx-state-tree';
import { IQuery } from '../../common/models/query';
import { IQueryStore, QueryStore } from './query';

export const QueryListStore = types
    .model('QueryListStore', {
        isLoading: types.optional(types.boolean, false),
        items: types.array(QueryStore),
    })
    .actions(self => {
        return {
            addQuery(values: IQuery): void {
                self.items.push(values);
            },
            removeQuery(idx: number): void {
                self.items.splice(idx, 1);
            },
            findQueryById(id: string): IQueryStore | undefined {
                return self.items.find(i => i.id === id);
            },
        };
    });

export interface IQueryListStore extends Instance<typeof QueryListStore> {}
