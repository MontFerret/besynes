import { Icon, Tabs } from 'antd';
import nanoid from 'nanoid';
import React from 'react';
import { Query } from '../../../models/query';
import { TabContent } from './content';

const { TabPane } = Tabs;
const TAB_ADD_KEY = '$add';
const TAB_MORE_KEY = '$more';
const styles = {
    tabs: {
        system: {
            width: '30px',
            padding: '0px 7px',
        },
    },
};

export interface Props {
    queries?: Query[];
}

interface State {
    selected: string;
    queries: Query[];
}

export class CodeTabs extends React.PureComponent<Props, State> {
    constructor(props: Props) {
        super(props);

        this.state = {
            selected: '',
            queries: props.queries || [],
        };

        this.handleEdit = this.handleEdit.bind(this);
        this.handleSelect = this.handleSelect.bind(this);
    }

    private handleSelect(targetKey: any): void {
        if (targetKey === TAB_ADD_KEY) {
            this.handleEdit('', 'add');

            return;
        }

        this.setState({
            selected: targetKey,
        });
    }

    private handleEdit(targetKey: any, action: 'add' | 'remove'): void {
        if (action === 'add') {
            const query: Query = {
                id: nanoid(),
                name: 'Untitled query',
                text: '',
                params: {},
            };
            const queries = [...this.state.queries, query];

            this.setState({
                queries,
                selected: query.id,
            });

            return;
        }

        if (action === 'remove') {
            const queries = this.state.queries.filter(i => i.id !== targetKey);

            this.setState({
                queries,
                selected:
                    queries.length > 0 ? queries[queries.length - 1].id : '',
            });
        }
    }

    private renderSystemTabs(): any {
        return [
            <TabPane
                tab={<Icon type="plus" />}
                key={TAB_ADD_KEY}
                style={styles.tabs.system}
                closable={false}
            ></TabPane>,
            <TabPane
                tab={<Icon type="more" />}
                key={TAB_MORE_KEY}
                style={styles.tabs.system}
                closable={false}
            ></TabPane>,
        ];
    }

    public render(): any {
        const { selected, queries } = this.state;

        const tabs = queries.map(query => (
            <TabPane tab={query.name} key={query.id} closable={true}>
                <TabContent query={query} />
            </TabPane>
        ));

        return (
            <Tabs
                activeKey={selected}
                type="editable-card"
                onEdit={this.handleEdit}
                onChange={this.handleSelect}
                tabBarGutter={5}
                hideAdd
                destroyInactiveTabPane
            >
                {tabs}
                {this.renderSystemTabs()}
            </Tabs>
        );
    }
}
