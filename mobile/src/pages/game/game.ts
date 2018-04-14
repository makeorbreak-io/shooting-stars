import { AuthProvider } from '../../providers/auth/auth';
import { BackgroundGeolocation, BackgroundGeolocationConfig, BackgroundGeolocationResponse } from '@ionic-native/background-geolocation';
import { Component } from '@angular/core';
import { DeviceMotion, DeviceMotionAccelerationData } from '@ionic-native/device-motion';
import { HTTP } from '@ionic-native/http'
import { Gyroscope, GyroscopeOrientation, GyroscopeOptions } from '@ionic-native/gyroscope';
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
  private backgroundGeolocationConfig: BackgroundGeolocationConfig
  constructor(
    public authProvider: AuthProvider,
    public http: HTTP,
    public navCtrl: NavController,
    public navParams: NavParams,
    private backgroundGeolocation: BackgroundGeolocation,
    public platform: Platform,
    private gyroscope: Gyroscope,
    private deviceMotion: DeviceMotion) {
      if (this.platform.is('cordova')) {
        this.mobileDevice = true
      } else {
        this.mobileDevice = false
      }
      if (this.mobileDevice) {
        this.backgroundGeolocationConfig = {
          desiredAccuracy: 0,
          stationaryRadius: 0,
          distanceFilter: 0,
          debug: true,
          notificationTitle: 'Shooting Stars',
          notificationText: 'Game is running, be prepared for combat!',
          startOnBoot: true,
          stopOnTerminate: false,
        };
      }
  }

  ionViewDidLoad() {
    console.log('ionViewDidLoad GamePage');
    if (this.mobileDevice) {
      this.platform.ready().then(() => this.onCordovaReady())
    }
  }

  onCordovaReady() {
    console.log('cordova ready')
    this.startPlaying()
  }

  startPlaying() : void {
    if (!this.mobileDevice) {
      console.warn('Cannot start background geolocation because the app is not being run in a mobile device.')
      return
    }
    this.http.post('https://shooting-stars.herokuapp.com/locations/' + this.authProvider.userID, {
      'latitude': 1,
      'longitude': 1
      }, {
        'Authorization': this.authProvider.token
      })
      .then(data => {
        console.log(data.data);
      }).catch(error => {
        console.log(error.status);
      });
    this.backgroundGeolocation.configure(this.backgroundGeolocationConfig).subscribe((location: BackgroundGeolocationResponse) => {
      console.log('received location')
      console.log(location.coords)
      console.log('locations/' + this.authProvider.userID)
      this.backgroundGeolocation.finish();
    });
    this.backgroundGeolocation.start()
    this.startMatch()
  }

  stopPlaying() : void {
    if (this.mobileDevice) {
      this.backgroundGeolocation.stop()
    }
  }

  startMatch() : void {
    let options: GyroscopeOptions = {
      frequency: 1000
    };
   
    this.gyroscope.getCurrent(options)
      .then((orientation: GyroscopeOrientation) => {
        console.log(orientation.x, orientation.y, orientation.z, orientation.timestamp);
      })
      .catch()
   
    let gyroscopeSubscription = this.gyroscope.watch(options)
      .subscribe((orientation: GyroscopeOrientation) => {
        console.log(orientation.x, orientation.y, orientation.z, orientation.timestamp);
      });
    
    // Get the device current acceleration
    this.deviceMotion.getCurrentAcceleration().then(
      (acceleration: DeviceMotionAccelerationData) => console.log(acceleration),
      (error: any) => console.log(error)
    );

    // Watch device acceleration
    var accelerometerSubscription = this.deviceMotion.watchAcceleration().subscribe((acceleration: DeviceMotionAccelerationData) => {
      console.log('acceleration', acceleration);
    });

    gyroscopeSubscription.unsubscribe();
    accelerometerSubscription.unsubscribe();
  }

  endMatch() : void {}
}
