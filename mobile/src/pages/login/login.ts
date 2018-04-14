import { Component } from '@angular/core';
import { IonicPage, NavController, NavParams } from 'ionic-angular';
import { BackgroundGeolocation, BackgroundGeolocationConfig, BackgroundGeolocationResponse } from '@ionic-native/background-geolocation';
import { Platform } from 'ionic-angular';
import { GamePageModule } from '../game/game.module';
import { GamePage } from '../game/game';

@IonicPage()
@Component({
  selector: 'page-login',
  templateUrl: 'login.html',
})
export class LoginPage {
  mobileDevice: boolean
  constructor(public navCtrl: NavController, public navParams: NavParams) {
  }

  ionViewDidLoad() {
    console.log('ionViewDidLoad LoginPage');
    this.onLoginSuccessfull()
  }

  onLoginSuccessfull() {
    this.navCtrl.setRoot(GamePage);
  }

}
