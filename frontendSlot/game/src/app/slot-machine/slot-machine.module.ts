import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { SlotMachineComponent } from './slot-machine.component';
import { RouterModule } from '@angular/router';
import { slotRoutes } from './slot-machine.routing';

import { MatButtonModule } from '@angular/material/button';
import { MatDialogModule } from '@angular/material/dialog';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { VictoryComponent } from './victory/victory.component';


const materialModules = [
  MatButtonModule,
  MatDialogModule,
  MatCheckboxModule,
  MatInputModule,
  MatFormFieldModule,
  MatProgressSpinnerModule,
  CommonModule,
  FormsModule,
];

@NgModule({
  declarations: [SlotMachineComponent, VictoryComponent],
  imports: [
    RouterModule.forChild(slotRoutes),
    ...materialModules
  ],
  exports: [
    ...materialModules
  ]
})
export class SlotMachineModule { }
