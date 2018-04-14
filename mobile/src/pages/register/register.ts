import { Component } from '@angular/core';
import { IonicPage, NavController, NavParams } from 'ionic-angular';
import { LoginPage } from '../login/login';
import { ApiProvider } from '../../providers/api/api';
import { ToastController } from 'ionic-angular';
import { AuthProvider } from '../../providers/auth/auth';
import { TabsPage } from '../tabs/tabs';

@IonicPage()
@Component({
  selector: 'page-register',
  templateUrl: 'register.html',
})
export class RegisterPage {
  name: String = "";
  email: String = "";
  password: String = "";
  confirmPassword: String = "";
  birthDate: String = "";
  gender: String = "M";
  constructor(public navCtrl: NavController, public navParams: NavParams, private api: ApiProvider, public toastCtrl: ToastController, public auth: AuthProvider) {
  }

  ionViewDidLoad() {
    console.log('ionViewDidLoad RegisterPage');
  }

  back() {
    this.navCtrl.setRoot(LoginPage);
  }

  register() {
    if (!this.verifyFields())
      return;

    let json = {
      "name": this.name,
      "email": this.email,
      "password": this.password,
      "confirmPassword": this.confirmPassword,
      "birthDate": this.birthDate,
      "gender": this.gender
    }
    this.api.post("/auth/register", json).then(data => {
      this.auth.setData(data);
      console.log(this.auth.getData());
      this.navCtrl.setRoot(TabsPage, {animate: true, direction: 'forward'});
    })
      .catch(err => {
        console.log(err)
      });
  }


  verifyFields() {
    let error = "";
    var regex = /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;

    if (this.name === "") {
      error = "Name is empty"
    }
    else if (!regex.test(this.email + "")) {
      error = "Invalid Email";
    }
    else if (this.password !== this.confirmPassword) {
      error = "Passwords don't match";
    }
    else if (this.birthDate !== "") {
      let date = new Date(this.birthDate + "");
      if (date.getTime() > (new Date().getTime()))
        error = "Invalid Birth Date";
    }
    else error = "Birth Date is empty"


    let toast = this.toastCtrl.create({
      message: error,
      duration: 3000,
      position: "top"
    });
    toast.present();

    if (error)
      return false;
    else return true;
  }
}
