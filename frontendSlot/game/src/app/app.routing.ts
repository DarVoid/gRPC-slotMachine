import { Route } from '@angular/router';

export const routes: Route[] = [
  {
    path: '**', loadChildren: (): any =>
      import(
        'src/app/slot-machine/slot-machine.module'
      ).then((m) => m.SlotMachineModule)
  },

];
