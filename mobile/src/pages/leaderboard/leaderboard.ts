import { Component } from '@angular/core';
import { IonicPage, NavController, NavParams } from 'ionic-angular';
import { ApiProvider } from '../../providers/api/api';
import { AuthProvider } from '../../providers/auth/auth';

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
  public players: {} = null

  constructor(public navCtrl: NavController, public navParams: NavParams, public api: ApiProvider, public auth: AuthProvider) {
    this.getLeaderboard();
  }

  getLeaderboard() {
    this.api.get('/stats/topWins', this.auth.token)
    .then(data => {
      this.players = data
    }).catch(err => console.log(err));
  }

  ionViewDidLoad() {
    console.log('ionViewDidLoad LeaderboardPage');
  }

}
