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
  userCredentials = { email: '', password: '' };
  constructor(public navCtrl: NavController, public navParams: NavParams, public auth: AuthProvider, public api: ApiProvider) {
  }

  ionViewDidLoad() {
    console.log('ionViewDidLoad LoginPage');
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
      this.navCtrl.setRoot(TabsPage, { animate: true, direction: 'forward' });
    })

      .catch(err => {
        console.log(err)
      });
  }

  createAccount() {
    this.navCtrl.setRoot(RegisterPage);
  }


}
