import { Card } from 'antd';
import * as codemirror from 'codemirror';
import React from 'react';
import { UnControlled as CodeMirror } from 'react-codemirror2';
require('codemirror/mode/javascript/javascript');

const styles = {
    card: {
        width: '100%',
    },
    cardBody: {
        padding: '0',
    },
};
const OPTIONS = {
    lineNumbers: true,
    mode: 'javascript',
    // theme: 'material',
};

export interface Props {
    text: string;
    onEditorDidMount?: (editor: codemirror.Editor) => void;
}

export class QueryEditor extends React.PureComponent<Props> {
    constructor(props: Props) {
        super(props);

        this.handleEditorDidMount = this.handleEditorDidMount.bind(this);
    }

    public handleEditorDidMount(editor: codemirror.Editor): void {
        editor.setValue(this.props.text);

        if (this.props.onEditorDidMount) {
            this.props.onEditorDidMount(editor);
        }
    }

    public render(): any {
        return (
            <Card style={styles.card} bodyStyle={styles.cardBody}>
                <CodeMirror
                    options={OPTIONS}
                    detach={false}
                    editorDidMount={this.handleEditorDidMount}
                />
            </Card>
        );
    }
}
