import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { GlobalsProvider } from "../globals/globals";


@Injectable()
export class ApiProvider {

  constructor(public http: HttpClient, private globals: GlobalsProvider) {
    console.log('Hello ApiProvider Provider');
  }

  post(path: string, data: Object) {
    let headers: HttpHeaders = new HttpHeaders({
      'Content-Type': 'application/json'
    });
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
