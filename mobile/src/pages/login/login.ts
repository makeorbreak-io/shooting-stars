import { Component } from '@angular/core';
import { IonicPage, NavController, NavParams } from 'ionic-angular';
import { RegisterPage } from '../register/register';
import { ApiProvider } from '../../providers/api/api';
import { AuthProvider } from '../../providers/auth/auth';
import { NativeAudio } from '@ionic-native/native-audio';
import { TabsPage } from '../tabs/tabs';


@IonicPage()
@Component({
  selector: 'page-login',
  templateUrl: 'login.html'
})
export class LoginPage {
  mobileDevice: boolean;
  userCredentials = { email: 'user@email.com', password: '123' };
  constructor(public navCtrl: NavController, public navParams: NavParams, public auth: AuthProvider, public api: ApiProvider, private nativeAudio: NativeAudio) {
    console.log(this.auth.getData());
  }

  ionViewDidLoad() {
  }

  login() {
    let json = {
      "email": this.userCredentials.email,
      "password": this.userCredentials.password
    }
    this.api.post('/auth/login', json).then(data => {
      this.auth.setData(data);
      this.navCtrl.setRoot(TabsPage, { animate: true, direction: 'forward' });
    })
      .catch(err => {
        console.log(err)
      });
  }t

  createAccount() {
    this.navCtrl.setRoot(RegisterPage);
  }


}
