import { Component } from '@angular/core';
import { IonicPage, NavController, NavParams } from 'ionic-angular';
import { AuthProvider } from '../../providers/auth/auth';
import { ApiProvider } from '../../providers/api/api';
import { getLocaleDayNames } from '@angular/common';
import { LoginPage } from '../login/login';
import { App } from 'ionic-angular';

@IonicPage()
@Component({
  selector: 'page-profile',
  templateUrl: 'profile.html',
})
export class ProfilePage {
  public userData: {} = null;
  constructor(public navCtrl: NavController, public navParams: NavParams, public auth: AuthProvider, public api: ApiProvider, public app: App) {
    this.getUserData();
  }

  ionViewDidLoad() {
  }

  getUserData() {
    this.api.get('/users/' + this.auth.userID, this.auth.token)
    .then(data => {
      console.log(data["name"]);
      this.userData = data;
    }).catch(err => console.log(err));
    this.userData = {gender: "M"};
  }

  logout() {
    this.auth.resetData();
    this.app.getRootNav().setRoot(LoginPage);
  }
}
