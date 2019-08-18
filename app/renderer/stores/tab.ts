import { Instance, flow, getEnv, types } from 'mobx-state-tree';
import { Tab } from '../../common/models/tab';
import { AppEnv } from '../env/env';

export const TabStore = types
    .compose(
        'TabStore',
        Tab,
        types.model({
            isDirty: types.optional(types.boolean, false),
            isLoading: types.optional(types.boolean, false),
            error: types.maybe(types.string),
            data: types.maybe(types.model({})),
        }),
    )
    .actions(self => {
        return {
            setQueryText(text: string): void {
                self.text = text;
            },

            execute: flow(function*(): any {
                const ctx = getEnv<AppEnv>(self);

                self.isLoading = true;
                self.error = undefined;
                self.data = undefined;

                try {
                    const data = yield ctx.api.executeQuery(self);
                    self.data = data;
                } catch (err) {
                    self.error = err.message;
                } finally {
                    self.isLoading = false;
                }
            }),
        };
    });

export interface ITabStore extends Instance<typeof TabStore> {}
