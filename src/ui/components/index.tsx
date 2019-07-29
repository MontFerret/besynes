import { History } from 'history';
import React from 'react';
// import { Route, Router } from 'react-router';
import { IndexScreen } from './screens/index';

export interface Props {
    history: History;
}

export class AppComponent extends React.PureComponent<Props> {
    public render(): any {
        // const { history } = this.props;

        // return (
        //     <Router history={history}>
        //         <Route path="/" component={IndexScreen as any} />
        //     </Router>
        // );

        return <IndexScreen />;
    }
}

export function create(history: History): React.ReactElement<any> {
    return <AppComponent history={history} />;
}
