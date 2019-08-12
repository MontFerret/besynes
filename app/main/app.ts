import { App, BrowserWindow } from 'electron';
import isDev from 'electron-is-dev';
import path from 'path';

export default class Application {
    private static __mainWindow?: BrowserWindow;
    private static __app: App;

    public static run(app: App): void {
        app.on('ready', Application.__onReady);
        // Quit when all windows are closed.
        app.on('window-all-closed', Application.__onWindowAllClosed);
        app.on('activate', Application.__onActivate);
    }

    private static __onReady(): void {
        if (Application.__mainWindow != null) {
            return;
        }

        // Create the browser window.
        Application.__mainWindow = new BrowserWindow({
            width: 1024,
            height: 768,
            webPreferences: {
                nodeIntegration: true,
            },
        });

        // and load the index.html of the app.
        Application.__mainWindow.loadFile(
            path.join(__dirname, '../renderer/index.html'),
        );
        Application.__mainWindow.on('close', Application.__onWindowClose);

        if (isDev) {
            Application.__mainWindow.webContents.openDevTools();
        }
    }

    private static __onWindowAllClosed(): void {
        // On macOS it is common for applications and their menu bar
        // to stay active until the user quits explicitly with Cmd + Q
        if (process.platform !== 'darwin') {
            Application.__app.quit();
        }
    }

    private static __onWindowClose(): void {
        // Dereference the window object.
        Application.__mainWindow = undefined;
    }

    private static __onActivate(): void {
        // On macOS it's common to re-create a window in the app when the
        // dock icon is clicked and there are no other windows open.
        if (Application.__mainWindow == null) {
            Application.__onReady();
        }
    }
}
