import { Editor } from 'codemirror';
import { observer } from 'mobx-react';
import React from 'react';
import { ITabStore } from '../../../../stores/tab';
import { QueryButtons } from './buttons';
import { QueryEditor } from './editor';
import { QueryResult } from './result';

export interface Props {
    store: ITabStore;
}

@observer
export class TabContent extends React.PureComponent<Props> {
    private __editor?: Editor;

    constructor(props: Props) {
        super(props);

        this.handleEditorDidMount = this.handleEditorDidMount.bind(this);
    }

    private handleEditorDidMount(editor: Editor): void {
        this.__editor = editor;
    }

    public componentWillUnmount(): void {
        if (this.__editor != null) {
            this.props.store.setQueryText(this.__editor.getValue());
        }
    }

    public render(): any {
        const { store } = this.props;

        return (
            <>
                <QueryEditor
                    text={store.text}
                    onEditorDidMount={this.handleEditorDidMount}
                />
                <QueryButtons />
                <QueryResult />
            </>
        );
    }
}
