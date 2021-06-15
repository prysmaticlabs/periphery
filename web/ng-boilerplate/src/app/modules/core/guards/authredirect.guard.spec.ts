import { AuthredirectGuard } from './authredirect.guard';
import { AuthService } from '../services/auth.service';
import { TestBed, async } from '@angular/core/testing';
import { AuthGuard } from './auth.guard';
import { RouterStateSnapshot, ActivatedRouteSnapshot } from '@angular/router';

class MockActivatedRouteSnapshot {
  public data: any;
}

class MockRouterStateSnapshot {
  url = '/';
}

describe('AuthredirectGuard', () => {
  let authredirectGuard: AuthredirectGuard;
  let authService: AuthService;
  let route: ActivatedRouteSnapshot;
  let state: RouterStateSnapshot;

  beforeEach(async(() => {
    const spy = jasmine.createSpyObj('AuthService', ['token']);
    TestBed.configureTestingModule({
      providers: [
        AuthGuard,
        { provide: AuthService, useValue: spy },
        { provide: ActivatedRouteSnapshot, useValue: MockActivatedRouteSnapshot },
        { provide: RouterStateSnapshot, useValue: MockRouterStateSnapshot },
      ],
    })
      .compileComponents();
    authredirectGuard = TestBed.inject(AuthredirectGuard);
    authService = TestBed.inject(AuthService);
    route = TestBed.inject(ActivatedRouteSnapshot);
    state = TestBed.inject(RouterStateSnapshot);
  }));

  describe('canActivate', () => {
    it('should return false for a logged in user', () => {
      authService.token = 'token';
      expect(authredirectGuard.canActivate(route, state)).toEqual(false);
    });

    it('should return true for an unauthenticated user', () => {
      authService.token = '';
      expect(authredirectGuard.canActivate(route, state)).toEqual(true);
    });
  });
});
