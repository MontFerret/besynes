import { Col, Row } from 'antd';
import { observer } from 'mobx-react';
import React from 'react';
import { IAppStore } from '../../../stores/app';
import { StoreConsumer } from '../../common/context';
import { CodeTabs } from './code-tabs/tabs';

@observer
export class AppContent extends React.PureComponent {
    public render(): any {
        return (
            <Row>
                <Col span={24}>
                    <StoreConsumer>
                        {(value: IAppStore) => {
                            return <CodeTabs tabs={value.tabs} />;
                        }}
                    </StoreConsumer>
                </Col>
            </Row>
        );
    }
}
