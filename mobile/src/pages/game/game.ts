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
import { NativeAudio } from '@ionic-native/native-audio';

/**
 * Generated class for the GamePage page.
 *
 * See https://ionicframework.com/docs/components/#navigation for more info on
 * Ionic pages and navigation.
 */

enum State {
  WAITING,
  IN_MATCH,
  COOLDOWN,
  RESTING
}

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
  private State = State;
  private state: State = State.RESTING;
  private searching: boolean = false;
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
    private nativeAudio: NativeAudio,
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

    //this.socket = new WebSocket("ws://" + this.globals.API_URL + "/matches");
    this.socket = new WebSocket("ws://echo.websocket.org");
    this.socket.onopen = function () {
      setTimeout(() => this.send(JSON.stringify({ message: 'hello server' })), 3000);
    }

    this.socket.onmessage = (event) => {
      let m = JSON.parse(event.data);
      console.log("Received message", m.message);
      if (this.state == State.WAITING) {
        game.startMatch();
      }
    }

    this.socket.onerror = function (err) {
      console.log(err);
    }
  }

  ionViewDidLoad() {
    this.state = State.RESTING
    if (this.mobileDevice) {
      this.platform.ready().then(() => this.startPlaying())
    }
  }

  startPlaying(): void {
    this.state = State.WAITING
    this.searching = true;
    this.backgroundMode.enable();
    if (!this.mobileDevice) {
      console.warn('Cannot start background geolocation because the app is not being run in a mobile device.')
      return
    }
    console.log('Waiting for a match.')
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
  }

  stopPlaying(): void {
    this.state = State.RESTING
    if (this.mobileDevice) {
      this.backgroundGeolocation.stop()
    }
  }

  startMatch() : void {
    this.state = State.IN_MATCH;
    this.backgroundMode.wakeUp()
    this.backgroundMode.moveToForeground()
    this.backgroundMode.unlock()
    this.platform.resume.asObservable().subscribe(() => {
      this.nativeAudio.play('westernWhistle', () => {});
      this.vibration.vibrate(3000);
    });
  }

  handleSuccessfullShot() {
    this.state = State.RESTING;
    console.log('SHOT A SHERIFF!!!')
    this.api.post('/shoot/' + this.authProvider.userID, {
    })
    .then(data => {
      console.log(data);
    }).catch((error: HttpErrorResponse) => {
      console.error(error);
    });
  }

  handleMissedShot() {
    this.state = State.COOLDOWN;
    setTimeout(() => this.state = State.IN_MATCH, 500)
  }

  shoot() {
    this.nativeAudio.play('gunshot', () => {})
    this.deviceMotion.getCurrentAcceleration().then(
      (acceleration) => {
        console.log(acceleration.x, acceleration.y, acceleration.z, acceleration.timestamp)
        if (Math.abs(acceleration.x) >= 9 && Math.abs(acceleration.y) <= 3 && Math.abs(acceleration.z) <= 3) {
          let options: GyroscopeOptions = {
            frequency: 30
          };
          this.gyroscope.getCurrent(options).then((orientation: GyroscopeOrientation) => {
            console.log(orientation.x, orientation.y, orientation.z, orientation.timestamp)
            if (Math.hypot(orientation.x, orientation.y, orientation.z) <= 1) {
              this.handleSuccessfullShot()
            } else {
              this.handleMissedShot()
            }
          })
          .catch()
        } else {
          this.handleMissedShot()
        }
      },
      (error: any) => console.log(error)
    );
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
