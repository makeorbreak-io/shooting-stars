import { Component } from '@angular/core';
import { IonicPage, NavController, NavParams } from 'ionic-angular';
import { AuthProvider } from '../../providers/auth/auth';
import { ApiProvider } from '../../providers/api/api';
import { LoginPage } from '../login/login';
import { App } from 'ionic-angular';
import { AlertController } from 'ionic-angular';

@IonicPage()
@Component({
  selector: 'page-profile',
  templateUrl: 'profile.html',
})
export class ProfilePage {
  public userData: any = {};
  public profileImg: string = "";
  constructor(public navCtrl: NavController, public navParams: NavParams, public auth: AuthProvider, public api: ApiProvider, public app: App, public alertCtrl: AlertController) {
    this.getUserData();
  }

  getUserData() {
    this.api.get('/users/' + this.auth.userID, this.auth.token)
      .then(data => {
        this.userData = data;
        if (this.userData.gender === "F")
          this.profileImg = "assets/imgs/cowgirl.png";
        else if (this.userData.gender == "M")
          this.profileImg = "assets/imgs/cowboy.png";

      }).catch(err => console.log(err));
  }

  logout() {
    let alert = this.alertCtrl.create({
      title: "Leave current session?",
      buttons: [
        {
          text: "Cancel",
          role: "cancel",
        },
        {
          text: "Logout",
          handler: () => {
            this.auth.resetData();
            this.app.getRootNav().setRoot(LoginPage);
          }
        }
      ]
    });
    alert.present();
  }
}
