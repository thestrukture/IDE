const electron = require('electron')
// Module to control application life.
const app = electron.app
// Module to create native browser window.
const BrowserWindow = electron.BrowserWindow

const path = require('path')
const url = require('url')

let mainWindow
function createWindow () {
  // Create the browser window.
  mainWindow = new BrowserWindow({width: 1200, height: 600,allowRunningInsecureContent: true})

  // and load the index.html of the app.

mainWindow.loadURL('http://localhost:8884/index')

}
app.on('ready', createWindow)