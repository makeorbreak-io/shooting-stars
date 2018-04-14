import { Component } from '@angular/core';
import { Platform } from 'ionic-angular';
import { StatusBar } from '@ionic-native/status-bar';
import { SplashScreen } from '@ionic-native/splash-screen';
<<<<<<< HEAD
=======
import { NativeAudio } from '@ionic-native/native-audio';
>>>>>>> fe9840a7c9e08e1e766b0a4444eb4056c3338d59
import { LoginPage } from '../pages/login/login';

@Component({
  templateUrl: 'app.html'
})
export class MyApp {
  rootPage: any = LoginPage;

<<<<<<< HEAD
  constructor(platform: Platform, statusBar: StatusBar, splashScreen: SplashScreen) {

=======
  constructor(platform: Platform, statusBar: StatusBar, splashScreen: SplashScreen, private nativeAudio: NativeAudio) {
>>>>>>> fe9840a7c9e08e1e766b0a4444eb4056c3338d59
    platform.ready().then(() => {
      // Okay, so the platform is ready and our plugins are available.
      // Here you can do any higher level native things you might need.
      statusBar.styleDefault();
      splashScreen.hide();
      this.nativeAudio.preloadSimple('westernWhistle', 'assets/sounds/western_whistle.mp3').then(() => console.log('here'), () => {})
    });

  }
}
