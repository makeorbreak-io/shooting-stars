import { NgModule, ErrorHandler } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { IonicApp, IonicModule, IonicErrorHandler } from 'ionic-angular';
import { MyApp } from './app.component';

import { GamePage } from '../pages/game/game';
import { TabsPage } from '../pages/tabs/tabs';
import { LoginPage } from '../pages/login/login';
import { RegisterPage } from '../pages/register/register';
import { ProfilePage } from '../pages/profile/profile';
import { LeaderboardPage } from '../pages/leaderboard/leaderboard';

import { BackgroundGeolocation } from '@ionic-native/background-geolocation';
import { BackgroundMode } from '@ionic-native/background-mode';
import { DeviceMotion } from '@ionic-native/device-motion';
import { Geolocation } from '@ionic-native/geolocation';
import { Gyroscope } from '@ionic-native/gyroscope';
import { HTTP } from '@ionic-native/http'
import { NativeAudio } from '@ionic-native/native-audio';
import { StatusBar } from '@ionic-native/status-bar';
import { SplashScreen } from '@ionic-native/splash-screen';
import { ApiProvider } from '../providers/api/api';
import { HttpClientModule} from '@angular/common/http';
import { GlobalsProvider } from '../providers/globals/globals';
import { AuthProvider } from '../providers/auth/auth';

@NgModule({
  declarations: [
    MyApp,
    GamePage,
    TabsPage,
    LoginPage,
    RegisterPage,
    LeaderboardPage,
    ProfilePage
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    IonicModule.forRoot(MyApp)
  ],
  bootstrap: [IonicApp],
  entryComponents: [
    MyApp,
    GamePage,
    TabsPage,
    LoginPage,
    RegisterPage,
    LeaderboardPage,
    ProfilePage
  ],
  providers: [
    BackgroundGeolocation,
    BackgroundMode,
    DeviceMotion,
    Geolocation,
    Gyroscope,
    HTTP,
    NativeAudio,
    StatusBar,
    SplashScreen,
    {provide: ErrorHandler, useClass: IonicErrorHandler},
    ApiProvider,
    HttpClientModule,
    GlobalsProvider,
    AuthProvider
  ]
})
export class AppModule {}
