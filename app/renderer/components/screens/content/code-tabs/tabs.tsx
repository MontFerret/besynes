import { Icon, Tabs } from 'antd';
import { observer } from 'mobx-react';
import React from 'react';
import { ITabListStore } from '../../../../stores/tab-list';
import { TabContent } from './tab';

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
    tabs?: ITabListStore;
}

interface State {
    selected: number;
}

@observer
export class CodeTabs extends React.PureComponent<Props, State> {
    constructor(props: Props) {
        super(props);

        this.state = {
            selected: -1,
        };

        this.handleEdit = this.handleEdit.bind(this);
        this.handleSelect = this.handleSelect.bind(this);
    }

    private handleSelect(targetKey: string): void {
        if (targetKey === TAB_ADD_KEY) {
            this.handleEdit('', 'add');

            return;
        }

        if (this.state.selected.toString() === targetKey) {
            return;
        }

        this.setState({
            selected: parseFloat(targetKey),
        });
    }

    private handleEdit(targetKey: any, action: 'add' | 'remove'): void {
        if (action === 'add') {
            this.props.tabs!.add({
                id: '',
                name: 'Untitled query',
                description: '',
                text: '',
                params: {},
            });

            this.setState({
                selected: this.props.tabs!.items.length - 1,
            });

            return;
        }

        if (action === 'remove') {
            const idx = parseFloat(targetKey);

            if (!isNaN(idx)) {
                this.props.tabs!.remove(idx);
            }

            this.setState({
                selected:
                    this.props.tabs!.items.length > 0
                        ? this.props.tabs!.items.length - 1
                        : -1,
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
        const { selected } = this.state;
        const { tabs } = this.props;

        return (
            <Tabs
                activeKey={selected.toString()}
                type="editable-card"
                onEdit={this.handleEdit}
                onChange={this.handleSelect}
                tabBarGutter={5}
                hideAdd
                destroyInactiveTabPane
            >
                {tabs!.items.map((tab, idx) => (
                    <TabPane tab={tab.name} key={idx.toString()} closable>
                        <TabContent store={tab} />
                    </TabPane>
                ))}
                {this.renderSystemTabs()}
            </Tabs>
        );
    }
}
