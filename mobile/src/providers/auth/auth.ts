import { Injectable } from '@angular/core';

@Injectable()
export class AuthProvider {
  id: string = "";
  createdAt: string = "";
  updatedAt: string = "";
  deletedAt: string = "";
  token: string = "";
  userID: string = "";
  gender: string = "";

  setData(data: any) {
    this.id = data.id;
    this.createdAt = data.createdAt;
    this.updatedAt = data.updatedAt;
    this.deletedAt = data.deletedAt;
    this.token = data.token;
    this.userID = data.userID;
    if (data.gender)
      this.gender = data.gender;
    else this.gender = Math.round(Math.random()) ? "M" : "F";
  }

  getData() {
    return {
      "id": this.id,
      "createdAt": this.createdAt,
      "updatedAt": this.updatedAt,
      "deletedAt": this.deletedAt,
      "token": this.token,
      "userID": this.userID,
      "gender": this.gender
    }
  }

  resetData() {
    this.id = "";
    this.createdAt = "";
    this.updatedAt = "";
    this.deletedAt = "";
    this.token = "";
    this.userID = "";
    this.gender = "";
  }

}
