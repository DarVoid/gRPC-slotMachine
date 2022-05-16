import {
  HttpClient,
} from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
@Injectable({
  providedIn: 'root'
})
export class SlotsService {

  constructor(private _httpClient: HttpClient,) { }

  public playGame(gameId: string , guyName: string, quote: string): Observable<any> {
    //groupID:string
    // const groupID = '6ac709f4-3b4c-47c8-87a8-3fa83a2708f9'; //mudar para groupLoadID da funcao de cima
    return this._httpClient.post<any>(`http://localhost:8000/play/${gameId}/${guyName}/${quote}`, {}
    );
  }

  public createGame(numJog: number, winPerc: number): Observable<any> {
    //groupID:string
    // const groupID = '6ac709f4-3b4c-47c8-87a8-3fa83a2708f9'; //mudar para groupLoadID da funcao de cima
    return this._httpClient.post<any>(`http://localhost:8000/create/${numJog}/${winPerc}`, {}
    );
  }
}
