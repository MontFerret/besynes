import { App as Electron, BrowserWindow, ipcMain } from 'electron';
import { Socket, Transport } from 'electron-ipc-socket';
import isDev from 'electron-is-dev';
import path from 'path';
import { Application } from './app';
import { DynamicOutput } from './ipc';

export class Bootloader {
    private static __mainWindow?: BrowserWindow;
    private static __electron: Electron;
    private static __output?: DynamicOutput;
    private static __app?: Application;

    public static load(electron: Electron): void {
        electron.on('ready', Bootloader.__onReady);

        // Quit when all windows are closed.
        electron.on('window-all-closed', Bootloader.__onWindowAllClosed);
        electron.on('activate', Bootloader.__onActivate);
        electron.on('will-quit', Bootloader.__onQuit);
    }

    private static __onReady(): void {
        if (Bootloader.__mainWindow != null) {
            return;
        }

        // Create the browser window.
        Bootloader.__mainWindow = new BrowserWindow({
            width: 1024,
            height: 768,
            webPreferences: {
                nodeIntegration: true,
            },
        });

        // and load the index.html of the app.
        Bootloader.__mainWindow.loadFile(
            path.join(__dirname, '../renderer/index.html'),
        );
        Bootloader.__mainWindow.on('close', Bootloader.__onWindowClose);

        if (isDev) {
            Bootloader.__mainWindow.webContents.openDevTools();
        }

        if (Bootloader.__output == null) {
            Bootloader.__output = new DynamicOutput(
                Bootloader.__mainWindow.webContents,
            );
        } else {
            Bootloader.__output.setRenderer(
                Bootloader.__mainWindow.webContents,
            );
        }

        if (Bootloader.__app == null) {
            Bootloader.__app = new Application({
                socket: new Socket(new Transport(ipcMain, Bootloader.__output)),
            });
        }
    }

    private static __onWindowAllClosed(): void {
        // On macOS it is common for applications and their menu bar
        // to stay active until the user quits explicitly with Cmd + Q
        if (process.platform !== 'darwin') {
            if (Bootloader.__app != null) {
                Bootloader.__app.dispose();
            }

            Bootloader.__electron.quit();
        }
    }

    private static __onWindowClose(): void {
        // Dereference the window object.
        Bootloader.__mainWindow = undefined;
    }

    private static __onActivate(): void {
        // On macOS it's common to re-create a window in the app when the
        // dock icon is clicked and there are no other windows open.
        if (Bootloader.__mainWindow == null) {
            Bootloader.__onReady();
        }
    }

    private static __onQuit(): void {
        if (Bootloader.__app != null) {
            Bootloader.__app.dispose();
        }
    }
}
