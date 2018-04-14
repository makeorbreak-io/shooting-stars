import { ApiProvider } from '../../providers/api/api';
import { AuthProvider } from '../../providers/auth/auth';
import { BackgroundMode } from '@ionic-native/background-mode';
import { BackgroundGeolocation, BackgroundGeolocationConfig, BackgroundGeolocationResponse } from '@ionic-native/background-geolocation';
import { Component } from '@angular/core';
import { DeviceMotion, DeviceMotionAccelerationData } from '@ionic-native/device-motion';
import { Gyroscope, GyroscopeOrientation, GyroscopeOptions } from '@ionic-native/gyroscope';
import { IonicPage, NavController, NavParams } from 'ionic-angular';
import { Platform } from 'ionic-angular';
import { HttpErrorResponse } from '@angular/common/http';
import { Vibration } from '@ionic-native/vibration';
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
  providers: [ Vibration ]
})
export class GamePage {
  private mobileDevice: boolean
  private backgroundGeolocationConfig: BackgroundGeolocationConfig;
  private socket: WebSocket;

  constructor(
    public api: ApiProvider,
    public authProvider: AuthProvider,
    public navCtrl: NavController,
    public navParams: NavParams,
    private backgroundMode: BackgroundMode,
    private backgroundGeolocation: BackgroundGeolocation,
    public platform: Platform,
    private gyroscope: Gyroscope,
    private deviceMotion: DeviceMotion,
    private globals: GlobalsProvider,
    private vibration: Vibration) {
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
    this.socket.onopen = function () {
      this.send(JSON.stringify({ message: 'hello server' }));
      console.log('sent');
    }

    this.socket.onmessage = function (event) {
      let m = JSON.parse(event.data);
      console.log("Received message", m.message);
      game.startPlaying();
    }

    this.socket.onerror = function (err) {
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
    console.log('Match starting.')
    this.backgroundGeolocation.configure(this.backgroundGeolocationConfig).subscribe((location: BackgroundGeolocationResponse) => {
      console.log('received location')
      console.log(location.latitude, location.longitude, location.speed)
      if (!location.speed) {
        location.speed = 0
      }
      this.updateLocation(location.latitude, location.longitude, location.speed)
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

<<<<<<< HEAD
  startMatch() : void {
    this.backgroundMode.enable();
    this.vibration.vibrate(3000);
    
=======
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

>>>>>>> e1f1918d3461bdb4757589446ec6ec9e74ef2faf
    // Get the device current acceleration
    this.deviceMotion.getCurrentAcceleration().then(
      this.handleAccelerometer,
      (error: any) => console.log(error)
    );
  }

  handleAccelerometer(acceleration: DeviceMotionAccelerationData) {
    console.log('acceleration', acceleration);
    if (Math.abs(acceleration.x) >= 9 && Math.abs(acceleration.y) <= 2 && Math.abs(acceleration.z) <= 2) {
      let options: GyroscopeOptions = {
        frequency: 30
      };
      this.gyroscope.getCurrent(options)
      .then((orientation: GyroscopeOrientation) => {
        if (Math.hypot(orientation.x, orientation.y, orientation.z) <= 1) {
          this.shoot()
        } else {
          this.deviceMotion.getCurrentAcceleration().then(
            this.handleAccelerometer,
            (error: any) => console.log(error)
          );
        }
      })
      .catch()
    } else {
      this.deviceMotion.getCurrentAcceleration().then(
        this.handleAccelerometer,
        (error: any) => console.log(error)
      );
    }
  }

  shoot() {
    this.state = State.WAITING
    console.log('SHOT A SHERIFF!!!')
  }

  endMatch(): void { }

  private updateLocation(latitude: number, longitude: number, speed: number): void {
    this.api.post('/locations/' + this.authProvider.userID, {
      "latitude": latitude,
      "longitude": longitude,
      "speed": speed
    })
      .then(data => {
        console.log(data);
      }).catch((error: HttpErrorResponse) => {
        console.error(error);
      });
  }
}
