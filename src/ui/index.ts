// import { createMemoryHistory as createHistory } from 'history';
import ReactDOM from 'react-dom';
import App from './app';
// import './index.css';

const app = new App(null as any);

ReactDOM.render(app.createElement(), document.getElementById('root'));
