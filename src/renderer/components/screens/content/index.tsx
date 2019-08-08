import { Col, Row } from 'antd';
import React from 'react';
import { CodeTabs } from '../../common/code-tabs/tabs';

export class AppContent extends React.PureComponent {
    public render(): any {
        return (
            <Row>
                <Col span={24}>
                    <CodeTabs />
                </Col>
            </Row>
        );
    }
}
