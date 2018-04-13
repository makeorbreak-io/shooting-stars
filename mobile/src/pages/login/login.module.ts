import { NgModule } from '@angular/core';
import { IonicPageModule } from 'ionic-angular';
import { LoginPage } from './login';
import { BackgroundGeolocation, BackgroundGeolocationConfig, BackgroundGeolocationResponse } from '@ionic-native/background-geolocation';

@NgModule({
  declarations: [
    LoginPage,
  ],
  imports: [
    IonicPageModule.forChild(LoginPage),
  ],
})
export class LoginPageModule {
  constructor(private backgroundGeolocation: BackgroundGeolocation) {
    const config: BackgroundGeolocationConfig = {
      desiredAccuracy: 10,
      stationaryRadius: 20,
      distanceFilter: 30,
      debug: false,
      startOnBoot: true,
      stopOnTerminate: false,
    };
  }
  onLogin() : void {
    this.backgroundGeolocation.start()
  }

  onLogout() : void {
    this.backgroundGeolocation.stop()
  }
}
