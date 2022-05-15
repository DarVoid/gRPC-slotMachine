import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { SlotMachineComponent } from './slot-machine.component';
import { RouterModule } from '@angular/router';
import { slotRoutes } from './slot-machine.routing';


@NgModule({
  declarations: [SlotMachineComponent],
  imports: [
    CommonModule,
    RouterModule.forChild(slotRoutes)
  ]
})
export class SlotMachineModule { }
