import { ApiProvider } from '../../providers/api/api';
import { AuthProvider } from '../../providers/auth/auth';
import { BackgroundMode } from '@ionic-native/background-mode';
import { BackgroundGeolocation, BackgroundGeolocationConfig, BackgroundGeolocationResponse } from '@ionic-native/background-geolocation';
import { Component } from '@angular/core';
import { DeviceMotion } from '@ionic-native/device-motion';
import { DeviceOrientation } from '@ionic-native/device-orientation';
import { Geolocation } from '@ionic-native/geolocation';
import { Gyroscope, GyroscopeOrientation, GyroscopeOptions } from '@ionic-native/gyroscope';
import { IonicPage, NavController, NavParams } from 'ionic-angular';
import { Platform } from 'ionic-angular';
import { HttpErrorResponse } from '@angular/common/http';
import { Vibration } from '@ionic-native/vibration';
import { NativeAudio } from '@ionic-native/native-audio';
import { Flashlight } from '@ionic-native/flashlight';
import { GlobalsProvider } from '../../providers/globals/globals';
import { AlertController } from 'ionic-angular';

enum State {
  WAITING,
  IN_MATCH,
  COOLDOWN,
  WAITING_MATCH_RESULT,
  VIEWING_MATCH_RESULT,
  RESTING,
  CONNECTING
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
  providers: [Flashlight, Vibration]
})
export class GamePage {
  private mobileDevice: boolean
  private socket: WebSocket;
  private State = State; // this is required for the angular template to have access to the State enum
  uselessVarToShutUpTSLint = this.State.IN_MATCH;
  private state: State = State.RESTING;
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
    public auth: AuthProvider,
    public navCtrl: NavController,
    public navParams: NavParams,
    private backgroundMode: BackgroundMode,
    private backgroundGeolocation: BackgroundGeolocation,
    public platform: Platform,
    private geolocation: Geolocation,
    private gyroscope: Gyroscope,
    private deviceMotion: DeviceMotion,
    private nativeAudio: NativeAudio,
    private vibration: Vibration,
    private globals: GlobalsProvider,
    private deviceOrientation: DeviceOrientation,
    private flashlight: Flashlight,
    private alertCtrl: AlertController) {
    if (this.platform.is('cordova')) {
      this.mobileDevice = true
    } else {
      this.mobileDevice = false
    }
  }

  ionViewDidLoad() {
    this.state = State.RESTING;
    /* if (this.mobileDevice) {
      this.platform.ready().then(() => this.startPlaying())
    } */
  }

  startPlaying(): void {
    this.socket = new WebSocket(this.globals.SOCKET_URL + "/websocket/" + this.auth.token);
    //this.socket = new WebSocket("ws://echo.websocket.org");

    this.socket.onopen = () => {
      this.socket.send(this.auth.token);
      this.state = State.CONNECTING;
      console.log('sent token to socket');
    }

    this.socket.onmessage = (event) => {

      if (event.data === "PONG") {
        setTimeout(() => {
          this.socket.send("PING");
        }, 1000);
      }

      if (this.state === State.CONNECTING && event.data === "OK") {
        this.state = State.WAITING;
        console.log("set state to WAITING");
        this.socket.send("PING");
        return;
      }
      if (this.state == State.WAITING && event.data === "DUEL") {
        console.log("received DUEL");
        this.startMatch();
        return;
      }

    }

    this.socket.onerror = () => {
      this.alertCtrl.create({
        title: "Connection failed.",
        subTitle: "Please, try again.",
        buttons: ['Ok']
      }).present();
    }

    this.socket.onclose = (message) => {
      this.state = State.RESTING;
      console.log("Connection Closed");
    }

    this.backgroundMode.enable();

    if (!this.mobileDevice) {
      console.warn('Cannot start background geolocation because the app is not being run in a mobile device.')
      return
    }
    console.log('Waiting for a match.')
    this.fetchLocation();
    this.backgroundGeolocation.start();
    this.backgroundMode.disableWebViewOptimizations()
    this.backgroundMode.moveToBackground();
  }

  stopPlaying(): void {
    this.socket.send("CLOSE");
    console.log("Sent message CLOSE");
    this.state = State.RESTING;
    this.backgroundGeolocation.stop();
    this.socket.close(1000);
  }

  spamWakeUp() {
    if (this.state != State.IN_MATCH) return;
    this.vibration.vibrate(5000);
    let interval = setInterval(() => {
      this.backgroundMode.unlock();
      this.vibration.vibrate(200);
    }, 200)
    setTimeout(() => clearInterval(interval), 5000)
  }

  startMatch(): void {
    this.state = State.IN_MATCH;
    this.backgroundMode.moveToForeground()
    this.platform.resume.asObservable().subscribe(() => {
      this.nativeAudio.play('westernWhistle', () => { });
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
        // console.log(this.totalRotation.x, this.totalRotation.y, this.totalRotation.z)
      }
    })
    setTimeout(() => {
      if (this.state === State.IN_MATCH) {
        this.sendShoot();
      }
    }, 5000);
  }

  handleSuccessfullShot() {
    this.state = State.WAITING_MATCH_RESULT;
    this.gyroscopeSubscription.unsubscribe()
    this.backgroundGeolocation.stop()
    this.vibration.vibrate(0);
    console.log('SHOT A SHERIFF!!!')
    this.sendShoot();
  }

  sendShoot() {
    this.api.post('/matches/' + this.auth.userID + '/shoot', {
    })
      .then(data => {
        console.log(JSON.stringify(data));
        this.state = State.VIEWING_MATCH_RESULT
      }).catch((error: HttpErrorResponse) => {
        this.state = State.VIEWING_MATCH_RESULT
      });
  }

  handleMissedShot() {
    this.state = State.COOLDOWN;
    setTimeout(() => this.state = State.IN_MATCH, 500);
  }

  shoot() {
    if (!this.flashlight.isSwitchedOn()) {
      this.flashlight.switchOn()
      setTimeout(() => {
        this.flashlight.switchOff()
      }, 50)
    }
    this.nativeAudio.play('gunshot', () => { })
    this.deviceMotion.getCurrentAcceleration().then(
      (acceleration) => {
        if (Math.abs(acceleration.x) >= 8 && Math.abs(acceleration.y) <= 3 && Math.abs(acceleration.z) <= 5) {
          this.handleSuccessfullShot()
        } else {
          this.handleMissedShot()
        }
      },
      (error: any) => console.log(error)
    ).catch(err => console.log(err));
  }

  endMatch(): void { }

  fetchLocation() {
    setTimeout(() => {
      if (this.state === State.WAITING) {
        this.geolocation.getCurrentPosition().then((location: BackgroundGeolocationResponse) => {
          console.log('received location')
          console.log(location.latitude, location.longitude, location.speed)
          if (!location.speed) {
            location.speed = 0
          }
          this.updateLocation(location.latitude, location.longitude, location.speed)
        }, (error: any) => {
          console.error(error)
        })
      }
      else return;
    }, 1000);


  }

  private updateLocation(latitude: number, longitude: number, speed: number): void {
    this.api.post('/locations/' + this.auth.userID, {
      "latitude": latitude,
      "longitude": longitude,
      "speed": speed
    })
      .then(data => {
        console.log(data);
        console.log('sent location')
      }).catch((error: HttpErrorResponse) => {
        console.error(error);
      });

      this.fetchLocation();
  }


}
