import { app } from 'electron';
import { Bootloader } from './boot';

app.commandLine.appendSwitch('remote-debugging-port', '8315');

Bootloader.load(app);
