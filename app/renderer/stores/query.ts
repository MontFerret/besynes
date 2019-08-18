import { Instance, types } from 'mobx-state-tree';
import { Query } from '../../common/models/query';

export const QueryStore = types.compose(
    'QueryStore',
    Query,
    types.model({}),
);

export interface IQueryStore extends Instance<typeof QueryStore> {}
