import { Component, Inject, Input, OnInit } from '@angular/core';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';
export interface DialogData {

  gameId: string,
  name: string,
  luckyQuote: string,
  reward: false
}

@Component({
  selector: 'app-victory',
  templateUrl: './victory.component.html',
  styleUrls: ['./victory.component.scss']
})
export class VictoryComponent implements OnInit {

  @Input()data!: DialogData;
  constructor() { }

  ngOnInit(): void {
  }

}
