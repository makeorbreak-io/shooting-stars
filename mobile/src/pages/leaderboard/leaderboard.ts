import { Component } from '@angular/core';
import { IonicPage, NavController, NavParams } from 'ionic-angular';

/**
 * Generated class for the LeaderboardPage page.
 *
 * See https://ionicframework.com/docs/components/#navigation for more info on
 * Ionic pages and navigation.
 */

@IonicPage()
@Component({
  selector: 'page-leaderboard',
  templateUrl: 'leaderboard.html',
})
export class LeaderboardPage {
  players = [{ 'name': 'John', 'wins': 5 },
    { 'name': 'Dow', 'wins': 4 },
    { 'name': 'Mary', 'wins': 3 },
    { 'name': 'Jane', 'wins': 2 },
    { 'name': 'Gustavo', 'wins': 1 }];

  constructor(public navCtrl: NavController, public navParams: NavParams) {
  }

  ionViewDidLoad() {
    console.log('ionViewDidLoad LeaderboardPage');
  }

}
