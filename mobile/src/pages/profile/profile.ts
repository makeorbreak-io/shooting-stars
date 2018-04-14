import { Component } from '@angular/core';
import { IonicPage, NavController, NavParams } from 'ionic-angular';
import { AuthProvider } from '../../providers/auth/auth';
import { ApiProvider } from '../../providers/api/api';
import { getLocaleDayNames } from '@angular/common';

@IonicPage()
@Component({
  selector: 'page-profile',
  templateUrl: 'profile.html',
})
export class ProfilePage {
  public userData: {} = null;
  constructor(public navCtrl: NavController, public navParams: NavParams, public auth: AuthProvider, public api: ApiProvider) {

    this.getUserData();
  }

  ionViewDidLoad() {
    console.log('ionViewDidLoad ProfilePage');
  }

  getUserData() {
    this.api.get('/user/profile?id=' + this.auth.userID, this.auth.token)
    .then(data => {
      console.log(data);
      this.userData = data;
    }).catch(err => console.log(err));
    this.userData = {gender: "F"};
  }
}
