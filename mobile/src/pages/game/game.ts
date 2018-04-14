import { BackgroundGeolocation, BackgroundGeolocationConfig, BackgroundGeolocationResponse } from '@ionic-native/background-geolocation';
import { Component } from '@angular/core';
import { IonicPage, NavController, NavParams } from 'ionic-angular';
import { Platform } from 'ionic-angular';

/**
 * Generated class for the GamePage page.
 *
 * See https://ionicframework.com/docs/components/#navigation for more info on
 * Ionic pages and navigation.
 */

@IonicPage()
@Component({
  selector: 'page-game',
  templateUrl: 'game.html',
})
export class GamePage {
  private mobileDevice: boolean
  constructor(public navCtrl: NavController, public navParams: NavParams,
    private backgroundGeolocation: BackgroundGeolocation, public platform: Platform) {
      if (this.platform.is('cordova')) {
        this.mobileDevice = true
      } else {
        this.mobileDevice = false
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
          console.log(location.coords)
          this.backgroundGeolocation.finish();
        });
      }
  }

  ionViewDidLoad() {
    console.log('ionViewDidLoad GamePage');
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
