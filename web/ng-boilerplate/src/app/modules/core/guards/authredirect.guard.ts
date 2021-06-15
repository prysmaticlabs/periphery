import { Injectable } from '@angular/core';
import {
  CanActivate,
  ActivatedRouteSnapshot,
  UrlTree,
  RouterStateSnapshot,
} from '@angular/router';
import { Observable } from 'rxjs';

import { AuthService } from '../services/auth.service';

// This redirect guard prevents users from visiting the home page by default
// if they are authenticated.
@Injectable({ providedIn: 'root' })
export class AuthredirectGuard implements CanActivate {
  constructor(
    private authService: AuthService,
  ) { }

  canActivate(
    next: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Observable<any> | Promise<boolean|UrlTree> | boolean {
    if (this.authService.token) {
      return false;
    }
    return true;
  }
}
