import { Col, Row } from 'antd';
import { inject, observer } from 'mobx-react';
import React from 'react';
import { ITabListStore } from '../../../stores/tab-list';
import { CodeTabs } from './code-tabs/tabs';

export interface Props {
    tabs?: ITabListStore;
}

@inject('tabs')
@observer
export class AppContent extends React.PureComponent<Props> {
    public render(): any {
        return (
            <Row>
                <Col span={24}>
                    <CodeTabs tabs={this.props.tabs} />
                </Col>
            </Row>
        );
    }
}
