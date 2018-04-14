import { ApiProvider } from '../../providers/api/api';
import { AuthProvider } from '../../providers/auth/auth';
import { BackgroundGeolocation, BackgroundGeolocationConfig, BackgroundGeolocationResponse } from '@ionic-native/background-geolocation';
import { Component } from '@angular/core';
import { DeviceMotion, DeviceMotionAccelerationData } from '@ionic-native/device-motion';
import { Gyroscope, GyroscopeOrientation, GyroscopeOptions } from '@ionic-native/gyroscope';
import { IonicPage, NavController, NavParams } from 'ionic-angular';
import { Platform } from 'ionic-angular';
import { HttpErrorResponse } from '@angular/common/http';
import { GlobalsProvider } from '../../providers/globals/globals';

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
  private backgroundGeolocationConfig: BackgroundGeolocationConfig;
  private socket : WebSocket;

  constructor(
    public api: ApiProvider,
    public authProvider: AuthProvider,
    public navCtrl: NavController,
    public navParams: NavParams,
    private backgroundGeolocation: BackgroundGeolocation,
    public platform: Platform,
    private gyroscope: Gyroscope,
    private deviceMotion: DeviceMotion,
    private globals: GlobalsProvider) {
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

    let game = this;

    this.socket = new WebSocket("ws://" + this.globals.API_URL + "/matches");
    this.socket.onopen = function() {
      this.send(JSON.stringify({message: 'hello server'}));
      console.log('sent');
    }

    this.socket.onmessage = function(event) {
      let m = JSON.parse(event.data);
      console.log("Received message", m.message);
      game.startPlaying();
    }

    this.socket.onerror = function(err) {
      console.log(err);
    }
  }

  ionViewDidLoad() {
    if (this.mobileDevice) {
      this.platform.ready().then(() => this.onCordovaReady())
    }
  }

  onCordovaReady() {
    this.startPlaying()
  }

  startPlaying(): void {
    if (!this.mobileDevice) {
      console.warn('Cannot start background geolocation because the app is not being run in a mobile device.')
      return
    }
    console.log('Math starting.')
    let json = {
      "latitude": 41.157944,
      "longitude": -8.629105
    }
    this.api.post('/locations/' + this.authProvider.userID, {
      "latitude": 41.157944,
      "longitude": -8.629105
    })
      .then(data => {
        console.log('aaaa')
        console.log(data);
      }).catch((error: HttpErrorResponse) => {
        console.log('bbbb')
        console.log(error);
      });
    this.backgroundGeolocation.configure(this.backgroundGeolocationConfig).subscribe((location: BackgroundGeolocationResponse) => {
      console.log('received location')
      console.log(location.coords)
      this.backgroundGeolocation.finish();
    }, (error: any) => {
      console.error(error)
    })
    this.backgroundGeolocation.start();
    this.startMatch()
  }

  stopPlaying(): void {
    if (this.mobileDevice) {
      this.backgroundGeolocation.stop()
    }
  }

  startMatch(): void {
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

  endMatch(): void {
    this.socket.close();
  }








}
