import React from 'react';
import { Query } from '../../../models/query';

export interface Props {
    query?: Query;
}

export class TabContent extends React.PureComponent<Props> {
    public render(): any {
        const { query } = this.props;

        if (!query) {
            return <span>""</span>;
        }

        return <span>{query.text}</span>;
    }
}
