import { Instance, types } from 'mobx-state-tree';

export const Query = types.model('Query', {
    id: types.identifier,
    name: types.string,
    description: types.optional(types.string, ''),
    text: types.optional(types.string, ''),
    params: types.maybeNull(types.model()),
});

export interface IQuery extends Instance<typeof Query> {}
