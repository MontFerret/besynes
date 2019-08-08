import ReactDOM from 'react-dom';
import App from './app';
// import './index.css';

const app = new App();

ReactDOM.render(app.createElement(), document.getElementById('app'));
