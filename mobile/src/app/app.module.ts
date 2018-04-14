import { NgModule, ErrorHandler } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { IonicApp, IonicModule, IonicErrorHandler } from 'ionic-angular';
import { MyApp } from './app.component';

import { GamePage } from '../pages/game/game';
import { HomePage } from '../pages/home/home';
import { TabsPage } from '../pages/tabs/tabs';
import { LoginPage } from '../pages/login/login';

import { BackgroundGeolocation } from '@ionic-native/background-geolocation';
import { DeviceMotion } from '@ionic-native/device-motion';
import { Geolocation } from '@ionic-native/geolocation';
import { Gyroscope } from '@ionic-native/gyroscope';
import { StatusBar } from '@ionic-native/status-bar';
import { SplashScreen } from '@ionic-native/splash-screen';
import { LoginProvider } from '../providers/login/login';

@NgModule({
  declarations: [
    MyApp,
    GamePage,
    HomePage,
    TabsPage,
    LoginPage
  ],
  imports: [
    BrowserModule,
    IonicModule.forRoot(MyApp)
  ],
  bootstrap: [IonicApp],
  entryComponents: [
    MyApp,
    GamePage,
    HomePage,
    TabsPage,
    LoginPage
  ],
  providers: [
    BackgroundGeolocation,
    DeviceMotion,
    Geolocation,
    Gyroscope,
    StatusBar,
    SplashScreen,
    {provide: ErrorHandler, useClass: IonicErrorHandler},
    LoginProvider
  ]
})
export class AppModule {}
