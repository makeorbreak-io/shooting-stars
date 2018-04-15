import { Component } from '@angular/core';
import { Platform } from 'ionic-angular';
import { StatusBar } from '@ionic-native/status-bar';
import { SplashScreen } from '@ionic-native/splash-screen';
import { NativeAudio } from '@ionic-native/native-audio';
import { LoginPage } from '../pages/login/login';
import { Flashlight } from '@ionic-native/flashlight';

@Component({
  templateUrl: 'app.html',
  providers: [ Flashlight ]
})
export class MyApp {
  rootPage: any = LoginPage;

  constructor(platform: Platform, statusBar: StatusBar, splashScreen: SplashScreen, private nativeAudio: NativeAudio, private flashlight: Flashlight) {
    platform.ready().then(() => {
      // Okay, so the platform is ready and our plugins are available.
      // Here you can do any higher level native things you might need.
      statusBar.styleDefault();
      splashScreen.hide();
      this.flashlight.available(); // Hopefully this will ask for permissions
      this.nativeAudio.preloadSimple('gunshot', 'assets/sounds/gun_shot.mp3').then((e) => console.log('Loaded gunshot', e)).catch((e) => console.error('Error loading gunshot', e))
      this.nativeAudio.preloadSimple('westernWhistle', 'assets/sounds/western_whistle.mp3').then((e) => console.log('Loaded whistle', e)).catch((e) => console.error('Error loading whistle', e))
    });

  }
}
