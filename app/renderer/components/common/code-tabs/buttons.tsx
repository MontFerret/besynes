import { Button, Col, Row } from 'antd';
import React from 'react';

const styles: { [key: string]: React.CSSProperties } = {
    container: {
        marginTop: '10px',
        marginBottom: '10px',
    },
    button: {
        float: 'right',
        width: '100px',
    },
};

export interface Props {
    onRun?: () => void;
}

export class QueryButtons extends React.PureComponent {
    public render(): any {
        return (
            <Row type="flex" justify="end" style={styles.container}>
                <Col span={3}>
                    <Button type="primary" style={styles.button}>
                        Run
                    </Button>
                </Col>
            </Row>
        );
    }
}
