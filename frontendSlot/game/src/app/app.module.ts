import { NgModule, CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { ExtraOptions, PreloadAllModules, RouterModule } from '@angular/router';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { routes } from './app.routing';
import { AppComponent } from './app.component';
import { SlotsService } from './services/slots.service';
import { HttpClientModule } from '@angular/common/http';

const routerConfig: ExtraOptions = {
  scrollPositionRestoration: 'enabled',
  preloadingStrategy: PreloadAllModules,
  relativeLinkResolution: 'legacy',
};

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    BrowserAnimationsModule,
    RouterModule.forRoot(routes, routerConfig),

  ],
  schemas:[CUSTOM_ELEMENTS_SCHEMA],
  providers: [SlotsService],
  bootstrap: [AppComponent]
})
export class AppModule { }
