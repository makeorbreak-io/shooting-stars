import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { GlobalsProvider } from "../globals/globals";
import { AuthProvider } from '../auth/auth';


@Injectable()
export class ApiProvider {

  constructor(public http: HttpClient, private globals: GlobalsProvider, private authProvider: AuthProvider) {
    console.log('Hello ApiProvider Provider');
  }

  post(path: string, data: Object) {
    let headers: HttpHeaders = new HttpHeaders({
      'Content-Type': 'application/json'
    });
    if (this.authProvider.token) {
      headers = headers.set('Authorization', this.authProvider.token)
    }
    return new Promise((resolve, reject) => {
      this.http.post(this.globals.API_URL + path, JSON.stringify(data), {headers: headers})
        .subscribe(res => {
          resolve(res);
        }, err => {
          reject(err);
        });
    })
  }
}
