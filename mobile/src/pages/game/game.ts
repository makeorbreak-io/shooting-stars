import { ApiProvider } from '../../providers/api/api';
import { AuthProvider } from '../../providers/auth/auth';
import { BackgroundMode } from '@ionic-native/background-mode';
import { BackgroundGeolocation, BackgroundGeolocationConfig, BackgroundGeolocationResponse } from '@ionic-native/background-geolocation';
import { Component } from '@angular/core';
import { DeviceMotion, DeviceMotionAccelerationData } from '@ionic-native/device-motion';
import { DeviceOrientation, DeviceOrientationCompassHeading } from '@ionic-native/device-orientation';
import { Gyroscope, GyroscopeOrientation, GyroscopeOptions } from '@ionic-native/gyroscope';
import { IonicPage, NavController, NavParams } from 'ionic-angular';
import { Platform } from 'ionic-angular';
import { HttpErrorResponse } from '@angular/common/http';
import { Vibration } from '@ionic-native/vibration';
import { GlobalsProvider } from '../../providers/globals/globals';
import { NativeAudio } from '@ionic-native/native-audio';
import { Flashlight } from '@ionic-native/flashlight';

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
  WAITING_MATCH_RESULT,
  VIEWING_MATCH_RESULT,
  RESTING
}
interface Rotation {
  x: number,
  y: number,
  z: number,
  timestamp: number
}

@IonicPage()
@Component({
  selector: 'page-game',
  templateUrl: 'game.html',
  providers: [ Flashlight, Vibration ]
})
export class GamePage {
  private mobileDevice: boolean
  private backgroundGeolocationConfig: BackgroundGeolocationConfig;
  private socket: WebSocket;
  private State = State;
  private state: State = State.RESTING;
  private searching: boolean = false;
  private totalRotation: Rotation = {
    x: 0,
    y: 0,
    z: 0,
    timestamp: 0
  }
  private gyroscopeSubscription
  private gyroscopeOptions: GyroscopeOptions = {
    frequency: 30
  };
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
    private vibration: Vibration,
    private deviceOrientation: DeviceOrientation,
    private flashlight: Flashlight) {
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
    this.backgroundMode.disableWebViewOptimizations()
    this.backgroundMode.moveToBackground();
    setTimeout(() => this.socket.send(JSON.stringify({ message: 'hello server' })), 3000);
  }

  stopPlaying(): void {
    this.state = State.RESTING
    this.backgroundGeolocation.stop()
  }

  spamWakeUp() {
    if (this.state != State.IN_MATCH) return;
    this.vibration.vibrate(5000);
    let vibrating: boolean = false;
    let interval = setInterval(() => {
      this.backgroundMode.unlock()
      this.vibration.vibrate(200);
    }, 200)
    setTimeout(() => clearInterval(interval), 5000)
  }

  startMatch() : void {
    this.state = State.IN_MATCH;
    this.backgroundMode.moveToForeground()
    this.platform.resume.asObservable().subscribe(() => {
      this.nativeAudio.play('westernWhistle', () => {});
    });
    this.spamWakeUp()
    this.totalRotation.x = 0
    this.totalRotation.y = 0
    this.totalRotation.z = 0
    this.gyroscopeSubscription = this.gyroscope.watch(this.gyroscopeOptions).subscribe((orientation: GyroscopeOrientation) => {
      if (this.state == State.IN_MATCH) {
        this.totalRotation.x += orientation.x
        this.totalRotation.y += orientation.y
        this.totalRotation.z += orientation.z
        console.log(this.totalRotation.x, this.totalRotation.y, this.totalRotation.z)
      }
    })
  }

  handleSuccessfullShot() {
    this.state = State.WAITING_MATCH_RESULT
    this.gyroscopeSubscription.unsubscribe()
    this.backgroundGeolocation.stop()
    this.vibration.vibrate(0);
    console.log('SHOT A SHERIFF!!!')
    this.api.post('/shoot/' + this.authProvider.userID, {
    })
    .then(data => {
      console.log(data);
      this.state = State.VIEWING_MATCH_RESULT
    }).catch((error: HttpErrorResponse) => {
      console.error(error);
      this.state = State.VIEWING_MATCH_RESULT
    });
  }

  handleMissedShot() {
    this.state = State.COOLDOWN;
    setTimeout(() => this.state = State.IN_MATCH, 500)
  }

  shoot() {
    if (!this.flashlight.isSwitchedOn()) {
      this.flashlight.switchOn()
      setTimeout(() => {
        this.flashlight.switchOff()
      }, 50)
    }
    this.nativeAudio.play('gunshot', () => {})
    this.deviceMotion.getCurrentAcceleration().then(
      (acceleration) => {
        if (Math.abs(acceleration.x) >= 8 && Math.abs(acceleration.y) <= 3 && Math.abs(acceleration.z) <= 5) {
          this.gyroscope.getCurrent(this.gyroscopeOptions).then((orientation: GyroscopeOrientation) => {
            if (Math.abs(this.totalRotation.x) > 30) {
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
