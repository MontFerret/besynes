import { Tabs } from 'antd';
import React from 'react';
import { Query } from '../../../models/query';
import { TabContent } from './content';

const { TabPane } = Tabs;
const DEFAULT_QUERIES: Query[] = [];

export interface Props {
    queries?: Query[];
}

interface State {
    selected: number;
}

export class CodeTabs extends React.PureComponent<Props, State> {
    constructor(props: Props) {
        super(props);

        this.state = {
            selected: -1,
        };
    }

    public render(): any {
        const { queries = DEFAULT_QUERIES } = this.props;
        const { selected } = this.state;

        const value = queries[selected];
        const key = value ? value.id : '';

        return (
            <Tabs
                // onChange={this.onChange}
                activeKey={key}
                type="editable-card"
                // sonEdit={this.onEdit}
            >
                {queries.map(query => (
                    <TabPane tab={query.name} key={query.id} closable={true}>
                        <TabContent query={query} />
                    </TabPane>
                ))}
            </Tabs>
        );
    }
}
