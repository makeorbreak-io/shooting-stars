import { Component } from '@angular/core';
import { IonicPage, NavController, NavParams } from 'ionic-angular';
import { BackgroundGeolocation, BackgroundGeolocationConfig, BackgroundGeolocationResponse } from '@ionic-native/background-geolocation';
import { Platform } from 'ionic-angular';

@IonicPage()
@Component({
  selector: 'page-login',
  templateUrl: 'login.html',
})
export class LoginPage {
  mobileDevice: boolean
  constructor(public navCtrl: NavController, public navParams: NavParams, private backgroundGeolocation: BackgroundGeolocation,
    public platform: Platform) {
      if (this.platform.is('core')) {
        this.mobileDevice = false
      } else {
        this.mobileDevice = true
      }
      if (this.mobileDevice) {
        const config: BackgroundGeolocationConfig = {
          desiredAccuracy: 10,
          stationaryRadius: 20,
          distanceFilter: 30,
          debug: false,
          startOnBoot: true,
          stopOnTerminate: false,
        };
        this.backgroundGeolocation.configure(config).subscribe((location: BackgroundGeolocationResponse) => {
          console.log('received location')
          console.log(location)
          this.backgroundGeolocation.finish();
        });
      }
  }

  ionViewDidLoad() {
    console.log('ionViewDidLoad LoginPage');
    this.startPlaying()
  }

  startPlaying() : void {
    if (!this.mobileDevice) {
      console.warn('Cannot start background geolocation because the app is not being run in a mobile device.')
      return
    }
    this.backgroundGeolocation.start()
  }

  stopPlaying() : void {
    if (this.mobileDevice) {
      this.backgroundGeolocation.stop()
    }
  }

}
