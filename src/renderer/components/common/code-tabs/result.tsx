import { Card } from 'antd';
import React from 'react';
import JSONTree from 'react-json-tree';

const styles = {
    card: {
        width: '100%',
    },
    cardBody: {
        padding: '0',
    },
};

const theme = {
    scheme: 'apathy',
    author: '',
    base00: '#031A16',
    base01: '#0B342D',
    base02: '#184E45',
    base03: '#2B685E',
    base04: '#5F9C92',
    base05: '#81B5AC',
    base06: '#A7CEC8',
    base07: '#D2E7E4',
    base08: '#3E9688',
    base09: '#3E7996',
    base0A: '#3E4C96',
    base0B: '#883E96',
    base0C: '#963E4C',
    base0D: '#96883E',
    base0E: '#4C963E',
    base0F: '#3E965B',
};

export interface Props {
    value?: any;
}

export class QueryResult extends React.PureComponent<Props> {
    public render(): any {
        if (this.props.value == null) {
            return '';
        }

        return (
            <Card style={styles.card} bodyStyle={styles.cardBody}>
                <JSONTree data={this.props.value} theme={theme} />
            </Card>
        );
    }
}
