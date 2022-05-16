import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { take } from 'rxjs/operators';
import { SlotsService } from '../services/slots.service';
import { VictoryComponent } from './victory/victory.component';

@Component({
  selector: 'app-slot-machine',
  templateUrl: './slot-machine.component.html',
  styleUrls: ['./slot-machine.component.scss']
})
export class SlotMachineComponent implements OnInit {

  guyName: string;
  quote: string;
  win: Boolean;

  gameId: string;

  numeroJogadas: number;
  percentagemWin: number;
  isOnPlay: boolean;
  constructor(private _slot: SlotsService,private _matDialog: MatDialog,) {
    this.guyName = '';
    this.quote = '';
    this.win = false;
    this.gameId = '';
    this.isOnPlay = false;

    this.numeroJogadas = 100;
    this.percentagemWin = 20;
  }

  ngOnInit(): void {
  }
  goIntoExistingGame(): void {
    this.isOnPlay = !this.isOnPlay;
    console.log(this.isOnPlay)
    console.log(this.guyName)
    console.log(this.quote)
  }
  newGame(): void {
    this._slot.createGame(this.numeroJogadas, this.percentagemWin).pipe(take(1)).subscribe((res) => {
      this.gameId = res.data.gameId;
    });;
  }
  play(): void {
    this._slot.playGame(this.gameId, this.guyName, this.quote).pipe(take(1)).subscribe((res) => {
      console.log(res);
      this._matDialog.open(VictoryComponent, {
        autoFocus: false,
        data     : res.data
    })
    });

  }
}

