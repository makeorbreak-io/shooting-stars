import { Component } from '@angular/core';

import { GamePage } from '../game/game';
import { LeaderboardPage } from '../leaderboard/leaderboard';
import { ProfilePage } from '../profile/profile';

@Component({
  templateUrl: 'tabs.html'
})
export class TabsPage {

  tab1Root = ProfilePage;
  tab2Root = GamePage;
  tab3Root = LeaderboardPage;

  constructor() {

  }
}
