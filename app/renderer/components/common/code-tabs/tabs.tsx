import { Icon, Tabs } from 'antd';
import { Editor } from 'codemirror';
import nanoid from 'nanoid';
import React from 'react';
import { Query } from '../../../models/query';
import { QueryButtons } from './buttons';
import { QueryEditor } from './editor';
import { QueryResult } from './result';

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
    private editor?: Editor;

    constructor(props: Props) {
        super(props);

        this.state = {
            selected: '',
            queries: props.queries || [],
        };

        this.handleEdit = this.handleEdit.bind(this);
        this.handleSelect = this.handleSelect.bind(this);
        this.handleEditorDidMount = this.handleEditorDidMount.bind(this);
    }

    private handleSelect(targetKey: any): void {
        if (targetKey === TAB_ADD_KEY) {
            this.handleEdit('', 'add');

            return;
        }

        if (this.state.selected === targetKey) {
            return;
        }

        this.setState({
            selected: targetKey,
            queries: this.updateCurrentQuery(),
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
            const queries = [...this.updateCurrentQuery(), query];

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

    private handleEditorDidMount(editor: Editor): void {
        this.editor = editor;
    }

    private updateCurrentQuery(): Query[] {
        // Fined query object of a current tab
        const idx = this.state.queries.findIndex(
            q => q.id === this.state.selected,
        );
        let queries = this.state.queries;

        // Should be true
        if (idx > -1) {
            const current = this.state.queries[idx];

            // If editor is available
            if (current && this.editor) {
                // Copy all quriees
                queries = this.state.queries.slice();
                // Replace current query with most recent values
                queries[idx] = {
                    id: current.id,
                    name: current.name,
                    params: current.params,
                    text: this.editor.getValue(),
                };
            }
        }

        return queries;
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

    private renderTab(query: Query): any {
        return (
            <TabPane tab={query.name} key={query.id} closable={true}>
                <QueryEditor
                    text={query.text}
                    onEditorDidMount={this.handleEditorDidMount}
                />
                <QueryButtons />
                <QueryResult />
            </TabPane>
        );
    }

    public render(): any {
        const { selected, queries } = this.state;

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
                {queries.map(query => this.renderTab(query))}
                {this.renderSystemTabs()}
            </Tabs>
        );
    }
}
