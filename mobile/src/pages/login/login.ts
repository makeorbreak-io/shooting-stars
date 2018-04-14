import { Component } from '@angular/core';
import { IonicPage, NavController, NavParams } from 'ionic-angular';
import { BackgroundGeolocation } from '@ionic-native/background-geolocation';
import { Platform } from 'ionic-angular';
import { RegisterPage } from '../register/register';
import { ApiProvider } from '../../providers/api/api';
import { AuthProvider } from '../../providers/auth/auth';
import { TabsPage } from '../tabs/tabs';

@IonicPage()
@Component({
  selector: 'page-login',
  templateUrl: 'login.html',
})
export class LoginPage {
  mobileDevice: boolean;
  userCredentials = { email: '', password: ''};
  constructor(public navCtrl: NavController, public navParams: NavParams, private backgroundGeolocation: BackgroundGeolocation,
              public platform: Platform, public api: ApiProvider, public auth: AuthProvider) {
      if (this.platform.is('core')) {
        this.mobileDevice = false
      } else {
        this.mobileDevice = true
      }
      // if (this.mobileDevice) {
      //   const config: BackgroundGeolocationConfig = {
      //     desiredAccuracy: 10,
      //     stationaryRadius: 20,
      //     distanceFilter: 30,
      //     debug: false,
      //     startOnBoot: true,
      //     stopOnTerminate: false,
      //   };
      //   this.backgroundGeolocation.configure(config).subscribe((location: BackgroundGeolocationResponse) => {
      //     console.log(location)
      //   });
      // }
  }

  login() {
    console.log('login');
    let json = {
      "email": this.userCredentials.email,
      "password": this.userCredentials.password
    }
    this.api.post('/auth/login', json).then(data => {
      this.auth.setData(data);
      console.log(this.auth.getData());
      this.navCtrl.setRoot(TabsPage, {animate: true, direction: 'forward'});
    })

      .catch(err => {
        console.log(err)
      });
  }

  createAccount() {
    console.log("createAccount");
    this.navCtrl.setRoot(RegisterPage);
  }

  ionViewDidLoad() {
    //this.startPlaying()
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
