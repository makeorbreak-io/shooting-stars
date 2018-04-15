import { Injectable } from '@angular/core';

/*
  Generated class for the GlobalsProvider provider.

  See https://angular.io/guide/dependency-injection for more info on providers
  and Angular DI.
*/
@Injectable()
export class GlobalsProvider {

    public API_URL: String = "https://shooting-stars.herokuapp.com";
    public SOCKET_URL: String = "ws://shooting-stars.herokuapp.com";
}
