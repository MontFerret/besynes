import { Layout } from 'antd';
import React from 'react';
import { AppContent } from './content/index';
import { AppHeader } from './header/index';
import { AppSidebar } from './sidebar/index';

const { Header, Content, Sider } = Layout;

const styles = {
    layout: {
        height: '100%',
    },
    sider: {
        height: '100%',
        background: '#fff',
        minWidth: '250px',
        maxWidth: '400px',
    },
};

export class IndexScreen extends React.PureComponent {
    public render(): any {
        return (
            <Layout style={styles.layout}>
                <Header className="header">
                    <AppHeader />
                </Header>
                <Layout>
                    <Sider width={250} style={styles.sider}>
                        <AppSidebar />
                    </Sider>
                    <Layout style={{ padding: '0 24px 24px' }}>
                        <Content
                            style={{
                                background: '#fff',
                                padding: 24,
                                margin: 0,
                                minHeight: 280,
                            }}
                        >
                            <AppContent />
                        </Content>
                    </Layout>
                </Layout>
            </Layout>
        );
    }
}
