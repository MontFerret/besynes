import { Card } from 'antd';
import React from 'react';
import { UnControlled as CodeMirror } from 'react-codemirror2';
import { Query } from '../../../models/query';

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
};

export interface Props {
    query: Query;
}

export class TabContent extends React.PureComponent<Props> {
    constructor(props: Props) {
        super(props);

        this.handleChange = this.handleChange.bind(this);
    }

    private handleChange(editor: any, data: any, value: any): void {
        console.log(editor, data, value);
    }

    public render(): any {
        return (
            <Card style={styles.card} bodyStyle={styles.cardBody}>
                <CodeMirror
                    value={this.props.query.text}
                    options={OPTIONS}
                    onChange={this.handleChange}
                />
            </Card>
        );
    }
}
