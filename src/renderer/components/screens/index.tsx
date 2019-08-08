import { Layout } from 'antd';
import React from 'react';
import SplitPane from 'react-split-pane';
import { AppContent } from './content/index';
import { AppHeader } from './header/index';
import { AppSidebar } from './sidebar/index';

const { Header, Content, Sider } = Layout;

export class IndexScreen extends React.PureComponent {
    public render(): any {
        return (
            <Layout>
                <Header className="header">
                    <AppHeader />
                </Header>

                <SplitPane split="vertical" minSize={50} defaultSize={200}>
                    <Layout>
                        <Sider width={200} style={{ background: '#fff' }}>
                            <AppSidebar />
                        </Sider>
                    </Layout>
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
                </SplitPane>
            </Layout>
        );
    }
}
